package main

import (
	"encoding/json"
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

func toJSON(now time.Time) string {
	timeOut := struct {
		DayOfWeek  string `json:"day_of_week"`
		DayOfMonth int    `json:"day_of_month"`
		Month      string `json:"month"`
		Year       int    `json:"year"`
		Hour       int    `json:"hour"`
		Minute     int    `json:"minute"`
		Second     int    `json:"second"`
	}{
		DayOfWeek:  now.Weekday().String(),
		DayOfMonth: now.Day(),
		Month:      now.Month().String(),
		Year:       now.Year(),
		Hour:       now.Hour(),
		Minute:     now.Minute(),
		Second:     now.Second(),
	}

	out, _ := json.Marshal(timeOut)
	return string(out)
}

func main() {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler: IPLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			var out string

			if r.Header.Get("Accept") == "application/json" {
				out = toJSON(now)
			} else {
				out = now.Format(time.RFC3339)
			}
			w.Write([]byte(out))
		})),
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
