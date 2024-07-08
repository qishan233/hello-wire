package config

import "github.com/google/wire"

var (
	ProductionProviderSet = wire.NewSet(
		NewProduction,
	)
	OfflineProviderSet = wire.NewSet(
		NewOffline,
	)
)
