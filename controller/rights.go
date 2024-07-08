package controller

import "qishan233/hello-wire/service"

type RightsController struct {
	rightSvc service.RightsService
}

func NewRightsController(r service.RightsService) *RightsController {
	return &RightsController{
		rightSvc: r,
	}
}
