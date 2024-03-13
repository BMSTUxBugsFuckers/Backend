package main

import (
	CAndC "candc"
	"candc/configs"
	"candc/internal/handler"
	"candc/internal/repository"
	"candc/internal/usecase"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	// Create global zap
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

	// Create context
	ctx := context.Background()

	app := fx.New(
		fx.Provide(
			zap.NewProduction,
			configs.NewServiceConfig,
			repository.NewRepoConfig,
			repository.NewPostgresDB,
			repository.NewRepo,
			usecase.NewUseCase,
			handler.NewHandler,
			handler.InitRoute,
			CAndC.NewServer,
		),
		fx.Invoke(
			configs.InitConfig,
			CAndC.StartServer,
		),
	)

	err := app.Start(ctx)
	if err != nil {
		zap.L().Error("can't start app")
		return
	}
}
