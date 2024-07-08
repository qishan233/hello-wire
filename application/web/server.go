package web

import (
	"fmt"
	"qishan233/hello-wire/controller"
)

type Server struct {
	LoginController  *controller.LoginController
	RightsController *controller.RightsController
}

func NewServer(loginController *controller.LoginController, rightsController *controller.RightsController) *Server {
	return &Server{
		LoginController:  loginController,
		RightsController: rightsController,
	}
}

func (s *Server) Run() error {
	fmt.Println("Server is running")
	return nil
}
