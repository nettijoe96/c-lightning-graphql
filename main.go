package main


import (
	"github.com/nettijoe96/c-lightning-api/plugin"
	"github.com/nettijoe96/c-lightning-api/lightning"
	"github.com/nettijoe96/c-lightning-api/schema"
	"github.com/graphql-go/handler"
	"net/http"
)


func main() {
	l := lightning.GetGlobalLightning()
	l.StartUp("lightning-rpc", LightningDir)
        s := schema.BuildSchema()
	h := handler.New(&handler.Config{
		Schema: &s,
		Pretty: true,
		GraphiQL: true,
	})
        http.Handle("/graphql", h)
	http.ListenAndServe(":10000", nil)
	plugin := plugin.StartPlugin()
	_ = plugin
}

