package utilities

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// func map is a map of functions that we are gonna use in a template
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	templateCache, err := CreateTemplateCache()
	fmt.Println(templateCache)
	if err != nil {
		fmt.Println("Error making tempalte cache")
		log.Fatal(err)
	}

	template, ok := templateCache[tmpl]

	if !ok {
		fmt.Println("Error finding the template cache")
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = template.Execute(buf, nil)
	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to the browser", err)
	}
}

// creates a template cache as a map
func CreateTemplateCache() (myCache map[string]*template.Template, err error) {

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
