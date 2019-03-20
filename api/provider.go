package api

import (
	"github.com/google/wire"
	"github.com/robojones/graphql/api/resolver"
)

var Provider = wire.NewSet(
	resolver.Provider,
	New,
)
