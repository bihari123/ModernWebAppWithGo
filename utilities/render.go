package utilities

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

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
