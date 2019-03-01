package server

import "github.com/google/wire"

var Providers = wire.NewSet(
	NewServer,
	NewConfig,
	NewServeMux,
)
