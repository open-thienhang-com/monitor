package internal

import (
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/services/shop/internal/inventory"
	"mono.thienhang.com/services/shop/internal/product"
)

func LoadPlugins() []plugins.Plugin {
	return []plugins.Plugin{
		product.New(),
		inventory.New(),
	}
}
