package services

import (
	"context"
	"fmt"
	"github.com/firdausalif/challenge-todolist/app/models"
	"github.com/firdausalif/challenge-todolist/app/requests"
	"github.com/firdausalif/challenge-todolist/pkg/constant"
	"time"
)

type TodoService struct {
	todoRepository models.TodoRepository
	ctxTimeout     time.Duration
}

func NewTodoService(td models.TodoRepository, ctxTimeout time.Duration) models.TodoService {
	return &TodoService{
		todoRepository: td,
		ctxTimeout:     ctxTimeout,
	}
}

func (t TodoService) Create(ctx context.Context, req requests.CreateTodo) (*models.Todo, int, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	req.IsActive = true

	if req.Priority == "" {
		req.Priority = "very-high"
	}

	dt := time.Now()
	record := &models.Todo{
		ActivityGroupID: req.ActivityGroupID,
		Title:           req.Title,
		IsActive:        &req.IsActive,
		Priority:        req.Priority,
		CreatedAt:       dt.Local(),
		UpdatedAt:       dt.Local(),
	}

	_, err := t.todoRepository.InsertTodos(ctx, record)
	if err != nil {
		fmt.Println(err)
		return nil, constant.CodeInternalServerError, err
	}

	return record, constant.CodeSuccess, nil
}

func (t TodoService) Update(ctx context.Context, id uint64, req requests.UpdateTodo) (*models.Todo, int, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	dt := time.Now()
	record := &models.Todo{
		ID:              id,
		Title:           req.Title,
		Priority:        req.Priority,
		ActivityGroupID: req.ActivityGroupID,
		IsActive:        &req.IsActive,
		UpdatedAt:       dt.Local(),
	}

	updateActivity, err := t.todoRepository.UpdateTodo(ctx, record)

	if err != nil {
		return nil, constant.CodeInternalServerError, err
	}

	return updateActivity, constant.CodeSuccess, nil
}

func (t TodoService) Remove(ctx context.Context, id uint64) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	err := t.todoRepository.RemoveTodo(ctx, id)
	if err != nil {
		return constant.CodeInternalServerError, err
	}

	return constant.CodeSuccess, nil
}

func (t TodoService) Fetch(ctx context.Context, id uint64) ([]*models.Todo, int, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fetchActivity, err := t.todoRepository.FetchTodos(ctx, id)
	if err != nil {
		return nil, constant.CodeInternalServerError, err
	}

	return fetchActivity, constant.CodeSuccess, nil
}

func (t TodoService) GetById(ctx context.Context, id uint64) (*models.Todo, int, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	activity, err := t.todoRepository.GetTodoByID(ctx, id)
	if err != nil {
		return nil, constant.CodeInternalServerError, err
	}

	return activity, constant.CodeSuccess, nil
}

func (t TodoService) GetActivityGroup(ctx context.Context) ([]*models.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	group, err := t.todoRepository.GetActivityGroup(ctx)
	if err != nil {
		return nil, err
	}

	return group, nil
}
