package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ahmadyogi543/pijarcamp-crud/internal/models"
)

type App struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	products *models.ProductModel
}

func main() {
	addr := flag.String("addr", ":5000", "HTTP network address")
	dsn := flag.String("dsn", "root:@/pijarcamp", "MySQL data source name")
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	db, err := openDB("mysql", *dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := App{
		errorLog: errorLog,
		infoLog:  infoLog,
		products: &models.ProductModel{DB: db},
	}

	server := http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	app.infoLog.Printf("Starting server on %s", *addr)
	err = server.ListenAndServe()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}
