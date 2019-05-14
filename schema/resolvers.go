package schema

import (
	"encoding/json"
	"log"
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-graphql/global"
	"github.com/nettijoe96/glightning/glightning"
	"github.com/pkg/errors"
	"strconv"
)


//close
func r_close(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var ptrCloseResult *glightning.CloseResult
	var id string = p.Args["id"].(string)
	var force bool = p.Args["force"].(bool)
	var timeout uint = uint(p.Args["timeout"].(int))
	l := global.GetGlobalLightning()
	ptrCloseResult, err = l.Close(id, force, timeout)
	if err != nil {
		return nil, errors.Wrap(err, "failed to close channel with id " + id)
	}
	return *ptrCloseResult, err

}


//connect
func r_connect(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var idReturned string
	var id string = p.Args["id"].(string)
	var host string = p.Args["host"].(string)
	var port uint = uint(p.Args["port"].(int))
	l := global.GetGlobalLightning()
	idReturned, err = l.Connect(id, host, port)
	if err != nil {
		err = errors.Wrap(err, "failed to connect to " + host + ":" + string(port))
		return nil, err
	}
	return idReturned, err
}


//decodepay
func r_decodepay(p graphql.ResolveParams) (interface{}, error) {
        var err error
	var bolt11 string = p.Args["bolt11"].(string)
	var description string = p.Args["description"].(string)
	var ptrDecoded *glightning.DecodedBolt11
	var decoded_ql DecodedBolt11_ql
	l := global.GetGlobalLightning()
	ptrDecoded, err = l.DecodePay(bolt11, description)
	if err != nil {
		err = errors.Wrap(err, "failed to decodeBolt11: " + bolt11)
		return nil, err
	}
	decoded_ql = decodedBolt11Toql(*ptrDecoded)
	return decoded_ql, err
}


//delinvoice
func r_delinvoice(p graphql.ResolveParams) (interface{}, error) {
        var err error
	var label string = p.Args["label"].(string)
	var status string = p.Args["status"].(string)
	var ptrInvoice *glightning.Invoice
	var invoice_ql Invoice_ql
	l := global.GetGlobalLightning()
	ptrInvoice, err = l.DeleteInvoice(label, status)
	if err != nil {
		err = errors.Wrap(err, "failed to deleteInvoice: " + label + " with status: " + status)
		return nil, err
	}
	invoice_ql = invoiceToql(*ptrInvoice)
	return invoice_ql, err
}


//disconnect
func r_disconnect(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var id string = p.Args["id"].(string)
	var force bool = p.Args["force"].(bool)
	l := global.GetGlobalLightning()
	err = l.Disconnect(id, force)
	if err != nil {
		err = errors.Wrap(err, "failed to disconnect from: " + id)
		return nil, err
	}
	return nil, err
}


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

//fundchannel
func r_fundchannel(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var ptrFundChannelResult *glightning.FundChannelResult
	var id string = p.Args["id"].(string)
	var amt uint64
	amt, err = strconv.ParseUint(p.Args["satoshi"].(string), 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse satoshi param in fundchannel resolver")
	}
	var announce bool = p.Args["announce"].(bool)
	var satoshi glightning.SatoshiAmount
	satoshi.SendAll = false
	satoshi.Amount = amt
	var feerate string = p.Args["feerate"].(string)
	var feeRate glightning.FeeRate
	var rate int
	var directive glightning.FeeDirective
	if (feerate == "slow" || feerate == "normal" || feerate == "urgent") {
		if feerate == "slow" {
			directive = glightning.Slow
		}else if feerate == "normal" {
			directive = glightning.Normal
		}else{
			directive = glightning.Urgent
		}
		feeRate = glightning.FeeRate {
			Directive: directive,
		}
	}else{
		if len(feerate) <= 5 {
			err = errors.Wrap(err, "perkb or perkw must be the suffix")
			return nil, err
		}else{
			var strStyle string = feerate[len(feerate)-5:]
			var strRate string = feerate[0:len(feerate)-5]
			if rate, err = strconv.Atoi(strRate); err == nil {
	                        var style FeeRateStyle_ql = FeeRateStyle_ql(strStyle)
	                        var feeRateStyle glightning.FeeRateStyle
	                        feeRateStyle, err = qlToFeeRateStyle(style)
				if err != nil {
					return nil, errors.Wrap(err, "valid values for fee rate: slow, normal, urgent, <num>perkb, <num>perkw")
				}
				feeRate = glightning.FeeRate {
					Rate: uint(rate),
					Style: feeRateStyle,
				}
			}else{
			        return nil, errors.Wrap(err, "valid values for fee rate: slow, normal, urgent, <num>perkb, <num>perkw")
			}


		}
	}

	l := global.GetGlobalLightning()
	ptrFundChannelResult, err = l.FundChannelExt(id, &satoshi, &feeRate, announce)
	if err != nil {
		return nil, errors.Wrap(err, "fundchannel call failed")

	}
	return *ptrFundChannelResult, err
}


//getinfo
func r_getinfo(p graphql.ResolveParams) (interface{}, error) {
        l := global.GetGlobalLightning()
	node, err := l.GetInfo()
	var nodeinfo *NodeInfo_ql = &NodeInfo_ql{}
        nodeToNodeInfo(node, nodeinfo)
        return nodeinfo, err
}


//getroute
func r_getroute(p graphql.ResolveParams) (interface{}, error) {
	l := global.GetGlobalLightning()
	var err error
	var id string = p.Args["id"].(string)
	var msatoshi uint64
	msatoshi, err = strconv.ParseUint(p.Args["msatoshis"].(string), 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "msatoshi not parsed in getroute")
	}
	var riskfactor float32 = float32(p.Args["riskfactor"].(float64))
	var cltv uint = uint(p.Args["cltv"].(int))
	var fromid string = p.Args["fromid"].(string)
	var fuzzpercent float32 = float32(p.Args["fuzzpercent"].(float64))
	var exclude []string = p.Args["exclude"].([]string)
	var maxhops int32 = int32(p.Args["maxhops"].(int))
        var hops []glightning.RouteHop
        var hops_ql []RouteHop_ql
	hops, err = l.GetRoute(id, msatoshi, riskfactor, cltv, fromid, fuzzpercent, exclude, maxhops)
	if err != nil {
		return nil, errors.Wrap(err, "failed getRoute")
	}
	for _, h := range hops {
		hops_ql = append(hops_ql, routeHopToql(h))
	}
	return hops_ql, err
}
//getroute^^


//invoice
func r_invoice(p graphql.ResolveParams) (interface{}, error) {
	var invoice_ql Invoice_ql
	var invoice *glightning.Invoice
	var err error
	var msatoshis uint64
	var description string = p.Args["description"].(string)
	var label string = p.Args["label"].(string)
	var expiry uint32 = uint32(p.Args["expiry"].(int))
	var fallbacks []string = p.Args["fallbacks"].([]string)
	var preimage string = p.Args["preimage"].(string)
	var exposeprivatechannels bool = p.Args["exposeprivatechannels"].(bool)
	msatoshis, err = strconv.ParseUint(p.Args["msatoshis"].(string), 10, 64)
	if err != nil {
		err = errors.Wrap(err, "failed to parse uint in r_invoice resolver")
		return nil, err
	}
	l := global.GetGlobalLightning()
	invoice, err = l.CreateInvoice(msatoshis, label, description, expiry, fallbacks, preimage, exposeprivatechannels)
	if err != nil {
		err = errors.Wrap(err, "create invoice failed. Called from r_invoice resolver")
		return nil, err
	}
        invoice_ql = invoiceToql(*invoice)
	invoice_ql.Label = label
	invoice_ql.Description = description
	return invoice_ql, err
}
//invoice ^^


//listchannels
func r_listchannels(p graphql.ResolveParams) (interface{}, error) {
	var scid string = p.Args["scid"].(string)
	var source string = p.Args["source"].(string)
	var channels []glightning.Channel
	var channels_ql []Channel_ql
	var err error
	l := global.GetGlobalLightning()
	if scid == "" && source == "" {
		channels, err = l.ListChannels()
	}else if scid != "" && source == "" {
		channels, err = l.GetChannel(scid)
	}else if source != "" && scid == "" {
		channels, err = l.ListChannelsBySource(source)
	}else {
		err = errors.New("cannot specify both scid and source")
	}
	for _, c := range channels {
		channels_ql = append(channels_ql, channelToql(c))
	}
	return channels_ql, err
}
//listchannels ^^


//listfowards
func r_listforwards(p graphql.ResolveParams) (interface{}, error) {
        var forwarding []glightning.Forwarding
	var forwarding_ql []Forwarding_ql
	var err error
	l := global.GetGlobalLightning()
	forwarding, err = l.ListForwards()
	if err != nil {
		return nil, err
	}
	for _, f := range forwarding {
		forwarding_ql = append(forwarding_ql, forwardingToql(f))
	}
	return forwarding_ql, err
}
//listforwards ^^


//listfunds
func r_listfunds(p graphql.ResolveParams) (interface{}, error) {
        var ptrFundsResult *glightning.FundsResult
	var fundsResult_ql FundsResult_ql
	var err error
	l := global.GetGlobalLightning()
	ptrFundsResult, err = l.ListFunds()
	if err != nil {
		return nil, err
	}
	fundsResult_ql = fundsResultToql(*ptrFundsResult)
	return fundsResult_ql, err
}
//listfunds ^^


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


//listnodes
func r_listpayments(p graphql.ResolveParams) (interface{}, error) {
	var lstPaymentFields []glightning.PaymentFields
	var lstPaymentFields_ql []PaymentFields_ql
        var err error
        l := global.GetGlobalLightning()
	var bolt11 string = p.Args["bolt11"].(string)
	var payment_hash string = p.Args["payment_hash"].(string)
	if bolt11 != "" && payment_hash != "" {
		err = errors.New("Cannot include both bolt11 and payment_hash optional args--only 1 or neither.")
		return nil, err
	}else if bolt11 != "" && payment_hash == "" {
                lstPaymentFields, err = l.ListPayments(bolt11)
	}else if bolt11 == "" && payment_hash != "" {
                lstPaymentFields, err = l.ListPaymentsHash(payment_hash)
	}else {
		lstPaymentFields, err = l.ListPaymentsAll()
	}
	for _, pf := range lstPaymentFields {
		lstPaymentFields_ql = append(lstPaymentFields_ql, paymentFieldsToql(pf))
	}

        return lstPaymentFields, err
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


//sendpay
func r_sendpay(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var ptrSendPayResult *glightning.SendPayResult
	var sendPayResult_ql SendPayResult_ql
        var strRoute string = p.Args["route"].(string)
	var getroute Getroute_ql
	var route glightning.Route
	err = json.Unmarshal([]byte(strRoute), &getroute)
	if err != nil {
		log.Println(err)
	}
	log.Println(getroute)
	var payment_hash string = p.Args["payment_hash"].(string)
	var label string = p.Args["label"].(string)
	var msatoshi uint64
	msatoshi, err = strconv.ParseUint(p.Args["msatoshi"].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	var bolt11 string = p.Args["bolt11"].(string)
	route, err = qlToRoute(getroute)
	if err != nil {
		return nil, err
	}
	l := global.GetGlobalLightning()
	ptrSendPayResult, err = l.SendPay(route.Hops, payment_hash, label, msatoshi, bolt11)
	if err != nil {
		return nil, err
	}
	sendPayResult_ql = sendPayResultToql(*ptrSendPayResult)
	return sendPayResult_ql, err


}


//waitanyinvoice
func r_waitanyinvoice(p graphql.ResolveParams) (interface{}, error) {
	var ptrCompletedInvoice *glightning.CompletedInvoice
	var completedInvoice_ql CompletedInvoice_ql
	var lastpay_index uint = uint(p.Args["lastpay_index"].(int))
        var err error
	l := global.GetGlobalLightning()
	ptrCompletedInvoice, err = l.WaitAnyInvoice(lastpay_index)
	if err != nil {
		err = errors.Wrap(err, "waitanyinvoice failed with lastpay_index param at" + string(lastpay_index))
		return nil, err
	}
	completedInvoice_ql = completedInvoiceToql(*ptrCompletedInvoice)
	return completedInvoice_ql, err
}


//waitinvoice
func r_waitinvoice(p graphql.ResolveParams) (interface{}, error) {
	var ptrCompletedInvoice *glightning.CompletedInvoice
	var completedInvoice_ql CompletedInvoice_ql
	var label string = p.Args["label"].(string)
        var err error
	l := global.GetGlobalLightning()
	ptrCompletedInvoice, err = l.WaitInvoice(label)
	if err != nil {
		err = errors.Wrap(err, "wait invoice failed with label param at" + string(label))
		return nil, err
	}
	completedInvoice_ql = completedInvoiceToql(*ptrCompletedInvoice)
	return completedInvoice_ql, err
}




