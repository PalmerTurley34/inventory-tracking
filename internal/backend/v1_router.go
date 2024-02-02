package backend

import "github.com/go-chi/chi"

func newV1Router(cfg *apiConfig) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/users", cfg.createUser)

	return router
}
