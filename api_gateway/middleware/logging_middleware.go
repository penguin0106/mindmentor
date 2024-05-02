package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware логирует каждый HTTP запрос
func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Вызываем следующий обработчик в цепочке
		next.ServeHTTP(w, r)

		// Логирование
		log.Printf("[%s] %s %s %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	}
}

// WrapHandlerFunc оборачивает функцию типа http.HandlerFunc в http.Handler
func WrapHandlerFunc(handlerFunc http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r)
	})
}
