package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"github.com/y-yu/kindle-clock-go/inject"
	"log/slog"
	"net/http"
)

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(error); ok {
					// スタックトレースを含むエラーを生成
					stackTrace := errors.WithStack(e)
					slog.Error("Recovered", "stacktrace", stackTrace)
				}
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

const port = 8080

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(errorMiddleware)

	ctx := context.Background()
	healthHandler := inject.HealthHandler(ctx)
	roomInfoHandler := inject.RoomInfoHandler(ctx)
	clockHandler := inject.ClockHandler(ctx)

	r.Get("/health", healthHandler.Handle)
	r.Get("/", roomInfoHandler.Handle)
	r.Get("/clock", clockHandler.Handle)

	slog.Info("Server started!", "port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		panic(err)
	}
}
