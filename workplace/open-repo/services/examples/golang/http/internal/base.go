package internal

import (
	"fmt"
	"sync"

	"mono.thienhang.com/pkg/engine"
	"mono.thienhang.com/pkg/plugins"
)

var (
	eng   *engine.Engine
	engMu sync.Mutex
)

func SetEngine(e *engine.Engine) {
	engMu.Lock()
	defer engMu.Unlock()
	if e == nil {
		panic("Empty engine")
	}
	//
	if eng == nil {
		eng = e
	} else {
		fmt.Println("Engine instance is already set.") //TODO: change log lib
	}
}

func LoadPlugins() []plugins.Plugin {
	if eng == nil {
		panic("Please set engine")
	}

	//
	return []plugins.Plugin{}
}
