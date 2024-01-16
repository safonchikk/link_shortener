package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	handler := &Link{
		Repository: &RedisRepo{
			Client: a.rdb,
		},
	}

	router.Get("/{id}", handler.MakeLong)
	router.Post("/", handler.MakeShort)
	a.router = router
}
