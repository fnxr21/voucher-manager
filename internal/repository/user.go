package repositories

import "github.com/fnxr21/voucher-manager/internal/model"

type User interface {
	Register(user model.User) (*model.User, error)
	Login(username string) (*model.User, error)
	Reauth(id uint) (*model.User, error)
}

func (r *repository) Register(user model.User) (*model.User, error) {

	err := r.db.Create(&user).
		Error

	return &user, err
}

func (r *repository) Login(username string) (*model.User, error) {

	var user model.User
	err := r.db.First(&user, "username=?", username).
		Error

	return &user, err
}
func (r *repository) Reauth(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "id=?", id).Error
	return &user, err
}
