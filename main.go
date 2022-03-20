package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/app1/config"
	"example.com/app1/handlers"
	"example.com/app1/utilities"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

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

	// change this to true when in production

	appConfig.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction //setting false for local testing , true for https connection

	appConfig.Session = session

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
