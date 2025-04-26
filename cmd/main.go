package main

import (
	"net/http"

	"github.com/go-microfrontend/host-page/internal/templates/layout"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := layout.Page()
		page.Render(r.Context(), w)
	})

	http.Handle("/assets/stylesheets/",
		http.StripPrefix("/assets/stylesheets/",
			http.FileServer(http.Dir("assets/stylesheets"))))

	http.ListenAndServe(":42069", nil)
}
