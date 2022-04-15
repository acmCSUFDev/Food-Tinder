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

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/api"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/fileserver"
	"github.com/diamondburned/listener"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func init() {
	envs, _ := filepath.Glob(".env*")
	if len(envs) > 0 {
		if err := godotenv.Load(envs...); err != nil {
			log.Fatalln("cannot load .env:", err)
		}
	}
}

func main() {
	var (
		dbAddress = os.Getenv("DB_ADDRESS")
		assetPath = os.Getenv("ASSET_PATH")
	)

	var fileServer foodtinder.FileServer
	if assetPath != "" {
		fileServer = fileserver.OnDisk(assetPath)
	} else {
		fileServer = fileserver.InMemory(nil)
	}

	dataStore, err := store.Open(dbAddress, store.Opts{
		FileServer: fileServer,
	})
	if err != nil {
		log.Fatalln("cannot set up data server:", err)
	}

	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Timeout(15 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Mount("/", api.Handler(dataStore))

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
