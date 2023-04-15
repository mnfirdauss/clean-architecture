package common

import (
	userHandler "github.com/mnfirdauss/clean-architecture/handler/api/user"
)

type Handler struct {
	UserHandler *userHandler.UserHandler
}
