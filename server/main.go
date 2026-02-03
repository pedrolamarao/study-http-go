package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"purpura.dev.br/study/http/server/service"
)

func main() {
	svc := service.NewService()

	mux := http.NewServeMux()
	mux.HandleFunc("/", svc.Handle)

	srv := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
