package server

import (
	"fmt"
	"github.com/firdausalif/challenge-todolist/pkg/config"
	"github.com/firdausalif/challenge-todolist/pkg/middleware"
	"github.com/firdausalif/challenge-todolist/pkg/route"
	"github.com/firdausalif/challenge-todolist/platform/database"
	"github.com/firdausalif/challenge-todolist/platform/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Serve ..
func Serve() {
	appCfg := config.AppCfg()

	logger.SetUpLogger()
	logr := logger.GetLogger()

	if err := database.ConnectDB(); err != nil {
		logr.Panicf("failed database setup. error: %v", err)
	}

	// Define Fiber config & app.
	fiberCfg := config.FiberConfig()
	app := fiber.New(fiberCfg)
	app.Use(cache.New(cache.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			path := c.Path()
			if c.Query("activity_group_id") != "" {
				path += "-" + c.Query("activity_group_id")
			}

			return path
		},
		CacheControl: true,
		CacheHeader:  "X-Cache",
		Expiration:   1 * time.Minute,
	}))

	// Attach Middlewares.
	middleware.FiberMiddleware(app)

	// Routes.
	route.GeneralRoute(app)
	//route.SwaggerRoute(app)
	route.PublicRoutes(app)
	route.PrivateRoutes(app)
	route.NotFoundRoute(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		logr.Infoln("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	serverAddr := fmt.Sprintf(":%d", appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		logr.Errorf("Oops... server is not running! error: %v", err)
	}
}
