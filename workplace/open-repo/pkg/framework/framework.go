// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package framework

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins"
)

// WebFrameWork is an interface which is used as an adapter of
// framework and goAdmin. It must implement two methods. Use registers
// the routes and the corresponding handlers. Content writes the
// response to the corresponding context of framework.
type WebFrameWork interface {
	// Name return the web framework name.
	Name() string

	// Use method inject the plugins to the web framework engine which is the
	// first parameter.
	Use(app interface{}, plugins []plugins.Plugin) error

	// Content add the panel html response of the given callback function to
	// the web framework context which is the first parameter.
	// Content(ctx interface{}, fn types.GetPanelFn, navButtons ...types.Button)

	// User get the auth user model from the given web framework context.
	// User(ctx interface{}) (models.UserModel, bool)

	// AddHandler inject the route and handlers of to the web framework.
	AddHandler(method, path string, handlers context.Handlers)

	// DisableLog()

	Static(prefix, path string)

	// Helper functions
	// ================================

	SetApp(app interface{}) error
	// SetConnection(db.Connection)
	// GetConnection() db.Connection
	// SetContext(ctx interface{}) WebFrameWork
	// GetCookie() (string, error)
	// Path() string
	// Method() string
	// FormParam() url.Values
	// IsPjax() bool
	// Redirect()
	// SetContentType()
	// Write(body []byte)
	// CookieKey() string
	// HTMLContentType() string
}
