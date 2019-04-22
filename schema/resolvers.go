package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-api/lightning"
	"github.com/niftynei/glightning/glightning"
	"errors"
)


func r_getinfo(p graphql.ResolveParams) (interface{}, error) {
        l := lightning.GetGlobalLightning()
	node, err := l.GetInfo()
	var nodeinfo *NodeInfo_ql = &NodeInfo_ql{}
        nodeToNodeInfo(node, nodeinfo)
        return nodeinfo, err
}


func r_listnodes(p graphql.ResolveParams) (interface{}, error) {
	var lstNode []glightning.Node
        var err error
        l := lightning.GetGlobalLightning()
	id, idPassed := p.Args["id"]
	if !idPassed {
		err = errors.New("Cannot find id in mapping.")
	} else if id.(string) == "" {
                lstNode, err = l.ListNodes()
	}else{
		lstNode, err = l.GetNode(id.(string))
	}
        return lstNode, err
}


func r_listinvoices(p graphql.ResolveParams) (interface{}, error) {
        var lstInvoice []glightning.Invoice
	var lstInvoice_ql []Invoice_ql
	var err error
	l := lightning.GetGlobalLightning()
	label, labelPassed := p.Args["label"]
	if !labelPassed {
		err = errors.New("Cannot find label in mapping.")
	}else if label.(string) == "" {
		lstInvoice, err = l.ListInvoices()
	}else{
		lstInvoice, err = l.GetInvoice(label.(string))
	}
        for _, invoice := range lstInvoice {
		lstInvoice_ql = append(lstInvoice_ql, invoiceToql(invoice))
	}
	return lstInvoice_ql, err
}


func r_listpeers(p graphql.ResolveParams) (interface{}, error) {
	var lstPeer []glightning.Peer
	var lstPeer_ql []Peer_ql
	var err error
	l := lightning.GetGlobalLightning()
	id, idPassed := p.Args["id"]
        level, levelPassed := p.Args["level"]
	var loglevel glightning.LogLevel
	switch level {
        case "":
	        loglevel = glightning.None
	case "info":
		loglevel = glightning.Info
	case "unusual":
		loglevel = glightning.Unusual
	case "debug":
		loglevel = glightning.Debug
	case "io":
		loglevel = glightning.Io
	}
	if !idPassed {
		err = errors.New("Cannot find id in mapping.")
	} else if !levelPassed {
		err = errors.New("Cannot find level in mapping.")
	}else{
		lstPeer, err = l.GetPeer(id.(string), loglevel)
		for _, peer := range lstPeer {
	                lstPeer_ql = append(lstPeer_ql, peerToql(peer))
		}
	}
	return lstPeer_ql, err
}


