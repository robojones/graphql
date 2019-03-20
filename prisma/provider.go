package prisma

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewClient,
	NewConfig,
)
