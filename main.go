package main

import (
	"main/internal/subscription"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	storage := subscription.NewInMemStorage()
	service := subscription.NewService(storage)
	handler := subscription.NewHandler(service)

	mux.HandleFunc("GET /subscriptions", handler.ListSubscriptions)
	mux.HandleFunc("PUT /subscribe", handler.Subscribe)

	//user_id передавати через params
	mux.HandleFunc("GET /user_subscriptions", handler.GetUserSubscriptions)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}
