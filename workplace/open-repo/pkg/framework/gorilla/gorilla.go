// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package gorilla

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"mono.thienhang.com/pkg/config"
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/framework"
	"mono.thienhang.com/pkg/framework/base"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/utils"
)

// Gorilla structure value is a Gorilla GoAdmin adapters.
type Gorilla struct {
	base.Adapter
	ctx Context
	app *mux.Router
}

func init() {
	framework.Register(new(Gorilla))
}

// type HandlerFunc func(ctx Context) (types.Panel, error)

// SetApp implements the method adapters.SetApp.
func (g *Gorilla) SetApp(app interface{}) error {
	var (
		eng *mux.Router
		ok  bool
	)
	if eng, ok = app.(*mux.Router); !ok {
		return errors.New("gorilla adapters SetApp: wrong parameter")
	}
	g.app = eng
	return nil
}

// AddHandler implements the method adapters.AddHandler.
func (g *Gorilla) AddHandler(method, path string, handlers context.Handlers) {

	reg1 := regexp.MustCompile(":(.*?)/")
	reg2 := regexp.MustCompile(":(.*?)$")

	url := path
	url = reg1.ReplaceAllString(url, "{$1}/")
	url = reg2.ReplaceAllString(url, "{$1}")

	g.app.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewContext(r)
		params := mux.Vars(r)

		for key, param := range params {
			if r.URL.RawQuery == "" {
				r.URL.RawQuery += strings.Replace(key, ":", "", -1) + "=" + param
			} else {
				r.URL.RawQuery += "&" + strings.Replace(key, ":", "", -1) + "=" + param
			}
		}

		ctx.SetHandlers(handlers).Next()
		for key, head := range ctx.Response.Header {
			w.Header().Add(key, head[0])
		}

		if ctx.Response.Body == nil {
			w.WriteHeader(ctx.Response.StatusCode)
			return
		}

		w.WriteHeader(ctx.Response.StatusCode)

		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(ctx.Response.Body)

		_, err := w.Write(buf.Bytes())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}).Methods(strings.ToUpper(method))
}

// Context wraps the Request and Response object of Gorilla.
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

// Name implements the method adapters.Name.
func (g *Gorilla) Name() string {
	return "gorilla"
}

// SetContext implements the method adapters.SetContext.
func (g *Gorilla) SetContext(contextInterface interface{}) framework.WebFrameWork {
	var (
		ctx Context
		ok  bool
	)
	if ctx, ok = contextInterface.(Context); !ok {
		panic("gorilla adapters SetContext: wrong parameter")
	}

	return &Gorilla{ctx: ctx}
}

// Use implements the method Adapter.Use.
func (g *Gorilla) Use(app interface{}, plugs []plugins.Plugin) error {
	return g.GetUse(app, plugs, g)
}

// Redirect implements the method adapters.Redirect.
func (g *Gorilla) Redirect() {
	http.Redirect(g.ctx.Response, g.ctx.Request, config.Url(config.GetLoginUrl()), http.StatusFound)
}

// SetContentType implements the method adapters.SetContentType.
func (g *Gorilla) SetContentType() {
	g.ctx.Response.Header().Set("Content-Type", g.HTMLContentType())
}

// Write implements the method adapters.Write.
func (g *Gorilla) Write(body []byte) {
	_, _ = g.ctx.Response.Write(body)
}

// Path implements the method adapters.Path.
func (g *Gorilla) Path() string {
	return g.ctx.Request.RequestURI
}

// Method implements the method adapters.Method.
func (g *Gorilla) Method() string {
	return g.ctx.Request.Method
}

// FormParam implements the method adapters.FormParam.
func (g *Gorilla) FormParam() url.Values {
	_ = g.ctx.Request.ParseMultipartForm(32 << 20)
	return g.ctx.Request.PostForm
}

// IsPjax implements the method adapters.IsPjax.
func (g *Gorilla) IsPjax() bool {
	return g.ctx.Request.Header.Get(utils.PjaxHeader) == "true"
}
