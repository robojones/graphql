package api

import (
	"github.com/google/wire"
	"github.com/robojones/graphql/api/directive/can"
	apiRoot "github.com/robojones/graphql/api/directive/root"
	"github.com/robojones/graphql/api/resolver/mutation"
	"github.com/robojones/graphql/api/resolver/mutation/auth"
	"github.com/robojones/graphql/api/resolver/query"
	directiveRoot "github.com/robojones/graphql/api/resolver/root"
)

var Providers = wire.NewSet(
	directiveRoot.New,
	can.New,
	auth.New,
	mutation.New,
	query.New,
	apiRoot.New,
	New,
)
