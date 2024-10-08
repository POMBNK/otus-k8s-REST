package main

import (
	"context"
	"github.com/POMBNK/kuberRest/internal/delivery/user"
	userStorage "github.com/POMBNK/kuberRest/internal/repository/user"
	userService "github.com/POMBNK/kuberRest/internal/service/user"
	"github.com/POMBNK/kuberRest/pkg/client/metrics"
	"github.com/POMBNK/kuberRest/pkg/client/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	mw "github.com/oapi-codegen/fiber-middleware"
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
	userServer := user.New(service)

	// init engine
	cfg := fiber.Config{
		AppName:           "kuberRest",
		EnablePrintRoutes: true,
		StrictRouting:     false,
	}

	engine := fiber.New(cfg)

	prometheus := metrics.New()

	engine.Use(cors.New())
	engine.Use(logger.New())

	// Swagger
	prometheus.Handle(engine, "/metrics")
	engine.Static("/swagger", "./doc/dist")
	engine.Use(prometheus.MetricMiddleware())
	engine.Use(mw.OapiRequestValidator(userServer.Swagger))
	userServer.Register(engine)

	//start server

	log.Panicln(engine.Listen(":8080"))

	//log.Panicln(engine.Listen("0.0.0.0:8080"))
}
