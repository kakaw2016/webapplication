package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kakaw2016/webapplication/pkg/config"

	"github.com/kakaw2016/webapplication/pkg/handlers"
	"github.com/kakaw2016/webapplication/pkg/render"
)

const portNumber = ":8080"

// main is the main application function

func main() {

	var app config.AppConfig

	session := scs.New()

	session.lifetime = 24 * time.Hour

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{

		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
