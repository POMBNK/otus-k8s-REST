package main

import (
	"context"
	"github.com/POMBNK/kuberRest/internal/delivery/user"
	userStorage "github.com/POMBNK/kuberRest/internal/repository/user"
	userService "github.com/POMBNK/kuberRest/internal/service/user"
	"github.com/POMBNK/kuberRest/pkg/client/postgres"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	pgClient, err := postgres.NewClient(context.Background(), 5)
	if err != nil {
		panic(err)
	}

	// user init
	repository := userStorage.New(pgClient)
	service := userService.New(repository, pgClient)

	// init engine
	cfg := fiber.Config{
		AppName:           "kuberRest",
		EnablePrintRoutes: true,
		StrictRouting:     false,
	}
	engine := fiber.New(cfg)
	// Swagger
	engine.Static("/swagger", "./doc/dist")
	user.New(service).Register(engine)

	// start server
	log.Panicln(engine.Listen(":8080"))
}
