package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%v %v request is now processing.", request.Method, request.URL.Path)

		start := time.Now().UTC()
		next.ServeHTTP(writer, request)

		log.Printf("Request to %v %v completed in %v.", request.Method, request.URL.Path, time.Since(start))
	})
}
