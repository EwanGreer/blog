package templator

import (
	"fmt"
	"html/template"
)

type Templator struct{}

func New() *Templator {
	return &Templator{}
}

var defaultPaths = []string{
	"static/views/layout.html",
}

func (t *Templator) Page(page string, components ...string) *template.Template {
	paths := append(defaultPaths, fmt.Sprintf("static/views/%s.html", page))

	tmpl, err := template.ParseFiles(loadComponents(paths, components)...)
	if err != nil {
		return nil
	}

	return tmpl
}

func (t *Templator) Partial(name string) *template.Template {
	tmpl, err := template.ParseFiles(fmt.Sprintf("static/components/%s.html", name))
	if err != nil {
		return nil
	}

	return tmpl
}

func loadComponents(templatePaths, components []string) []string {
	for _, comp := range components {
		templatePaths = append(templatePaths, fmt.Sprintf("static/components/%s.html", comp))
	}

	return templatePaths
}
