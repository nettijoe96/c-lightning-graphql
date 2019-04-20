package main


import (
	"github.com/nettijoe96/c-lightning-api/plugin"
	"os"
)


func main() {
	plugin.Init()
	p := plugin.GetGlobalPlugin()
	p.Start(os.Stdin, os.Stdout)
}

