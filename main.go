package main

import (
	"log/slog"
	"net/http"
)

func main() {

	router := NewRouter()
	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	slog.Info("Server started on http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Server error", "error", err)
		return
	}
}
