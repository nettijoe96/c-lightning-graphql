package main


import (
	"flag"
	"github.com/nettijoe96/c-lightning-api/plugin"
	"os"
)


func main() {
	var flagMap map[string]interface{} = flags()
	if flagMap["plugin"].(bool) {
	    plugin.Init()
	    p := plugin.GetGlobalPlugin()
	    p.Start(os.Stdin, os.Stdout)
	}else{
	    var api *plugin.StartApi
	    api.Standalone("9043", "graphql", LightningDir)
        }
}


func flags() map[string]interface{} {
	flagMap := make(map[string]interface{})
	var isPlugin *bool = flag.Bool("plugin", true, "is running as a plugin")
	flag.Parse()
	flagMap["plugin"] = *isPlugin
	return flagMap
}
