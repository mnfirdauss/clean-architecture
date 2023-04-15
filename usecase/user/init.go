package user

import (
	ut "github.com/mnfirdauss/clean-architecture/entity/user"
	ur "github.com/mnfirdauss/clean-architecture/repo/user"
)

type UserUsecase interface {
	CreateUser(user *ut.User) error
}

type userUsecase struct {
	userRepo ur.UserRepo
}

func New(userRepo ur.UserRepo) *userUsecase {
	return &userUsecase{
		userRepo,
	}
}
