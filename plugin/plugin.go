package plugin

import (
	"fmt"
	"github.com/graphql-go/handler"
	"github.com/nettijoe96/c-lightning-api/schema"
	"github.com/niftynei/glightning/glightning"
	"github.com/nettijoe96/c-lightning-api/lightning"
	"github.com/niftynei/glightning/jrpc2"
	"log"
	"net/http"
)

var plugin *glightning.Plugin


func GetGlobalPlugin() *glightning.Plugin {
	return plugin
}


func Init() {
	plugin = glightning.NewPlugin(InitFunc)
	plugin.RegisterOption(glightning.NewOption("graphql-port", "port api is available on. default: 9042", "9042"))
	plugin.RegisterOption(glightning.NewOption("graphql-page", "page api is available on. default: graphql", "graphql"))
	rpcStartApi := glightning.NewRpcMethod(&StartApi{}, "run lightning graphql api")
	rpcStartApi.LongDesc = "run lightning graphql api on provided --port (default: 9042) and at --page (default: graphql). Access api at localhost:<port>/<page>/"
	plugin.RegisterMethod(rpcStartApi)
	registerSubscriptions(plugin)
}

func InitFunc(p *glightning.Plugin, o map[string]string, config *glightning.Config) {
	l := lightning.GetGlobalLightning()
	l.StartUp(config.RpcFile, config.LightningDir)
}

func OnConnect(c *glightning.ConnectEvent) {
	log.Printf("connect called: id %s at %s:%d", c.PeerId, c.Address.Addr, c.Address.Port)
}

func OnDisconnect(d *glightning.DisconnectEvent) {
	log.Printf("disconnect called for %s\n", d.PeerId)
}

func registerSubscriptions(p *glightning.Plugin) {
	p.SubscribeConnect(OnConnect)
	p.SubscribeDisconnect(OnDisconnect)
}

type StartApi struct{}

func (api *StartApi) New() interface{} {
	return &StartApi{}
}

func (api *StartApi) Name() string {
	return "graphql"
}

func (api *StartApi) Call() (jrpc2.Result, error) { //TODO: call this method from lightning option and/or command. Check if already running before continuing
	var port string = plugin.GetOptionValue("graphql-port")
	var page string = plugin.GetOptionValue("graphql-page")
        s := schema.BuildSchema()
	h := handler.New(&handler.Config{
		Schema: &s,
		Pretty: true,
		GraphiQL: true,
	})
        http.Handle("/" + page, h)
	go http.ListenAndServe(":" + port, nil)
	return fmt.Sprintf("running api on localhost:" + port + "/" + page + "/"), nil
}




