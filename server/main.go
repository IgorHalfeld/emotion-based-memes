package main

import (
	"emotion-based-memes/container"
	"emotion-based-memes/handlers"
	"emotion-based-memes/services"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	azureClient := services.NewAzureService()
	faceHandlers := handlers.NewFaceHandler(container.ServiceContainer{
		AzureService: azureClient,
	})
	memeHandlers := handlers.NewMemeHandler(container.ServiceContainer{
		AzureService: azureClient,
	})

	r.Post("/face/analyze", faceHandlers.Analyze)
	r.Post("/meme", memeHandlers.Create)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "temp-images"))
	FileServer(r, "/static", filesDir)

	http.ListenAndServe(":3000", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		log.Println("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
