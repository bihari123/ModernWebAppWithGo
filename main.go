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
	r.App.UseCache =false
}

func main() {
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Starting the application on port number", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
