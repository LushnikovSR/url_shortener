package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LushnikovSR/url_shortener/internal/config"
	greetingHandler "github.com/LushnikovSR/url_shortener/internal/handler/greeting"
	linkHandler "github.com/LushnikovSR/url_shortener/internal/handler/link"
	url "github.com/LushnikovSR/url_shortener/internal/repository/url/db"
	"github.com/LushnikovSR/url_shortener/internal/usecase/link"
	"github.com/LushnikovSR/url_shortener/internal/usecase/web"
	"github.com/LushnikovSR/url_shortener/pkg/db"
	//"github.com/LushnikovSR/url_shortener/pkg/memory"
	//"github.com/LushnikovSR/url_shortener/internal/repository/url/memory"
)

func main() {

	//store := memory.New(1000)

	// 1. Создаём connection и закрываем чтобы отпустить под
	conf := config.LoadConfig()
	db := db.MustInitDB(conf)
	db.Close() // нужно чтобы отпускать коннекты, потому что есть пул соединений, чтобы когда у тебя под гасится чтобы он сразу коннект отпускал и новые поды имели доступ к этим коннектам, потому что в какой-то момент может нехватить коннектов если старые не будут отпускать.
	// 2. Прокидываем connection в repository

	// --repository
	repository := url.NewRepository(db)

	// --usecases
	greeting := web.New()
	repo := link.New(repository)

	// --handlers
	greet := greetingHandler.New(greeting)
	link := linkHandler.New(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/", greet.Root)
	mux.HandleFunc("/name", greet.Name)
	mux.HandleFunc("/add", link.Add)
	mux.HandleFunc("/get", link.Get)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		fmt.Println("Server starting on port 8080...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
