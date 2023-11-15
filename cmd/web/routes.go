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

	router.HandlerFunc(http.MethodGet, "/products/create", app.addProductFormHandler)
	router.HandlerFunc(http.MethodGet, "/products/edit/:id", app.editProductFormHandler)

	router.HandlerFunc(http.MethodGet, "/products", app.getProductsHandler)
	router.HandlerFunc(http.MethodPost, "/products", app.createProductHandler)
	router.HandlerFunc(http.MethodPut, "/products/:id", app.updateProductHandler)
	router.HandlerFunc(http.MethodDelete, "/products/:id", app.deleteProductHandler)

	return router
}
