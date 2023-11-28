package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("File server program running")

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("/assets/")))

	muxCors := muxCorsMiddleWare(mux)

	http.ListenAndServe(":8080", muxCors)
}

func muxCorsMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "GET POST PUT PATCH DELETE OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.Header().Set("Access-Control-Max-Age", "3600")
		w.Header().Set("Accept", "text/html, application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		//test push out

		next.ServeHTTP(w, r)
	})
}
