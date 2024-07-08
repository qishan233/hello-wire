//go:build wireinject
// +build wireinject

package application

import (
	"qishan233/hello-wire/application/rpc"
	"qishan233/hello-wire/application/web"
	"qishan233/hello-wire/config"
	"qishan233/hello-wire/controller"
	"qishan233/hello-wire/service"

	"github.com/google/wire"
)

var Name = "hello-wire"

const Version = "v0.0.1"

func InitWebServer(config config.Config) *web.Server {
	wire.Build(
		WebProviderSet,
		controller.ProviderSet,
		service.ProviderSet,
	)
	return nil
}

func InitRpcServer(config config.Config) *rpc.BusinessService {
	wire.Build(
		RpcProviderSet,
		service.ProviderSet,
	)
	return nil
}
