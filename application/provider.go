package application

import (
	"qishan233/hello-wire/application/rpc"
	"qishan233/hello-wire/application/web"

	"github.com/google/wire"
)

var (
	WebProviderSet = wire.NewSet(
		web.NewServer,
	)
	RpcProviderSet = wire.NewSet(
		rpc.NewBusinessService,
	)
)
