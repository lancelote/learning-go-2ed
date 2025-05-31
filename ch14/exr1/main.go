package main

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

func Timeout(ms int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(ms)*time.Millisecond)
			defer cancelFunc()

			req = req.WithContext(ctx)
			h.ServeHTTP(rw, req)
		})
	}
}

func main() {
	timeout := Timeout(100)
	server := http.Server{
		Handler: timeout(http.HandlerFunc(sleepy)),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}

func sleepy(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	message, err := doThing(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			rw.WriteHeader(http.StatusGatewayTimeout)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		rw.WriteHeader(http.StatusOK)
	}

	rw.Write([]byte(message))
}

func doThing(ctx context.Context) (string, error) {
	wait := rand.Intn(200)
	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		return "Done!", nil
	case <-ctx.Done():
		return "Too slow!", ctx.Err()
	}
}
