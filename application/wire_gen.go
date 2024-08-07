// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"qishan233/hello-wire/application/rpc"
	"qishan233/hello-wire/application/web"
	"qishan233/hello-wire/config"
	"qishan233/hello-wire/controller"
	"qishan233/hello-wire/service"
)

// Injectors from wire.go:

func InitWebServer(config2 config.Config) *web.Server {
	specialRightServiceImpl := service.NewSpecialRightService(config2)
	userService := specialRightServiceImpl.SpecialUserService
	loginController := controller.NewLoginController(userService)
	rightsServiceImpl := &service.RightsServiceImpl{
		UserService: userService,
	}
	rightsController := controller.NewRightsController(rightsServiceImpl)
	server := web.NewServer(loginController, rightsController)
	return server
}

func InitRpcServer(config2 config.Config) *rpc.BusinessService {
	specialRightServiceImpl := service.NewSpecialRightService(config2)
	userService := specialRightServiceImpl.SpecialUserService
	rightsServiceImpl := &service.RightsServiceImpl{
		UserService: userService,
	}
	businessService := rpc.NewBusinessService(userService, rightsServiceImpl)
	return businessService
}

// wire.go:

var Name = "hello-wire"

const Version = "v0.0.1"
