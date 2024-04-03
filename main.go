package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")


    PORT := os.Getenv("PORT")
    if PORT == "" {
        log.Fatal("PORT environment variable is not set or loaded properly")
    }

    mux:= http.NewServeMux()

    mux.HandleFunc("GET /v1/readiness", handlerReady)

    mux.HandleFunc("GET /v1/err", handlerError)


    corsMux := corsMiddleware(mux)

    server := &http.Server{
        Addr:   ":" + PORT,
        Handler:  corsMux,
    }



    log.Printf("Serving on port: %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}