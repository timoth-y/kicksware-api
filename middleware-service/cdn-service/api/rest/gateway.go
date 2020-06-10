package rest

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func ProvideRoutes(rest RestfulHandler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
	)
	router.Mount("/", restRoutes(rest))
	return router
}

func restRoutes(rest RestfulHandler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get(endpointOf(""), rest.Get)
	r.Get(endpointOf("crop"), rest.GetCropped)
	r.Get(endpointOf("resize"), rest.GetResized)
	r.Get(endpointOf("thumbnail"), rest.GetThumbnail)
	r.Post(endpointOf("thumbnail"), rest.Post)
	return
}

func endpointOf(operation string) string {
	if len(operation) != 0 {
		operation = "/" + operation
	}
	return fmt.Sprintf("%v/{collection}/{filename}", operation)
}