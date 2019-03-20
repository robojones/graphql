package resolver

import (
	"github.com/google/wire"
	"github.com/robojones/graphql/api/resolver/mutation"
	"github.com/robojones/graphql/api/resolver/project"
	"github.com/robojones/graphql/api/resolver/query"
	"github.com/robojones/graphql/api/resolver/root"
)

var Provider = wire.NewSet(
	mutation.Provider,
	project.New,
	query.New,
	root.New,
)
