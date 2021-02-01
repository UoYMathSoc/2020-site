package views

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/UoYMathSoc/2020-site/utils"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
	"time"
)

var (
	ViewsDir    = "views"
	LayoutDir   = "partials"
	ElementsDir = "elements"
	FileExt     = ".gohtml"
)

func New(layout string, content string, elements ...string) *View {
	filePaths := []string{filepath.Join(ViewsDir, content+FileExt)}
	for _, element := range elements {
		filepath := filepath.Join(ViewsDir, ElementsDir, element+FileExt)
		filePaths = append(filePaths, filepath)
	}
	filePaths = append(layoutFiles(), filePaths...)

	return &View{
		Layout:    layout,
		filePaths: filePaths,
		err:       nil,
	}
}

type View struct {
	Template  *template.Template
	Layout    string
	filePaths []string
	err       error
	init      sync.Once
}

func (v *View) Render(w http.ResponseWriter, context structs.PageContext, data interface{}) error {
	v.init.Do(func() {
		v.Template = template.New(v.Layout + FileExt)
		v.Template.Funcs(template.FuncMap{
			"url":       func(s string) string { return utils.PrefixURL(s, context.URLPrefix) },
			"html":      utils.RenderHTML,
			"MonthYear": func(t time.Time) string { return fmt.Sprintf("%s %d", t.Month().String(), t.Year()) },
		})
		v.Template, v.err = v.Template.ParseFiles(v.filePaths...)
	})

	if v.err != nil {
		return v.err
	}

	td := structs.Globals{
		PageContext: context,
		PageData:    data,
	}

	return v.Template.ExecuteTemplate(w, v.Layout+FileExt, td)
}

func (v *View) parse() {
	t := template.New(v.Layout + FileExt)
	t, v.err = template.ParseFiles(v.filePaths...)
	v.Template = t
}

func layoutFiles() []string {
	wildcard := "*" + FileExt
	files, err := filepath.Glob(filepath.Join(ViewsDir, LayoutDir, wildcard))

	if err != nil {
		panic(err)
	}
	return files
}
