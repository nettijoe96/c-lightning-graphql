package main


import (
	"flag"
	"github.com/nettijoe96/c-lightning-api/plugin"
	"os"
)


func main() {
	var flagMap map[string]interface{} = standaloneFlags()
	//if plugin=false, then we do not create a plugin! (plugin=true by default)
	if flagMap["plugin"].(bool) {
	    plugin.Init()
	    p := plugin.GetGlobalPlugin()
	    p.Start(os.Stdin, os.Stdout)
	}else{
	    var isTLS bool = flagMap["tls"].(bool)
	    var certfile string = flagMap["certfile"].(string)
	    var keyfile string = flagMap["keyfile"].(string)
	    var api *plugin.StartApi
	    api.Standalone(isTLS, "9043", "graphql", certfile, keyfile, LightningDir)
        }
}


func standaloneFlags() map[string]interface{} {
	/* standalone app flags set here. See plugin/plugin.go for plugin options" */
	var isPlugin *bool = flag.Bool("plugin", true, "is running as a plugin")
	var isTLS *bool = flag.Bool("tls", true, "is running tls")
	var certfile *string = flag.String("certfile", "cert.pem", "is running tls")
	var keyfile *string = flag.String("keyfile", "key.pem", "is running tls")
	flagMap := make(map[string]interface{})

	flag.Parse()
	flagMap["plugin"] = *isPlugin
	flagMap["tls"] = *isTLS
	flagMap["certfile"] = *certfile
	flagMap["keyfile"] = *keyfile

	return flagMap
}
