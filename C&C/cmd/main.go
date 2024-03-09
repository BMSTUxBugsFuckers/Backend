package main

import (
	CAndC "candc"
	"candc/configs"
	"candc/internal/handler"
	"candc/internal/repository"
	"candc/internal/usecase"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			configs.NewServiceConfig,
			repository.NewRepoConfig,
			repository.NewPostgresDB,
			repository.NewRepo,
			usecase.NewUseCase,
			handler.NewHandler,
			handler.InitRoute,
			CAndC.NewServer,
		),
		fx.Invoke(CAndC.StartServer),
	).Run()

}
