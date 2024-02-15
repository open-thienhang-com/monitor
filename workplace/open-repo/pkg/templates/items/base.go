package items

import (
	"html/template"
	"sync"
)

type Theme interface {
	Name() string

	// Components

	// layout
	// Col() types.ColAttribute
	// Row() types.RowAttribute

	// form and table
	// Form() types.FormAttribute
	// Table() types.TableAttribute
	// DataTable() types.DataTableAttribute

	// TreeView() types.TreeViewAttribute
	// Tree() types.TreeAttribute
	// Tabs() types.TabsAttribute
	// Alert() types.AlertAttribute
	// Link() types.LinkAttribute

	// Paginator() types.PaginatorAttribute
	// Popup() types.PopupAttribute
	// Box() types.BoxAttribute

	// Label() types.LabelAttribute
	// Image() types.ImgAttribute

	// Button() types.ButtonAttribute

	// Builder methods
	GetTmplList() map[string]string
	GetAssetList() []string
	GetAssetImportHTML(exceptComponents ...string) template.HTML
	GetAsset(string) ([]byte, error)
	GetTemplate(bool) (*template.Template, string)
	GetVersion() string
	GetRequirements() []string
	GetHeadHTML() template.HTML
	GetFootJS() template.HTML
	Get404HTML() template.HTML
	Get500HTML() template.HTML
	Get403HTML() template.HTML
}

var (
	templateMap = make(map[string]Theme)

	templateMu sync.Mutex
	compMu     sync.Mutex
)

func Get(theme string) Theme {
	if temp, ok := templateMap[theme]; ok {
		return temp
	}
	panic("wrong theme name")
}

func Add(name string, temp Theme) {
	templateMu.Lock()
	defer templateMu.Unlock()
	if temp == nil {
		panic("template is nil")
	}
	if _, dup := templateMap[name]; dup {
		panic("add template twice " + name)
	}
	templateMap[name] = temp
}

func Themes() []string {
	names := make([]string, len(templateMap))
	i := 0
	for k := range templateMap {
		names[i] = k
		i++
	}
	return names
}

func Default() Theme {
	if temp, ok := templateMap["default"]; ok {
		return temp
	}
	panic("wrong theme name")
}
