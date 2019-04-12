package plugin


import (
	"github.com/niftynei/glightning/glightning"
	"os"
)

var lightning *glightning.Lightning


func StartPlugin() *glightning.Plugin {
	plugin := glightning.NewPlugin(initFunc)
	glightning.NewLightning()
	plugin.Start(os.Stdin, os.Stdout)
	return plugin
}

func initFunc(p *glightning.Plugin, o map[string]string, config *glightning.Config) {
	lightning.StartUp(config.RpcFile, config.LightningDir)
}

