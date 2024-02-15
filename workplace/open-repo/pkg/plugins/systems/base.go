package systems

import (
	"mono.thienhang.com/pkg/context"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/plugins/base"
)

type System struct {
	base.Base
}

func NewSystem() plugins.Plugin {
	s := &System{
		Base: base.Base{
			App:       context.NewApp(),
			PlugName:  "System Plugin",
			GroupName: "system",
			URLPrefix: "internal",
		},
	}
	return s
}
