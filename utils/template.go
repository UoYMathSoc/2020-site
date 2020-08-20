package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/UoYMathSoc/2020-site/structs"
)

const TemplatePrefix = "views"

var BaseTemplates = []string{
	"partials/base.gohtml",
	"partials/header.gohtml",
}

// Think of better variable name
var OtherTemplates = []string{
	"partials/footer.gohtml",
	"elements/navbar.gohtml",
}

func RenderContent(w http.ResponseWriter, context structs.PageContext, data interface{}, content string) error {
	templates := append(OtherTemplates, content)
	return RenderTemplates(w, context, data, templates...)
}

func RenderTemplates(w http.ResponseWriter, context structs.PageContext, data interface{}, templates ...string) error {
	var err error
	var templatePaths []string
	templates = append(BaseTemplates, templates...)

	for _, template := range templates {
		templatePaths = append(templatePaths, filepath.Join(TemplatePrefix, template))
	}

	t := template.New("base.gohtml")
	t.Funcs(template.FuncMap{
		"url":  func(s string) string { return PrefixURL(s, context.URLPrefix) },
		"html": renderHTML,
	})
	t, err = t.ParseFiles(templatePaths...)
	if err != nil {
		return err
	}

	return t.Execute(w, data)
}

func renderHTML(value interface{}) template.HTML {
	return template.HTML(fmt.Sprint(value))
}
