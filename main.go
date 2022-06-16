package main

import (
	"github.com/firdausalif/challenge-todolist/cmd/server"
	"github.com/firdausalif/challenge-todolist/pkg/config"
	"os"
)

func main() {

	env := ".env"
	if os.Getenv("APP_ENV") == "test" {
		env = ".env.test"
	}

	// setup various configuration for app
	config.LoadAllConfigs(env)

	server.Serve()
}
