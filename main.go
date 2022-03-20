package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/app1/config"
	"example.com/app1/handlers"
	"example.com/app1/utilities"
)

const portNumber = ":8080"

var appConfig config.AppConfig

func init() {

	tc, err := utilities.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache")
	}
	appConfig.TemplateCache = tc
	utilities.NewTemplates(&appConfig)

	r := handlers.NewRepo(&appConfig)
	handlers.NewHandler(r)
	r.App.UseCache = false
}

func main() {

	server := http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfig),
	}
	
	fmt.Println("Starting the application on port number", portNumber)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}
