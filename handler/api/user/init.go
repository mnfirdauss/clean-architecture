package user

import (
	uc "github.com/mnfirdauss/clean-architecture/usecase/user"
)

type UserHandler struct {
	userUC uc.UserUsecase
}

func New(userUC uc.UserUsecase) *UserHandler {
	return &UserHandler{
		userUC,
	}
}
