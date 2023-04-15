package user

import (
	ut "github.com/mnfirdauss/clean-architecture/entity/user"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *ut.User) error
}

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepo {
	return &userRepo{
		db,
	}
}
