package main

import (
	"fmt"
	"net/http"
	"time"
)

type responseWriter struct {
	respWriter    http.ResponseWriter
	statusCode    int
	headerWritten bool
}

func (rw *responseWriter) Header() http.Header {
	return rw.respWriter.Header()
}

func (rw *responseWriter) WriteHeader(status int) {
	if rw.headerWritten {
		return
	}
	rw.statusCode = status
	rw.respWriter.WriteHeader(status)
	rw.headerWritten = true
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	if !rw.headerWritten {
		rw.statusCode = 200
		rw.headerWritten = true
	}
	return rw.respWriter.Write(data)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		rw := &responseWriter{
			respWriter: w,
		}
		next.ServeHTTP(rw, r)
		respTime := time.Since(startTime)
		fmt.Printf("%s %s -> %d -> %+v\n", r.Method, r.URL.Path, rw.statusCode, respTime)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if rw, ok := w.(*responseWriter); ok {
					if !rw.headerWritten {
						http.Error(rw.respWriter, "Internal Server Error", http.StatusInternalServerError)
					} else {
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					}
				} else {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}
		}()

		next.ServeHTTP(w, r)
		fmt.Println("Request completed successfully")
	})
}
