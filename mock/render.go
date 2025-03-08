package mock

import (
	"html/template"
	"io"
	"log"
	"strings"
)

func RenderTemplate(w io.Writer, filename string, data interface{}) {
	// load layout and page files
	tmpl := template.Must(template.New("layout").ParseFiles("mock/views/layout.html", "mock/views/pages/"+filename))
	// load component files
	tmpl = template.Must(tmpl.ParseGlob("mock/views/components/*.html"))

	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Fatal("error rendering template", err)
		return
	}
}

func RenderComponent(w io.Writer, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("mock/views/components/" + filename))

	err := tmpl.ExecuteTemplate(w, strings.TrimRight(filename, ".html"), data)
	if err != nil {
		log.Fatal("error component template", err)
		return
	}
}
