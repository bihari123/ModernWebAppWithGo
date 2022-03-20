package utilities

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"example.com/app1/config"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// func map is a map of functions that we are gonna use in a template
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache

	} else {
		templateCache, _ = CreateTemplateCache()
	}

	template, ok := templateCache[tmpl]

	if !ok {
		log.Fatal("Error finding the template cache")

	}

	buf := new(bytes.Buffer)

	_ = template.Execute(buf, nil)
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to the browser", err)
	}
}

// creates a template cache as a map
func CreateTemplateCache() (myCache map[string]*template.Template, err error) {

	myCache = map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return
	}

	for _, page := range pages {
		name := filepath.Base(page)

		template_set, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			fmt.Println("found some error in preparing  matches, quitting")
			return myCache, err

		}

		if len(matches) > 0 {
			template_set, err = template_set.ParseGlob("./templates/*layout.tmpl")

			if err != nil {
				fmt.Println("err in template set")
				return myCache, err

			}

		}
		myCache[name] = template_set

	}
	return myCache, nil

}
