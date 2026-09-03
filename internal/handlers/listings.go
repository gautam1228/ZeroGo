package handlers

import (
	"database/sql"
	"net/http"
)

func List(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		db.Query("")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{ "message": "All good :)" }`))
	}
}
