package mutation

import (
	"github.com/google/wire"
	"github.com/robojones/graphql/api/resolver/mutation/auth"
	"github.com/robojones/graphql/api/resolver/mutation/project"
)

var Provider = wire.NewSet(
	New,
	auth.New,
	project.New,
)
