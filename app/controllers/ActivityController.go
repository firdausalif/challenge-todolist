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

type ActivityController struct {
	activityService models.ActivityService
	Validate        *validator.Validate
}

type ActivityControllerInterface interface {
	FetchHandler(ctx *fiber.Ctx) error
	StoreActivityHandler(ctx *fiber.Ctx) error
	DeleteActivityHandler(ctx *fiber.Ctx) error
	EditActivityHandler(ctx *fiber.Ctx) error
	DetailActivityHandler(ctx *fiber.Ctx) error
}

func NewActivityController(as models.ActivityService, validate *validator.Validate) ActivityControllerInterface {
	return &ActivityController{
		activityService: as,
		Validate:        validate,
	}
}

func (d *ActivityController) FetchHandler(ctx *fiber.Ctx) error {
	key := "activities"
	activities, errCache := cache.Get(key)
	if errCache == cacheNotFound {
		activities, _, err := d.activityService.Fetch(ctx.Context())
		if err != nil {
			return helper.JsonERROR(ctx, err)
		}
		go cache.SetWithTTL(key, activities, time.Hour)
		return helper.JsonSUCCESS(ctx, activities)
	}
	return helper.JsonSUCCESS(ctx, activities)
}

func (d *ActivityController) StoreActivityHandler(ctx *fiber.Ctx) error {
	var req requests.CreateActivity
	if err := ctx.BodyParser(&req); err != nil {
		return helper.JsonERROR(ctx, err)
	}

	if req.Title == "" {
		return helper.JsonValidationError(ctx, "title cannot be null")
	}

	resp, _, err := d.activityService.Create(ctx.Context(), req)
	if err != nil {
		return helper.JsonERROR(ctx, err)
	}

	if resp != nil {
		key := fmt.Sprintf("activity-id-%d", resp.ID)
		go cache.SetWithTTL(key, resp, time.Hour)
		go cache.Remove("activities")
	}

	return helper.JsonCreated(ctx, resp)
}

func (d *ActivityController) DeleteActivityHandler(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.JsonERROR(ctx, err)
	}

	_, err = d.activityService.Remove(ctx.Context(), id)
	if err != nil {
		return helper.JsonNotFound(ctx, fmt.Sprintf("Activity with ID %d Not Found", id))
	}

	key := fmt.Sprintf("activity-id-%d", id)
	go cache.Remove(key)
	go cache.Remove("activities")

	return helper.JsonSuccessDelete(ctx)
}

func (d *ActivityController) EditActivityHandler(ctx *fiber.Ctx) error {
	var req requests.UpdateActivity

	if err := ctx.BodyParser(&req); err != nil {
		return helper.JsonERROR(ctx, err)
	}

	if req.Title == "" {
		return helper.JsonValidationError(ctx, "title cannot be null")
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.JsonERROR(ctx, err)
	}

	resp, _, err := d.activityService.Update(ctx.Context(), id, req)
	if err != nil {
		return helper.JsonNotFound(ctx, fmt.Sprintf("Activity with ID %d Not Found", id))
	}

	if resp != nil {
		key := fmt.Sprintf("activity-id-%d", id)
		go cache.SetWithTTL(key, resp, time.Hour)
		go cache.Remove("activities")
	}

	return helper.JsonSUCCESS(ctx, resp)
}

func (d *ActivityController) DetailActivityHandler(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.JsonERROR(ctx, err)
	}

	key := fmt.Sprintf("activity-id-%d", id)
	activity, err := cache.Get(key)

	if err == cacheNotFound {
		activity, _, err := d.activityService.GetById(ctx.Context(), id)
		if err != nil {
			return helper.JsonNotFound(ctx, fmt.Sprintf("Activity with ID %d Not Found", id))
		}
		go cache.SetWithTTL(key, activity, time.Hour)
		return helper.JsonSUCCESS(ctx, activity)
	}

	return helper.JsonSUCCESS(ctx, activity)
}
