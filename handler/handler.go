package handler

import (
	"fmt"
	"net/http"
)

// HomeHandler おうちを建てます
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome My Home!")
}

// ProductsHandler Productをよしなにします
func ProductsHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome My Products!")
}

// ArticlesHandler 論文
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome My Articles!")
}
