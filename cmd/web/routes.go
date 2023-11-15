package main

import (
	"net/http"

	"github.com/ahmadyogi543/pijarcamp-crud/ui"
	"github.com/julienschmidt/httprouter"
)

func (app *App) routes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/static/*filepath", http.FileServer(http.FS(ui.Files)))
	router.HandlerFunc(http.MethodGet, "/", app.homeHandler)
	router.HandlerFunc(http.MethodGet, "/products", app.getProductsHandler)

	router.HandlerFunc(http.MethodGet, "/products/create", app.createProductGetHandler)
	router.HandlerFunc(http.MethodPost, "/products/create", app.createProductPostHandler)

	router.HandlerFunc(http.MethodGet, "/products/edit/:id", app.editProductGetHandler)
	router.HandlerFunc(http.MethodPut, "/products/edit/:id", app.editProductPutHandler)

	router.HandlerFunc(http.MethodDelete, "/products/delete/:id", app.deleteProductHandler)

	return router
}
