package theme

import (
	"fmt"
	"html/template"
	"path"
	"path/filepath"
	// "github.com/GoAdminGroup/go-admin/template/types"
)

var dir string = "themes"

func HTML(s string) template.HTML {
	return template.HTML(s)
}

func CSS(s string) template.CSS {
	return template.CSS(s)
}

func JS(s string) template.JS {
	return template.JS(s)
}

func Set(s string) {
	if s != "" {
		dir = s
	}
}

// func Execute(t *template.Template, d interface{}) string {
// 	var resultBuffer strings.Builder
// 	err := t.Execute(&resultBuffer, d)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return resultBuffer.String()
// }

type ExecuteParam struct {
	// User       models.UserModel
	Tmpl     *template.Template
	TmplName string
	// IsPjax     bool
	// Panel      types.Panel
	// Logo       template.HTML
	// Config     *c.Config
	// Menu       *menu.Menu
	// Animation  bool
	// Buttons    types.Buttons
	NoCompress bool
	Iframe     bool
}

var templates *template.Template

func LoadTemplates(s string) *template.Template {
	// loadTemplates(s)
	// rootDir := path.Join(dir, s)

	// templatePattern := filepath.Join(rootDir, "admin/*.html")
	// templates = template.Must(template.ParseGlob(templatePattern))

	// return templates

	return loadTemplates(s)
}

func loadTemplates(templatesDir string) *template.Template {
	fmt.Println("loadTemplates")

	rootDir := path.Join(dir, templatesDir)
	// templatePattern := filepath.Join(rootDir, "/layouts/*.html")

	// layouts, err := filepath.Glob(templatePattern)
	// if err != nil {
	// 	panic(err.Error())
	// }

	templatePattern2 := filepath.Join(rootDir, "/*.html")
	includes, err := filepath.Glob(templatePattern2)
	if err != nil {
		panic(err.Error())
	}

	// // Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		fmt.Println(include)
		// 	layoutCopy := make([]string, len(layouts))
		// 	copy(layoutCopy, layouts)
		// 	files := append(layoutCopy, include)
		// 	// r.AddFromFiles(filepath.Base(include), files...)
		// fmt.Println(files)
		// 	// templates = template.Must(template.ParseGlob(templatePattern))

	}

	templates = template.Must(template.ParseFiles(includes...))

	return templates
}

func Load(name string) *template.Template {
	tmpl := template.New(name)
	// tmplFile := path.Join(dir, s)
	// tmpl, err := template.ExecuteTemplate(s)
	// if err != nil {
	// 	panic(err)
	// }
	return tmpl
}

// func Execute(param *ExecuteParam) *bytes.Buffer {

// 	buf := new(bytes.Buffer)
// 	err := param.Tmpl.ExecuteTemplate(buf, param.TmplName,
// 		types.NewPage(&types.NewPageParam{
// 			// User:       param.User,
// 			// Menu:       param.Menu,
// 			// Assets:     GetComponentAssetImportHTML(),
// 			// Buttons:    param.Buttons,
// 			// Iframe:     param.Iframe,
// 			// UpdateMenu: param.IsPjax,
// 			// Panel: param.Panel.
// 			// 	GetContent(append([]bool{param.Config.IsProductionEnvironment() && !param.NoCompress},
// 			// 		param.Animation)...).AddJS(param.Menu.GetUpdateJS(param.IsPjax)).
// 			// 	AddJS(updateNavAndLogoJS(param.Logo)).AddJS(updateNavJS(param.IsPjax)),
// 			// TmplHeadHTML: Default().GetHeadHTML(),
// 			// TmplFootJS:   Default().GetFootJS(),
// 			// Logo:         param.Logo,
// 		}))
// 	if err != nil {
// 		logger.Error("template execute error", err)
// 	}
// 	return buf
// }

// func Render() *bytes.Buffer {
// 	// Render a specific template
// 	buf := new(bytes.Buffer)
// 	err := templates.ExecuteTemplate(buf, "index.html", nil)
// 	if err != nil {
// 		// http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return nil
// 	}
// 	return buf
// }
