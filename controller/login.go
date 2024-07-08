package controller

import "qishan233/hello-wire/service"

type LoginController struct {
	userSvc service.UserService
}

func NewLoginController(userSvc service.UserService) *LoginController {
	return &LoginController{
		userSvc: userSvc,
	}
}
