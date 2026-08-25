package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gautam1228/ZeroGo/internal/config"
	"github.com/gautam1228/ZeroGo/internal/db"
	"github.com/gautam1228/ZeroGo/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	_, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	log.Printf("Starting server ...")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.Health)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server runnning in %v mode on port: %v", cfg.Env, srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
