package main

import (
	"log/slog"
	"net/http"
	"time"
)

func IPLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("new connection", "ip", r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

func main() {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler: IPLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now().Format(time.RFC3339)
			w.Write([]byte(t))
		})),
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
