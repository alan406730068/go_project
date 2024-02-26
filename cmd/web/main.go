package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alan406730068/go_project/pkg/config"
	"github.com/alan406730068/go_project/pkg/handlers"
	"github.com/alan406730068/go_project/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.Appconfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not get template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
