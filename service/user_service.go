package service

import (
	"qishan233/hello-wire/config"
	"qishan233/hello-wire/model"
)

type UserServiceImpl struct {
	adminId string
}

type UserService interface {
	GetUser(id string) *model.User
	GetAdmin() *model.Admin
}

func (u *UserServiceImpl) GetUser(id string) *model.User {
	return &model.User{Name: "user_" + id}
}

func (u *UserServiceImpl) GetAdmin() *model.Admin {
	return &model.Admin{
		Name:  u.adminId,
		Scope: "admin",
	}
}

func NewUserService(config config.Config) UserService {
	return &UserServiceImpl{
		adminId: config.GetString("admin_id"),
	}
}

type SpecialRightServiceImpl struct {
	SpecialUserService UserService
}

func NewSpecialRightService(config config.Config) *SpecialRightServiceImpl {
	return &SpecialRightServiceImpl{
		SpecialUserService: &UserServiceImpl{adminId: "special user"},
	}
}
