package backend

import "github.com/go-chi/chi"

func newV1Router(cfg *apiConfig) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/users", cfg.createUser)

	router.Post("/login", cfg.loginUser)

	router.Get("/inventory_items", cfg.getAllInventoryItems)
	router.Post("/inventory_items", cfg.createInventoryItem)

	return router
}
