package main

import (
	Backend "backend"

	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(Backend.NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
