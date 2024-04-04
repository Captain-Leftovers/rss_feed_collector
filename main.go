package main

import (
	"database/sql"

	"github.com/Captain-Leftovers/rss_feed_collector/internal/database"
	_ "github.com/lib/pq"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")


    PORT := os.Getenv("PORT")
    if PORT == "" {
        log.Fatal("PORT environment variable is not set or loaded properly")
    }

    dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal(err)
    }
    
    dbQueries := database.New(db)

    apiCfg := apiConfig{
        DB: dbQueries,
    }

    mux:= http.NewServeMux()
    


    mux.HandleFunc("GET /v1/readiness", handlerReady)

    mux.HandleFunc("GET /v1/err", handlerError)

    mux.HandleFunc("POST /v1/users", apiCfg.handlerUsersCreate)


    corsMux := corsMiddleware(mux)

    server := &http.Server{
        Addr:   ":" + PORT,
        Handler:  corsMux,
    }



    log.Printf("Serving on port: %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}



