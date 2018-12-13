package view

import (
	"html/template"
	"net/http"
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
		"asset": func(filename string) string {

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
