package main

import (
	"net/http"

	"github.com/ahmadyogi543/pijarcamp-crud/ui"
	"github.com/julienschmidt/httprouter"
)

func (app *App) routes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/static/*filepath", http.FileServer(http.FS(ui.Files)))
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/products", app.getProducts)

	router.HandlerFunc(http.MethodGet, "/create", app.addProductView)
	router.HandlerFunc(http.MethodPost, "/create", app.addProductPost)

	router.HandlerFunc(http.MethodGet, "/edit/:id", app.editProductView)
	router.HandlerFunc(http.MethodPut, "/edit/:id", app.editProductPost)

	router.HandlerFunc(http.MethodDelete, "/delete/:id", app.delete)

	return router
}
