package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ahmadyogi543/pijarcamp-crud/internal/models"
	"github.com/ahmadyogi543/pijarcamp-crud/ui/components"
	"github.com/julienschmidt/httprouter"
)

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	products, err := app.products.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}
	components.HomePage(products).Render(w)
}

func (app *App) addProductFormHandler(w http.ResponseWriter, r *http.Request) {
	components.AddProductForm().Render(w)
}

func (app *App) editProductFormHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	product, err := app.products.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	components.EditProductForm(product).Render(w)
}

func (app *App) getProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := app.products.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}
	components.ProductItemList(products).Render(w)
}

func (app *App) createProductHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("name")
	description := r.PostForm.Get("desc")
	price, err := strconv.Atoi(r.PostForm.Get("price"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	qty, err := strconv.Atoi(r.PostForm.Get("qty"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	_, err = app.products.Create(name, description, price, qty)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func (app *App) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	err = r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("name")
	description := r.PostForm.Get("desc")
	price, err := strconv.Atoi(r.PostForm.Get("price"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	qty, err := strconv.Atoi(r.PostForm.Get("qty"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	err = app.products.Update(&models.Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Qty:         qty,
	})
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			app.clientError(w, http.StatusBadRequest)
		} else {
			app.serverError(w, err)
		}
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func (app *App) deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	err = app.products.Delete(id)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
