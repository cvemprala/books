package main

import (
	"context"
	"fmt"
	logging "github.com/cvemprala/golog"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	logger := logging.NewDefault().WithFields(map[string]interface{}{logging.TagKey: "Books"})
	ctx := logging.WithLogger(context.Background(), logger)

	r := createRouter(ctx)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.WithFields(map[string]interface{}{logging.ErrorKey: err}).Errorln("HTTP server error")
			os.Exit(1)
		}
	}()

	logger.Debugln("Services instantiated. Start listening...")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c

	logger.WithFields(map[string]interface{}{"signal": sig}).Infoln("Got interrupt signal, aborting...")
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	if err := server.Shutdown(timeoutCtx); err != nil {
		logger.Errorln("Timeout when shutting down the server")
		os.Exit(1)
	}
	cancel()
	os.Exit(0)
}

func createRouter(ctx context.Context) http.Handler {
	logger := logging.GetLogger(ctx)
	middlewareChain := func(handler http.Handler) http.Handler {
		return logging.NewMiddleware(handler, logger)
	}

	router := mux.NewRouter()
	router.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	repo := NewRepository()
	getBooks := GetBooksHandler{repo: repo}
	getBook := GetBookHandler{repo: repo}
	createBook := CreateBookHandler{repo: repo}
	updateBook := UpdateBookHandler{repo: repo}
	deleteBook := DeleteBookHandler{repo: repo}

	router.Handle("/books", middlewareChain(getBooks)).Methods("GET")
	router.Handle("/books", middlewareChain(createBook)).Methods("POST")
	router.Handle("/books/{id}", middlewareChain(getBook)).Methods("GET")
	router.Handle("/books/{id}", middlewareChain(updateBook)).Methods("PUT")
	router.Handle("/books/{id}", middlewareChain(deleteBook)).Methods("DELETE")

	return router
}
