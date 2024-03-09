package Backend

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	srv := new(http.Server)
	srv.Addr = ":8080"

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := srv.ListenAndServe()
			if err != nil {
				return err
			}
			zap.L().Info("Starting http server")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
