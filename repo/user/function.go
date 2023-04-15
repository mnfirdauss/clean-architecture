package user

import ut "github.com/mnfirdauss/clean-architecture/entity/user"

func (r *userRepo) CreateUser(user *ut.User) error {
	err := r.db.Save(&user)
	if err != nil {
		return err.Error
	}

	return nil
}
