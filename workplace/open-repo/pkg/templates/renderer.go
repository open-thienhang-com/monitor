package theme

import (
	"html/template"
)

// Renderer type is the Agnostic Renderer for themes.
// When gin is in debug mode then all themes works with
// hot reloading allowing you modify file templates and seeing changes instantly.
// Renderer should be created using theme.NewRenderer() constructor.
type Renderer interface {
	Add(name string, tmpl *template.Template)
	AddFromFiles(name string, files ...string) *template.Template
	AddFromGlob(name, glob string) *template.Template
	AddFromString(name, templateString string) *template.Template
	AddFromStringsFuncs(name string, funcMap template.FuncMap, templateStrings ...string) *template.Template
	AddFromFilesFuncs(name string, funcMap template.FuncMap, files ...string) *template.Template
}
