package api

import (
	"github.com/google/wire"
	"github.com/robojones/graphql/api/resolver/mutation"
	"github.com/robojones/graphql/api/resolver/mutation/auth"
	"github.com/robojones/graphql/api/resolver/query"
	"github.com/robojones/graphql/api/resolver/root"
)

var Providers = wire.NewSet(
	auth.New,
	mutation.New,
	query.New,
	root.New,
	New,
)
