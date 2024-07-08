package rpc

import (
	"fmt"
	"qishan233/hello-wire/service"
)

type BusinessService struct {
	UserSvc   service.UserService
	RightsSvc service.RightsService
}

func (r *BusinessService) Run() error {
	fmt.Print("RightsService is running")

	return nil
}

func NewBusinessService(u service.UserService, r service.RightsService) *BusinessService {
	return &BusinessService{
		UserSvc:   u,
		RightsSvc: r,
	}
}
