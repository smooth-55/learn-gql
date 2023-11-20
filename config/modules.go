package config

import (
	"github.com/smooth/learn-gql/apps/todo"
	"github.com/smooth/learn-gql/gql"
	"github.com/smooth/learn-gql/infrastructure"
	"github.com/smooth/learn-gql/routes"
	"go.uber.org/fx"
)

// Module exports dependency
var InstalledModule = fx.Options(
	routes.Module,
	infrastructure.Module,
	gql.Module,
	todo.Module,
)
