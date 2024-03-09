package C_C

import (
	"candc/configs"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Server struct {
	httpServer *fiber.App
}

func NewServer(handlers *fiber.App) *Server {
	return &Server{
		httpServer: handlers,
	}
}

func StartServer(lifecycle fx.Lifecycle, s *Server, conf configs.ServiceConfig) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				return s.httpServer.Listen(conf.Addr)
			},
			OnStop: func(ctx context.Context) error {
				return s.httpServer.ShutdownWithContext(ctx)
			},
		},
	)
}
