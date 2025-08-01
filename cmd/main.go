package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url_shortener/internal/interfaces/repository"
	"url_shortener/internal/interfaces/web"
	"url_shortener/internal/usecases"
	"url_shortener/pkg/safemap"
)

func main() {
	store := safemap.New(1000)
	repo := repository.NewInMemoryRepo(store)
	greetingUC := usecases.NewGreetingUsecase()
	keyValueUC := usecases.NewKeyValueUsecase(repo)
	handler := web.NewHandler(greetingUC, keyValueUC)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("name", handler.Name)
	mux.HandleFunc("add", handler.Add)
	mux.HandleFunc("get", handler.Get)

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
