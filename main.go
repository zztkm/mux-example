package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zztkm/mux-example/handler"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/products", handler.ProductsHandler)
	r.HandleFunc("/articles", handler.ArticlesHandler)
	http.Handle("/", r)
	r.Use(loggingMiddleware)

	// server start
	http.ListenAndServe(":8080", r)

}
