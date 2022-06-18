package controllers

import (
	"fmt"
	"github.com/firdausalif/challenge-todolist/app/models"
	"github.com/firdausalif/challenge-todolist/app/requests"
	"github.com/firdausalif/challenge-todolist/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type TodoController struct {
	todoService models.TodoService
	Validate    *validator.Validate
}

type TodoControllerInterface interface {
	FetchHandler(ctx *fiber.Ctx) error
	StoreTodoHandler(ctx *fiber.Ctx) error
	DeleteTodoHandler(ctx *fiber.Ctx) error
	EditTodoHandler(ctx *fiber.Ctx) error
	DetailTodoHandler(ctx *fiber.Ctx) error
}

func NewTodoController(ts models.TodoService, validate *validator.Validate) TodoControllerInterface {
	return &TodoController{
		todoService: ts,
		Validate:    validate,
	}
}

func (t *TodoController) FetchHandler(ctx *fiber.Ctx) error {
	var id int
	var err error
	agi := ctx.Query("activity_group_id")

	if agi != "" {
		id, err = strconv.Atoi(agi)
		if err != nil {
			id = 0
		}
	}

	key := fmt.Sprintf("todos-%d", id)
	todo, err := cache.Get(key)
	if err == cacheNotFound {
		go func() {
			todo, _, _ := t.todoService.Fetch(ctx.Context(), uint64(id))
			respsChan <- todo
		}()
		todo = <-respsChan
		go cache.SetWithTTL(key, todo, 10*time.Minute)
		return helper.JsonSUCCESS(ctx, todo)
	}
	return helper.JsonSUCCESS(ctx, todo.([]*models.Todo))
}

func (t *TodoController) StoreTodoHandler(ctx *fiber.Ctx) error {
	var req requests.CreateTodo

	if err := ctx.BodyParser(&req); err != nil {
		return helper.JsonERROR(ctx, err)
	}

	if req.Title == "" {
		return helper.JsonValidationError(ctx, "title cannot be null")
	}

	if req.ActivityGroupID == 0 {
		return helper.JsonValidationError(ctx, "activity_group_id cannot be null")
	}

	go func() {
		resp, _, _ := t.todoService.Create(ctx.Context(), req)
		respChan <- resp
	}()
	resp := <-respChan
	key := fmt.Sprintf("todo-id-%d", resp.ID)
	go cache.SetWithTTL(key, resp, time.Hour)
	go cache.Remove("todos-0")
	return helper.JsonCreated(ctx, resp)
}

func (t *TodoController) DeleteTodoHandler(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.JsonERROR(ctx, err)
	}

	_, err = t.todoService.Remove(ctx.Context(), uint64(id))
	if err != nil {
		return helper.JsonNotFound(ctx, fmt.Sprintf("Todo with ID %d Not Found", id))
	}
	key := fmt.Sprintf("todo-id-%d", id)
	go cache.Remove(key)
	go cache.Remove("todos-0")
	return helper.JsonSuccessDelete(ctx)
}

func (t *TodoController) EditTodoHandler(ctx *fiber.Ctx) error {
	var req requests.UpdateTodo

	if err := ctx.BodyParser(&req); err != nil {
		return helper.JsonERROR(ctx, err)
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.JsonERROR(ctx, err)
	}

	resp, _, err := t.todoService.Update(ctx.Context(), uint64(id), req)
	if err != nil {
		return helper.JsonNotFound(ctx, fmt.Sprintf("Todo with ID %d Not Found", id))
	}

	if resp != nil {
		key := fmt.Sprintf("todo-id-%d", resp.ID)
		go cache.SetWithTTL(key, resp, 10*time.Minute)
		go cache.Remove("todos-0")
	}

	return helper.JsonSUCCESS(ctx, resp)
}

func (t *TodoController) DetailTodoHandler(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.JsonERROR(ctx, err)
	}

	key := fmt.Sprintf("todo-id-%d", id)
	todo, err := cache.Get(key)

	if err == cacheNotFound {
		todo, _, err := t.todoService.GetById(ctx.Context(), uint64(id))
		if err != nil {
			return helper.JsonNotFound(ctx, fmt.Sprintf("Todo with ID %d Not Found", id))
		}
		go cache.SetWithTTL(key, todo, 10*time.Minute)

		return helper.JsonSUCCESS(ctx, todo)
	}

	return helper.JsonSUCCESS(ctx, todo)
}
