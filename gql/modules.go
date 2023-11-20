package gql

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRootMutation),
	fx.Provide(NewRootQuery),
)
