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
	"partials/header.tmpl",
	"partials/footer.tmpl",
	"partials/base.tmpl",
}

func RenderTemplate(w http.ResponseWriter, context structs.PageContext, data interface{}, mainTemplate string) error {
	var err error

	templates := append(BaseTemplates, mainTemplate)

	var templatePaths []string
	for _, template := range templates {
		templatePaths = append(templatePaths, filepath.Join(TemplatePrefix, template))
	}

	t := template.New("base.tmpl")
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
