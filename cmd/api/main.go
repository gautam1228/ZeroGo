package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{ "message": "Healthy :)" }`))
	})

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
