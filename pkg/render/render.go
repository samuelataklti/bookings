package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/samuelataklti/bookings/pkg/config"
	"github.com/samuelataklti/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplate Sets the Config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func addDefualtData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//get the template cache from the app config
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCash
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("could not get template cache from template cache")
	}

	buf := new(bytes.Buffer)

	td = addDefualtData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser: ", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCashe := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		fmt.Println("error parsing templates ", err)
		return myCashe, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			fmt.Println("error parsing templates ", err)
			return myCashe, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil || matches == nil {
			fmt.Println("error parsing templates ", err)
			return myCashe, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				fmt.Println("error parsing templates ", err)
				return myCashe, err
			}
		}

		myCashe[name] = ts

	}

	return myCashe, nil
}
