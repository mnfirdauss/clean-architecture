package main

import (
	"github.com/labstack/echo/v4"

	"github.com/mnfirdauss/clean-architecture/common"
	"github.com/mnfirdauss/go-gorm/middleware"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()

	middleware.LogMiddleware(e)

	userGroup := e.Group("/users")
	userGroup.POST("/", handler.UserHandler.CreateUser)

	return e
}
