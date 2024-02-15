package plugins

type Plugins []Plugin

var (
	pluginList    = make(Plugins, 0)
	allPluginList = make(Plugins, 0)
)

func (pp Plugins) Add(p Plugin) Plugins {
	if !pp.Exist(p) {
		pp = append(pp, p)
	}
	return pp
}

func (pp Plugins) Exist(p Plugin) bool {
	for _, v := range pp {
		if v.Name() == p.Name() {
			return true
		}
	}
	return false
}

func FindByName(name string) (Plugin, bool) {
	for _, v := range pluginList {
		if v.Name() == name {
			return v, true
		}
	}
	return nil, false
}

func FindByNameAll(name string) (Plugin, bool) {
	for _, v := range allPluginList {
		if v.Name() == name {
			return v, true
		}
	}
	return nil, false
}

func Exist(p Plugin) bool {
	return pluginList.Exist(p)
}

func Add(p Plugin) {
	pluginList = pluginList.Add(p)
}
