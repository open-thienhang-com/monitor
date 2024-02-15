package base

import (
	"mono.thienhang.com/pkg/config"
	db "mono.thienhang.com/pkg/database"
	adapters "mono.thienhang.com/pkg/framework"
	plugins "mono.thienhang.com/pkg/plugins"
)

// Adapter is a base adapter contains some helper functions.
type Adapter struct {
	db db.Connection
}

// SetConnection set the db connection.
func (base *Adapter) SetConnection(conn db.Connection) {
	base.db = conn
}

// GetConnection get the db connection.
func (base *Adapter) GetConnection() db.Connection {
	return base.db
}

// HTMLContentType return the default content type header.
func (*Adapter) HTMLContentType() string {
	return "text/html; charset=utf-8"
}

// CookieKey return the cookie key.
// func (*Adapter) CookieKey() string {
// 	return auth.DefaultCookieKey
// }

// GetUser is a helper function get the auth user model from the context.
// func (*Adapter) GetUser(ctx interface{}) (models.UserModel, bool) {
// 	cookie, err := wf.SetContext(ctx).GetCookie()

// 	if err != nil {
// 		return models.UserModel{}, false
// 	}

// 	user, exist := auth.GetCurUser(cookie, wf.GetConnection())
// 	return user.ReleaseConn(), exist
// }

// GetUse is a helper function adds the plugins to the framework.
func (*Adapter) GetUse(app interface{}, plugin []plugins.Plugin, wf adapters.WebFrameWork) error {
	if err := wf.SetApp(app); err != nil {
		return err
	}

	for _, plug := range plugin {
		for path, handlers := range plug.GetHandler() {
			if plug.Prefix() == "" {
				wf.AddHandler(path.Method, path.URL, handlers)
			} else {
				wf.AddHandler(path.Method, config.Url("/"+plug.Prefix()+path.URL), handlers)
			}
		}
	}

	return nil
}

func (*Adapter) Run() error         { panic("not implement") }
func (*Adapter) DisableLog()        { panic("not implement") }
func (*Adapter) Static(_, _ string) { panic("not implement") }

// GetContent is a helper function of adapter.Content
// func (base *Adapter) GetContent(ctx interface{}, getPanelFn types.GetPanelFn, wf adapters.WebFrameWork,
// 	navButtons types.Buttons, fn context.NodeProcessor) {

// 	var (
// 		newBase          = wf.SetContext(ctx)
// 		cookie, hasError = newBase.GetCookie()
// 	)

// 	if hasError != nil || cookie == "" {
// 		newBase.Redirect()
// 		return
// 	}

// 	user, authSuccess := auth.GetCurUser(cookie, wf.GetConnection())

// 	if !authSuccess {
// 		newBase.Redirect()
// 		return
// 	}

// 	var (
// 		panel types.Panel
// 		err   error
// 	)

// 	if !auth.CheckPermissions(user, newBase.Path(), newBase.Method(), newBase.FormParam()) {
// 		panel = template.WarningPanel(errors.NoPermission, template.NoPermission403Page)
// 	} else {
// 		panel, err = getPanelFn(ctx)
// 		if err != nil {
// 			panel = template.WarningPanel(err.Error())
// 		}
// 	}

// 	fn(panel.Callbacks...)

// 	tmpl, tmplName := template.Default().GetTemplate(newBase.IsPjax())

// 	buf := new(bytes.Buffer)
// 	hasError = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(&types.NewPageParam{
// 		User:         user,
// 		Menu:         menu.GetGlobalMenu(user, wf.GetConnection(), newBase.Lang()).SetActiveClass(config.URLRemovePrefix(newBase.Path())),
// 		Panel:        panel.GetContent(config.IsProductionEnvironment()),
// 		Assets:       template.GetComponentAssetImportHTML(),
// 		Buttons:      navButtons.CheckPermission(user),
// 		TmplHeadHTML: template.Default().GetHeadHTML(),
// 		TmplFootJS:   template.Default().GetFootJS(),
// 		Iframe:       newBase.Query().Get(constant.IframeKey) == "true",
// 	}))

// 	if hasError != nil {
// 		logger.Error(fmt.Sprintf("error: %s adapter content, ", newBase.Name()), hasError)
// 	}

// 	newBase.SetContentType()
// 	newBase.Write(buf.Bytes())
// }
