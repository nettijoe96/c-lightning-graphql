package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-graphql/global"
	"github.com/niftynei/glightning/glightning"
	"github.com/pkg/errors"
)

//feerates
func r_feerates(p graphql.ResolveParams) (interface{}, error) {
	var style FeeRateStyle_ql = FeeRateStyle_ql(p.Args["style"].(string))
	var feeRateEstimate_ql FeeRateEstimate_ql
	var feeRateStyle glightning.FeeRateStyle
	var err error
	l := global.GetGlobalLightning()
	feeRateStyle, err = qlToFeeRateStyle(style)
	if err != nil {
		return nil, err
	}
	feeRateEstimate, err := l.FeeRates(feeRateStyle)
	if err != nil {
		return nil, err
	}
	feeRateEstimate_ql = feeRateEstimateToql(*feeRateEstimate)
	return feeRateEstimate_ql, err
}

//getinfo
func r_getinfo(p graphql.ResolveParams) (interface{}, error) {
        l := global.GetGlobalLightning()
	node, err := l.GetInfo()
	var nodeinfo *NodeInfo_ql = &NodeInfo_ql{}
        nodeToNodeInfo(node, nodeinfo)
        return nodeinfo, err
}

//listinvoices
func r_listinvoices(p graphql.ResolveParams) (interface{}, error) {
        var lstInvoice []glightning.Invoice
	var lstInvoice_ql []Invoice_ql
	var err error
	l := global.GetGlobalLightning()
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

//listnodes
func r_listnodes(p graphql.ResolveParams) (interface{}, error) {
	var lstNode []glightning.Node
        var err error
        l := global.GetGlobalLightning()
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


//listpeers
func r_listpeers(p graphql.ResolveParams) (interface{}, error) {
	var lstPeer []glightning.Peer
	var lstPeer_ql []Peer_ql
	var err error
	l := global.GetGlobalLightning()
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

//pay
func r_pay(p graphql.ResolveParams) (interface{}, error) {
        var paymentSuccess *glightning.PaymentSuccess
	var paymentSuccess_ql PaymentSuccess_ql
        var err error
	l := global.GetGlobalLightning()
	bolt11, isBolt11 := p.Args["bolt11"]
	if !isBolt11 {
		err = errors.New("Cannot find bolt11 in mapping.")
	}else{
		paymentSuccess, err = l.PayBolt(bolt11.(string))
	}
	paymentSuccess_ql = paymentSuccessToql(*paymentSuccess)
	return paymentSuccess_ql, err
}

