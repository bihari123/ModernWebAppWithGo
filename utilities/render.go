package utilities

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// func map is a map of functions that we are gonna use in a template
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTempate, err := template.ParseFiles("./templates/" + tmpl)

	if err != nil {
		fmt.Println("\n\n\n")
		log.Fatalln("[ERROR] Error parsing the template: ", err)
		fmt.Println("\n\n\n")
	}
	err = parsedTempate.Execute(w, nil)

	if err != nil {
		fmt.Println("\n\n\n")
		log.Fatalln("[ERROR] Error parsing the template: ", err)
		fmt.Println("\n\n\n")
	}

}

func RenderTemplateUsingBaseTemplate(w http.ResponseWriter, tmpl string) (myCache map[string]*template.Template, err error) {

	myCache = map[string]*template.Template{}

	pages, err := filepath.Glob("./tempates/*.page.tmpl")

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
			return myCache, err

		}

		if len(matches) > 0 {
			template_set, err = template_set.ParseGlob("./templates/*layout.tmpl")

			if err != nil {
				return myCache, err

			}

		}

		myCache[name] = template_set

	}
	return myCache, nil

}
