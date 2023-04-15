package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mnfirdauss/clean-architecture/common"
	conf "github.com/mnfirdauss/clean-architecture/conf"
	userHandler "github.com/mnfirdauss/clean-architecture/handler/api/user"
	userRepo "github.com/mnfirdauss/clean-architecture/repo/user"
	userUsecase "github.com/mnfirdauss/clean-architecture/usecase/user"
)

func startApp() *echo.Echo {
	userRepo := userRepo.New(conf.DBMysql)
	userUsecase := userUsecase.New(userRepo)
	userHandler := userHandler.New(userUsecase)

	handler := common.Handler{
		UserHandler: userHandler,
	}

	router := StartRoute(handler)
	return router
}
