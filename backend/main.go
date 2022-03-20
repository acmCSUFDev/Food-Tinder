package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/api"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/inmemory"
	"github.com/diamondburned/listener"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	dbURL, err := url.Parse(os.Getenv("DB_ADDRESS"))
	if err != nil {
		log.Fatalln("invalid $DB_ADDRESS:", err)
	}

	var dataServer foodtinder.Server

	switch dbURL.Scheme {
	case "mock":
		// Hack because url.Parse confuses some of the path as the domain.
		b, err := os.ReadFile(strings.TrimPrefix(dbURL.String(), "mock://"))
		if err != nil {
			log.Fatalf("cannot read mock database at %s: %v", dbURL.Path, err)
		}

		var state inmemory.State
		if err := json.Unmarshal(b, &state); err != nil {
			log.Fatalln("cannot parse mock database:", err)
		}

		dataServer = inmemory.NewServer(state)

	default:
		log.Fatalf("unsupported database address scheme %q", dbURL.Scheme)
	}

	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Timeout(15 * time.Second))
	r.Mount("/", api.Handler(dataServer))

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
