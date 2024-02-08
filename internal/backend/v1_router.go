package backend

import "github.com/go-chi/chi"

func newV1Router(cfg *apiConfig) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/users", cfg.createUser)

	router.Post("/login", cfg.loginUser)

	router.Post("/valid_username", cfg.validateUsername)

	router.Get("/inventory_items", cfg.getAllInventoryItems)
	router.Post("/inventory_items", cfg.createInventoryItem)
	router.Delete("/inventory_items/{ID}", cfg.deleteInventoryItem)

	return router
}
