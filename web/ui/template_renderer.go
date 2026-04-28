package ui

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"sync"
)

type Template struct {
	Mutex       sync.RWMutex
	Env         string
	TemplateDir string
}

func NewTemplate() *Template {
	return &Template{
		Mutex:       sync.RWMutex{},
		TemplateDir: "./templates",
		Env:         "DEV",
	}
}

func (tmpl *Template) RenderTemplate(writer http.ResponseWriter, templateName string, data any) {
	tmpl.Mutex.RLock()
	defer tmpl.Mutex.RUnlock()

	templateFile := path.Join(tmpl.TemplateDir, templateName)
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Fprintf(writer, "Error parsing file: %+v", err)
		return
	}
	err = t.Execute(writer, data)
	if err != nil {
		fmt.Fprintf(writer, "Error executing template: %v", err)
	}
}
