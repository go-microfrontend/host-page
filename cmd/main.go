package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/go-microfrontend/host-page/internal/templates/layout"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := layout.Page("/catalogue")
		page.Render(r.Context(), w)
	})

	itemsRendererURL, err := url.Parse(os.Getenv("ITEMS_RENDERER_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	itemsRendererProxy := httputil.NewSingleHostReverseProxy(itemsRendererURL)
	http.Handle("/catalogue/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isHXRequest(r) {
			layout.Page(r.URL.Path).Render(r.Context(), w)
			return
		}

		http.StripPrefix("/catalogue", itemsRendererProxy).ServeHTTP(w, r)
	}))

	imagesProviderURL, err := url.Parse(os.Getenv("IMAGES_PROVIDER_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	imagesProviderProxy := httputil.NewSingleHostReverseProxy(imagesProviderURL)
	http.Handle("/images/", http.StripPrefix("/images", imagesProviderProxy))

	authURL, err := url.Parse(os.Getenv("AUTH_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	authProxy := httputil.NewSingleHostReverseProxy(authURL)
	http.HandleFunc("/auth/", func(w http.ResponseWriter, r *http.Request) {
		if !isHXRequest(r) {
			layout.Page(r.URL.Path).Render(r.Context(), w)
			return
		}

		http.StripPrefix("/auth", authProxy).ServeHTTP(w, r)
	})

	http.Handle("/assets/css/",
		http.StripPrefix("/assets/css/",
			http.FileServer(http.Dir("assets/css"))))
	http.Handle("/assets/js/",
		http.StripPrefix("/assets/js/",
			http.FileServer(http.Dir("assets/js"))))

	http.ListenAndServe(":42069", nil)
}

func isHXRequest(r *http.Request) bool {
	return strings.EqualFold(r.Header.Get("HX-Request"), "true")
}
