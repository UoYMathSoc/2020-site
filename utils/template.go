package utils

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/UoYMathSoc/2020-site/structs"
)

const (
	TemplatePrefix = "views"
	AdminPrefix    = "internal"
	TextPrefix     = "text"
)

var BaseTemplates = []string{
	"partials/header.gohtml",
	"partials/footer.gohtml",
	"elements/navbar.gohtml",
	"partials/base.gohtml",
}

var AdminTemplates = []string{
	"partials/header.gohtml",
	"partials/footer.gohtml",
	"elements/adminbar.gohtml",
	"partials/base.gohtml",
}

func RenderContent(w http.ResponseWriter, context structs.PageContext, data interface{}, content string) error {
	if strings.HasPrefix(content, AdminPrefix+"/") {
		templates := append(AdminTemplates, content)
		return RenderTemplates(w, context, data, templates...)
	}
	templates := append(BaseTemplates, content)
	return RenderTemplates(w, context, data, templates...)
}

func RenderTemplates(w http.ResponseWriter, context structs.PageContext, data interface{}, templates ...string) error {
	var err error

	td := structs.Globals{
		PageContext: context,
		PageData:    data,
	}

	var templatePaths []string
	for _, template := range templates {
		templatePaths = append(templatePaths, filepath.Join(TemplatePrefix, template))
	}

	t := template.New("base.gohtml")
	t.Funcs(template.FuncMap{
		"url":       func(s string) string { return PrefixURL(s, context.URLPrefix) },
		"html":      renderHTML,
		"MonthYear": func(t time.Time) string { return fmt.Sprintf("%s %d", t.Month().String(), t.Year()) },
	})
	t, err = t.ParseFiles(templatePaths...)
	if err != nil {
		return err
	}

	return t.Execute(w, td)
}

func renderHTML(value interface{}) template.HTML {
	return template.HTML(fmt.Sprint(value))
}

func RenderICal(w http.ResponseWriter, data interface{}, text string) error {
	templatePath := filepath.Join(TemplatePrefix, TextPrefix, text)

	td := struct {
		Now       time.Time
		EventData interface{}
	}{
		Now:       time.Now(),
		EventData: data,
	}

	t := template.New("ical.tmpl")
	t.Funcs(template.FuncMap{
		"date":     func(date time.Time) string { return date.Format("20060102T") },
		"time":     func(time time.Time) string { return time.Format("150405Z") },
		"time2":    func(time sql.NullTime) string { return time.Time.Format("150405Z") },
		"datetime": func(time time.Time) string { return time.Format("20060102T150405Z") },
	})
	t, err := t.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	return t.Execute(w, td)
}
