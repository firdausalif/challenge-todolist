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

	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo, timeoutCtx)
	todoController := controllers.NewTodoController(todoService, validate)

	activityGroup := app.Group("/activity-groups")
	activityGroup.Get("", activityController.FetchHandler)
	activityGroup.Post("", activityController.StoreActivityHandler)
	activityGroup.Get("/:id", activityController.DetailActivityHandler)
	activityGroup.Delete("/:id", activityController.DeleteActivityHandler)
	activityGroup.Patch("/:id", activityController.EditActivityHandler)

	todoGroup := app.Group("/todo-items")
	todoGroup.Get("", todoController.FetchHandler)
	todoGroup.Post("", todoController.StoreTodoHandler)
	todoGroup.Get("/:id", todoController.DetailTodoHandler)
	todoGroup.Delete("/:id", todoController.DeleteTodoHandler)
	todoGroup.Patch("/:id", todoController.EditTodoHandler)

}
