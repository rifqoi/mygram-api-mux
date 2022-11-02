package main

import (
	"github.com/rifqoi/mygram-api-mux/api/middlewares"
	"github.com/rifqoi/mygram-api-mux/api/routes"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}
	logger.Sugar()
	middleware := middlewares.NewMiddleware(logger)

	app := routes.NewRouter(middleware)
	app.Run()
}
