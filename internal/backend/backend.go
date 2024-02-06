package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func StartBackendServer() {
	godotenv.Load()

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	if port == "" {
		log.Fatal("PORT not found in environments")
	}
	if dbURL == "" {
		log.Fatal("DB_URL not found in environment")
	}

	database, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbQueries := db.New(database)

	cfg := apiConfig{DB: dbQueries}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))
	router.Get("/healthz", healthcheck)

	v1Router := newV1Router(&cfg)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Println("Server listening on port:", port)
	server.ListenAndServe()
}
