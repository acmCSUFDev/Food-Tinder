package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/diamondburned/listener"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func init() {
	envs, _ := filepath.Glob(".env*")

	if err := godotenv.Load(envs...); err != nil {
		log.Fatalln("cannot load .env:", err)
	}
}

func main() {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Get("/", hello)

	// SIGINT handler in a cancellable context.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	server := &http.Server{
		Addr:    os.Getenv("HTTP_ADDRESS"),
		Handler: r,
	}

	log.Println("Listening at", server.Addr)
	listener.MustHTTPListenAndServeCtx(ctx, server)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, 世界")
}
