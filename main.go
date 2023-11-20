package main

import (
	"github.com/smooth/learn-gql/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
