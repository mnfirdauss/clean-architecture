package user

import ut "github.com/mnfirdauss/clean-architecture/entity/user"

func (uc *userUsecase) CreateUser(user *ut.User) error {
	return uc.userRepo.CreateUser(user)
}
