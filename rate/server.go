package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func startServer(addr string) error {
	http.HandleFunc("/", handleRequest)

	log.Printf("Starting server on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// ランダムな遅延（10ms〜500ms）
	delay := time.Duration(rand.Intn(490)+10) * time.Millisecond
	time.Sleep(delay)

	response := fmt.Sprintf("OK (delayed: %v)", delay)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

	log.Printf("Responded after %v\n", delay)
}
