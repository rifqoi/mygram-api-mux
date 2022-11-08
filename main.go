package main

import (
	"github.com/rifqoi/mygram-api-mux/api/controller"
	"github.com/rifqoi/mygram-api-mux/api/middleware"
	"github.com/rifqoi/mygram-api-mux/api/routes"
	"github.com/rifqoi/mygram-api-mux/config"
	"github.com/rifqoi/mygram-api-mux/repository/postgres"
	"github.com/rifqoi/mygram-api-mux/repository/postgres/db"
	"github.com/rifqoi/mygram-api-mux/services"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	pgDatabase, err := config.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	q := db.New(pgDatabase)

	userRepo := postgres.NewUserRepository(q)

	userService := services.NewUserService(userRepo)

	userController := controller.NewUserController(userService)

	middleware := middleware.NewMiddleware(logger, userService)

	app := routes.NewRouter(middleware, userController)
	app.Run()
}
