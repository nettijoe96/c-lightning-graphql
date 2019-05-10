package global

import (
	"github.com/nettijoe96/glightning/glightning"
)

var lightning *glightning.Lightning = glightning.NewLightning()
var plugin *glightning.Plugin


func GetGlobalLightning() *glightning.Lightning {
	return lightning
}

func GetGlobalPlugin() *glightning.Plugin {
	return plugin
}

func SetGlobalPlugin(p *glightning.Plugin) {
	plugin = p
}

func SetGlobalLightning(l *glightning.Lightning) {
	lightning = l
}

