package view

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

//Config represents the view config
type Config struct {
	AssetURL string
}

//View represents a view struct
type View struct {
	Config Config
}

//New creates a new View struct
func New(c Config) *View {
	return &View{
		Config: c,
	}
}

//Render creates and executes a template with options
func (v View) Render(w http.ResponseWriter, templateName string, vars interface{}) {

	t, err := template.New("react.html").Funcs(template.FuncMap{
		"asset": func(filename string, hasManifest bool) string {

			if hasManifest == true {

				//load JSON manifest
				var data map[string]string
				plan, err := ioutil.ReadFile("frontend/dist/manifest.json")

				if err != nil {
					panic(err.Error())
				}

				err = json.Unmarshal(plan, &data)

				if err != nil {
					panic(err.Error())
				}

				fmt.Println(data)

				stripped := strings.TrimLeft(filename, "/")

				if fl, ok := data[stripped]; ok == true {
					return v.Config.AssetURL + "/" + fl
				}

				return v.Config.AssetURL + filename
			}

			return v.Config.AssetURL + filename
		},
	}).ParseFiles(templateName)

	if err != nil {
		panic(err.Error())
	}

	t.Execute(w, vars)
}

//RenderJSON renders json response
func (v View) RenderJSON(w http.ResponseWriter, data []byte) {

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
