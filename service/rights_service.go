package service

import "qishan233/hello-wire/model"

type RightsServiceImpl struct {
	UserService UserService
}

type RightsService interface {
	GetCurrentRights(userId string) (*model.Rights, error)
}

func (r *RightsServiceImpl) GetCurrentRights(userId string) (*model.Rights, error) {
	return &model.Rights{}, nil
}

func NewRightsService(userSvc UserService) *RightsServiceImpl {
	return &RightsServiceImpl{
		UserService: userSvc,
	}
}

