package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/pusher/pusher-http-go"
)

type AuthenticatingObjectRequest struct {
	SocketID string `json:"socket_id"`
}

type MessagePostRequest struct {
	Message string `json:"message"`
}

type MessagePostResponse struct {
	Message string `json:"message"`
}

type AuthenticatingObjectResponse struct {
	Message string `json:"message"`
	ID      string `json:"id"`
	Token   string `json:"token"`
}

func main() {
	port := "8888"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	pusherClient := pusher.Client{
		AppID:   "1486927",
		Key:     "279a52598e5317c2f0f9",
		Secret:  "42c6d2c7e740bc91aeb2",
		Cluster: "ap1",
		Secure:  true,
	}

	r.Post("/messages", func(w http.ResponseWriter, r *http.Request) {
		mRq := MessagePostRequest{}
		err := json.NewDecoder(r.Body).Decode(&mRq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		mRs := MessagePostResponse(mRq)
		pusherClient.Trigger("private-my-channel", "my-event", mRs)
		w.WriteHeader(http.StatusOK)
	})

	r.Post("/pusher/auth", func(w http.ResponseWriter, r *http.Request) {
		params, _ := io.ReadAll(r.Body)

		auth, err := pusherClient.AuthenticatePrivateChannel(params)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Content-Type", "application/json")
		w.Write(auth)
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}
