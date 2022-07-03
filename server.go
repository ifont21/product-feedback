package main

import (
	"github.com/ifont21/product-feedback-service/internal/handlers"
	"github.com/ifont21/product-feedback-service/internal/helpers"
	"github.com/ifont21/product-feedback-service/internal/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db := infrastructure.NewConnection()

	defer db.Close()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, error=${error}\n",
	}))

	e.Use(echo.WrapMiddleware(helpers.EnsureValidToken()))

	v1 := e.Group("/api")

	routerHandler := handlers.NewRouterHandleFactory(db)
	routerHandler.RegisterRoutes(v1)

	e.Logger.Fatal(e.Start(":1323"))
}
