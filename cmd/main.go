package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url_shortener/internal/handler"
	"url_shortener/internal/repository/url"
	"url_shortener/internal/usecase"
	db "url_shortener/pkg/db"
)

func main() {
	store := db.New(1000)
	repository := url.NewRepository(store)
	greetUC := usecase.NewGreeting()
	repoUC := usecase.NewRepo(repository)
	greethandler := handler.NewGreetHandler(greetUC)
	repohandler := handler.NewRepoHandler(repoUC)
	mux := http.NewServeMux()
	mux.HandleFunc("/", greethandler.Root)
	mux.HandleFunc("/name", greethandler.Name)
	mux.HandleFunc("/add", repohandler.Add)
	mux.HandleFunc("/get", repohandler.Get)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		fmt.Println("Server starting on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server failed: %v\n", err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 5*time.Second)
	defer shutdownCancel()

	fmt.Println("Shutting down server...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Shutdown error: %v\n", err)
	}
	fmt.Println("Server stopped")
}
