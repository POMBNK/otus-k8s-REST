package main

import (
	"context"
	"github.com/POMBNK/kuberRest/internal/delivery/user"
	userStorage "github.com/POMBNK/kuberRest/internal/repository/user"
	userService "github.com/POMBNK/kuberRest/internal/service/user"
	"github.com/POMBNK/kuberRest/pkg/client/postgres"
	_ "github.com/POMBNK/kuberRest/statik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/rakyll/statik/fs"

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
	}
	engine := fiber.New(cfg)
	// Swagger
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("Swagger unable: failed to create statik fs: %s", err.Error())
	}
	engine.Use("/docs", filesystem.New(filesystem.Config{
		Root: statikFS,
	}))

	user.New(service).Register(engine)
	engine.Listen(":8080")
	//var listener net.Listener
	//var listenErr error
	//listener, listenErr = net.Listen("tcp", "127.0.0.1:8080")
	//if listenErr != nil {
	//	log.Fatal(listenErr)
	//}
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	//
	//go func() {
	//	if err := engine.Server().Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		panic(err)
	//	}
	//}()

	//s.logs.Println("Server started")

	//<-interrupt
	//s.logs.Println("Shutting down server...")

}
