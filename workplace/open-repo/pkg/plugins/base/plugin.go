package base

import "mono.thienhang.com/pkg/plugins"

type BasePlugin struct {
	Base
	Info      Info
	IndexURL  string
	Installed bool
}

func (b *BasePlugin) GetInfo() Info       { return b.Info }
func (b *BasePlugin) Name() string        { return b.Info.Name }
func (b *BasePlugin) GetIndexURL() string { return b.IndexURL }
func (b *BasePlugin) IsInstalled() bool   { return b.Installed }

func NewBasePluginWithInfo(info Info) plugins.Plugin {
	return &BasePlugin{Info: info}
}

func NewBasePluginWithInfoAndIndexURL(info Info, u string, installed bool) plugins.Plugin {
	return &BasePlugin{Info: info, IndexURL: u, Installed: installed}
}
