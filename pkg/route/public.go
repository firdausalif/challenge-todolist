package route

import (
	"github.com/firdausalif/challenge-todolist/app/controllers"
	"github.com/firdausalif/challenge-todolist/app/repositories"
	"github.com/firdausalif/challenge-todolist/app/services"
	"github.com/firdausalif/challenge-todolist/pkg/validator"
	"github.com/firdausalif/challenge-todolist/platform/database"
	"github.com/gofiber/fiber/v2"
	"time"
)

// PublicRoutes func for describe group of public route.
func PublicRoutes(app *fiber.App) {
	db := database.GetDB()
	validate := validator.NewValidator()
	timeoutCtx := time.Duration(100) * time.Second

	activityRepo := repositories.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepo, timeoutCtx)
	activityController := controllers.NewActivityController(activityService, validate)

	activityGroup := app.Group("/activity-groups")
	activityGroup.Get("", activityController.FetchHandler)
	activityGroup.Post("", activityController.StoreActivityHandler)
}
