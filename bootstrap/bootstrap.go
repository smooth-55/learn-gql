package bootstrap

import (
	"context"
	"fmt"

	"github.com/smooth/learn-gql/config"
	"github.com/smooth/learn-gql/infrastructure"
	"github.com/smooth/learn-gql/routes"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.InstalledModule,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	routes routes.Routes,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Printf("Starting Application")
			go func() {
				fmt.Printf("🛠️ setting up routes")
				routes.Setup()
				handler.Gin.LoadHTMLGlob("templates/*")
				_ = handler.Gin.Run(":8010")
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("App stopped")
			return nil
		},
	})
}
