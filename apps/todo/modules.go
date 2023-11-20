package todo

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewQuery),
	fx.Provide(NewMutation),
)
