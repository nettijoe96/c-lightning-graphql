package schema


import (
        "github.com/nettijoe96/glightning/glightning"
	"github.com/pkg/errors"
        "strconv"
)

//decodepay 
func decodedBolt11Toql(d glightning.DecodedBolt11) DecodedBolt11_ql {
        var ql DecodedBolt11_ql
	ql.Currency = d.Currency
	ql.CreatedAt = strconv.FormatUint(d.CreatedAt, 10)
	ql.Expiry = strconv.FormatUint(d.Expiry, 10)
	ql.Payee = d.Payee
	ql.MilliSatoshis = strconv.FormatUint(d.MilliSatoshis, 10)
	ql.Description = d.Description
	ql.DescriptionHash = d.DescriptionHash
	ql.Fallbacks = d.Fallbacks
	var lstRoutes []BoltRoute_ql
	for _, rs := range d.Routes {
	        lstRoutes = make([]BoltRoute_ql, 0)
		for _, r := range rs {
	                lstRoutes = append(lstRoutes, boltRouteToql(r))
		}
		ql.Routes = append(ql.Routes, lstRoutes)
	}
	ql.Extra = d.Extra
	ql.PaymentHash = d.PaymentHash
	ql.Signature = d.Signature
	return ql
}
func boltRouteToql(r glightning.BoltRoute) BoltRoute_ql {
        var ql BoltRoute_ql
	ql.Pubkey = r.Pubkey
	ql.ShortChannelId = r.ShortChannelId
	ql.FeeBaseMilliSatoshis = strconv.FormatUint(r.FeeBaseMilliSatoshis, 10)
	ql.FeeProportionalMillionths = strconv.FormatUint(r.FeeProportionalMillionths, 10)
	ql.CltvExpiryDelta = r.CltvExpiryDelta
	return ql
}
//decodepay ^^

//feerates
func feeRateEstimateToql(feeRateEstimate glightning.FeeRateEstimate) FeeRateEstimate_ql {
	var ql FeeRateEstimate_ql
	ql.Style = feeRateStyleToql(feeRateEstimate.Style)
	ql.Details = feeRateEstimate.Details
	var onchainEstimate_ql OnchainEstimate_ql = onchainEstimateToql(*feeRateEstimate.OnchainEstimate)
	ql.OnchainEstimate = &onchainEstimate_ql
	ql.Warning = feeRateEstimate.Warning
	return ql
}
func qlToFeeRateStyle(feeRateStyle_ql FeeRateStyle_ql) (glightning.FeeRateStyle, error) {
	var feeRateStyle glightning.FeeRateStyle
	var err error
	if feeRateStyle_ql == SatPerKiloByte {
		feeRateStyle = 0
	}else if feeRateStyle_ql == SatPerKiloSipa {
		feeRateStyle = 1
	}else{
		err = errors.New("fee rate style must be perkb or perkw")
	}
	return feeRateStyle, err
}
func feeRateStyleToql(feeRateStyle glightning.FeeRateStyle) FeeRateStyle_ql {
	var ql FeeRateStyle_ql
	if feeRateStyle == 0 {
		ql = SatPerKiloByte
	}else{
		ql = SatPerKiloSipa
	}
	return ql
}
func onchainEstimateToql(onchainEstimate glightning.OnchainEstimate) OnchainEstimate_ql {
	var ql OnchainEstimate_ql
	ql.OpeningChannelSatoshis = strconv.FormatUint(onchainEstimate.OpeningChannelSatoshis, 10)
	ql.MutualCloseSatoshis = strconv.FormatUint(onchainEstimate.MutualCloseSatoshis, 10)
	ql.UnilateralCloseSatoshis = strconv.FormatUint(onchainEstimate.UnilateralCloseSatoshis, 10)
	return ql
}
//feerates ^^


//getinfo
func nodeToNodeInfo(nodeinfo *glightning.NodeInfo, nodeinfo_ql *NodeInfo_ql) {
	nodeinfo_ql.Id = nodeinfo.Id
	nodeinfo_ql.Alias = nodeinfo.Alias
	nodeinfo_ql.Color = nodeinfo.Color
	nodeinfo_ql.PeerCount = nodeinfo.PeerCount
	nodeinfo_ql.PendingChannelCount = nodeinfo.PendingChannelCount
        nodeinfo_ql.ActiveChannelCount = nodeinfo.ActiveChannelCount
	nodeinfo_ql.InactiveChannelCount = nodeinfo.InactiveChannelCount
	nodeinfo_ql.Addresses = nodeinfo.Addresses
        nodeinfo_ql.Binding = nodeinfo.Binding
	nodeinfo_ql.Version = nodeinfo.Version
	nodeinfo_ql.Blockheight = nodeinfo.Blockheight
	nodeinfo_ql.Network = nodeinfo.Network
	nodeinfo_ql.FeesCollectedMilliSatoshis = strconv.FormatUint(nodeinfo.FeesCollectedMilliSatoshis, 10)
}
//getinfo ^^

//getroute
func routeHopToql(routeHop glightning.RouteHop) RouteHop_ql {
	var ql RouteHop_ql
	ql.Id = routeHop.Id
	ql.ShortChannelId = routeHop.ShortChannelId
	ql.MilliSatoshi = strconv.FormatUint(routeHop.MilliSatoshi, 10)
	ql.Delay = routeHop.Delay
	return ql
}
//getroute ^^


//listchannels
func channelToql(channel glightning.Channel) Channel_ql {
	var ql Channel_ql
	ql.Source = channel.Source
	ql.Destination = channel.Destination
	ql.ShortChannelId = channel.ShortChannelId
	ql.IsPublic = channel.IsPublic
	ql.Satoshis = strconv.FormatUint(channel.Satoshis, 10)
	ql.MessageFlags = channel.MessageFlags
	ql.ChannelFlags = channel.ChannelFlags
	ql.IsActive = channel.IsActive
	ql.LastUpdate = channel.LastUpdate
	ql.BaseFeeMillisatoshi = strconv.FormatUint(channel.BaseFeeMillisatoshi, 10)
	ql.FeePerMillionth = strconv.FormatUint(channel.FeePerMillionth, 10)
	ql.Delay = channel.Delay
	return ql
}
//listchannels ^^


//listfunds
func fundsResultToql(fr glightning.FundsResult) FundsResult_ql {
        var ql FundsResult_ql
	for _, output := range fr.Outputs {
		var fo FundOutput_ql = fundOutputToql(*output)
		ql.Outputs = append(ql.Outputs, &fo)
	}
	for _, channel := range fr.Channels {
		var fc FundingChannel_ql = fundingChannelToql(*channel)
		ql.Channels = append(ql.Channels, &fc)
	}
	return ql
}
func fundOutputToql(fo glightning.FundOutput) FundOutput_ql {
	var ql FundOutput_ql
	ql.TxId = fo.TxId
	ql.Output = fo.Output
	ql.Value = strconv.FormatUint(fo.Value, 10)
	ql.Address = fo.Address
	ql.Status = fo.Status
	return ql
}
func fundingChannelToql(fc glightning.FundingChannel) FundingChannel_ql {
	var ql FundingChannel_ql
	ql.Id = fc.Id
	ql.ShortChannelId = fc.ShortChannelId
	ql.ChannelSatoshi = strconv.FormatUint(fc.ChannelSatoshi, 10)
	ql.ChannelTotalSatoshi = strconv.FormatUint(fc.ChannelTotalSatoshi, 10)
	ql.FundingTxId = fc.FundingTxId
	return ql
}
//listfunds


//listinvoices
func invoiceToql(invoice glightning.Invoice) Invoice_ql {
	var ql Invoice_ql
	ql.PaymentHash = invoice.PaymentHash
	ql.ExpiresAt = strconv.FormatUint(invoice.ExpiresAt, 10)
	ql.Bolt11 = invoice.Bolt11
	ql.WarningOffline = invoice.WarningOffline
	ql.WarningCapacity = invoice.WarningCapacity
	ql.Label = invoice.Label
	ql.Status = invoice.Status
	ql.Description = invoice.Description
	return ql
}
//listinvoices ^^


//listpeers
func peerToql(peer glightning.Peer) Peer_ql {
        var ql Peer_ql
	ql.Id = peer.Id
	ql.Connected = peer.Connected
	ql.NetAddresses = peer.NetAddresses
	ql.GlobalFeatures = peer.GlobalFeatures
	ql.LocalFeatures = peer.LocalFeatures
	for _, channel := range peer.Channels {
	        ql.Channels = append(ql.Channels, peerChannelToql(channel))
	}
	ql.Logs = peer.Logs
	return ql
}
func peerChannelToql(peerChannel glightning.PeerChannel) PeerChannel_ql {
	var ql PeerChannel_ql
	ql.State = peerChannel.State
	ql.ScratchTxId = peerChannel.ScratchTxId
	ql.Owner = peerChannel.Owner
	ql.ShortChannelId = peerChannel.ShortChannelId
	ql.ChannelDirection = peerChannel.ChannelDirection
	ql.ChannelId = peerChannel.ChannelId
	ql.FundingTxId = peerChannel.FundingTxId
	ql.Funding = peerChannel.Funding
	ql.Status = peerChannel.Status
	ql.Private = peerChannel.Private
	for key, val := range peerChannel.FundingAllocations {
		ql.FundingAllocations = append(ql.FundingAllocations, FundingAllocations_ql{key, strconv.FormatUint(val, 10)})
	}
	ql.MilliSatoshiToUs = strconv.FormatUint(peerChannel.MilliSatoshiToUs, 10)
	ql.MilliSatoshiToUsMin = strconv.FormatUint(peerChannel.MilliSatoshiToUsMin, 10)
	ql.MilliSatoshiToUsMax = strconv.FormatUint(peerChannel.MilliSatoshiToUsMax, 10)
	ql.MilliSatoshiTotal = strconv.FormatUint(peerChannel.MilliSatoshiTotal, 10)
	ql.DustLimitSatoshi = strconv.FormatUint(peerChannel.DustLimitSatoshi, 10)
	ql.MaxHtlcValueInFlightMilliSatoshi = strconv.FormatUint(peerChannel.MaxHtlcValueInFlightMilliSatoshi, 10)
	ql.TheirChannelReserveSatoshi = strconv.FormatUint(peerChannel.TheirChannelReserveSatoshi, 10)
	ql.OurChannelReserveSatoshi = strconv.FormatUint(peerChannel.OurChannelReserveSatoshi, 10)
	ql.SpendableMilliSatoshi = strconv.FormatUint(peerChannel.SpendableMilliSatoshi, 10)
	ql.HtlcMinMilliSatoshi = strconv.FormatUint(peerChannel.HtlcMinMilliSatoshi, 10)
	ql.TheirToSelfDelay = peerChannel.TheirToSelfDelay
	ql.OurToSelfDelay = peerChannel.OurToSelfDelay
	ql.MaxAcceptedHtlcs = peerChannel.MaxAcceptedHtlcs
	ql.InPaymentsOffered = strconv.FormatUint(peerChannel.InPaymentsOffered, 10)
	ql.InMilliSatoshiOffered = strconv.FormatUint(peerChannel.InMilliSatoshiOffered, 10)
	ql.InPaymentsFulfilled = strconv.FormatUint(peerChannel.InPaymentsFulfilled, 10)
	ql.InMilliSatoshiFulfilled = strconv.FormatUint(peerChannel.InMilliSatoshiFulfilled, 10)
	ql.OutPaymentsOffered = strconv.FormatUint(peerChannel.OutPaymentsOffered, 10)
	ql.OutMilliSatoshiOffered = strconv.FormatUint(peerChannel.OutMilliSatoshiOffered, 10)
	ql.OutPaymentsFulfilled = strconv.FormatUint(peerChannel.OutPaymentsFulfilled, 10)
	ql.OutMilliSatoshiFulfilled = strconv.FormatUint(peerChannel.OutMilliSatoshiFulfilled, 10)
	for _, htlc := range peerChannel.Htlcs {
		ql.Htlcs = append(ql.Htlcs, htlcToql(htlc))
	}
	return ql
}
func htlcToql(htlc *glightning.Htlc) Htlc_ql {
	var ql Htlc_ql
	ql.Direction = htlc.Direction
	ql.Id = strconv.FormatUint(htlc.Id, 10)
	ql.MilliSatoshi = strconv.FormatUint(htlc.MilliSatoshi, 10)
	ql.Expiry = strconv.FormatUint(htlc.Expiry, 10)
	ql.PaymentHash = htlc.PaymentHash
	ql.State = htlc.State
	return ql
}
//listpeers ^^



//pay
func paymentSuccessToql(paymentSuccess glightning.PaymentSuccess) PaymentSuccess_ql {
        var ql PaymentSuccess_ql
	ql.Id = strconv.FormatUint(paymentSuccess.Id, 10)
	ql.PaymentHash = paymentSuccess.PaymentHash
	ql.Destination = paymentSuccess.Destination
	ql.MilliSatoshi = strconv.FormatUint(paymentSuccess.MilliSatoshi, 10)
	ql.MilliSatoshiSent = strconv.FormatUint(paymentSuccess.MilliSatoshiSent, 10)
	ql.CreatedAt = strconv.FormatUint(paymentSuccess.CreatedAt, 10)
	ql.Status = paymentSuccess.Status
	ql.PaymentPreimage = paymentSuccess.PaymentPreimage
	ql.Description = paymentSuccess.Description
	ql.GetRouteTries = paymentSuccess.GetRouteTries
	ql.SendPayTries = paymentSuccess.SendPayTries
	for _, routeHop := range paymentSuccess.Route {
	        ql.Route = append(ql.Route, routeHopToql(routeHop))
	}
	for _, failure := range paymentSuccess.Failures {
		ql.Failures = append(ql.Failures, payFailureToql(failure))
	}
        return ql
}
func paymentFieldsToql(paymentFields glightning.PaymentFields) PaymentFields_ql {
	var ql PaymentFields_ql
	ql.Id = strconv.FormatUint(paymentFields.Id, 10)
	ql.PaymentHash = paymentFields.PaymentHash
	ql.Destination = paymentFields.Destination
	ql.MilliSatoshi = strconv.FormatUint(paymentFields.MilliSatoshi, 10)
	ql.MilliSatoshiSent = strconv.FormatUint(paymentFields.MilliSatoshiSent, 10)
	ql.CreatedAt = strconv.FormatUint(paymentFields.CreatedAt, 10)
	ql.Status = paymentFields.Status
	ql.PaymentPreimage = paymentFields.PaymentPreimage
	ql.Description = paymentFields.Description
	return ql
}
func payFailureToql(payFailure glightning.PayFailure) PayFailure_ql {
	var ql PayFailure_ql
	ql.Message = payFailure.Message
	ql.Type = payFailure.Type
	ql.OnionReply = payFailure.OnionReply
	ql.ErringIndex = payFailure.ErringIndex
	ql.FailCode = payFailure.FailCode
	ql.ErringNode = payFailure.ErringNode
	ql.ErringChannel = payFailure.ErringChannel
	ql.ChannelUpdate = payFailure.ChannelUpdate
	for _, routeHop := range payFailure.Route {
		ql.Route = append(ql.Route, routeHopToql(routeHop))
	}
        return ql
}
//pay ^^


//sendpay
func sendPayResultToql(r glightning.SendPayResult) SendPayResult_ql {
	var ql SendPayResult_ql
	ql.Message = r.Message
	ql.Id = strconv.FormatUint(r.Id, 10)
	ql.PaymentHash = r.PaymentHash
	ql.Destination = r.Destination
	ql.MilliSatoshi = strconv.FormatUint(r.MilliSatoshi, 10)
	ql.MilliSatoshiSent = strconv.FormatUint(r.MilliSatoshiSent, 10)
	ql.CreatedAt = strconv.FormatUint(r.CreatedAt, 10)
	ql.Status = r.Status
	ql.PaymentPreimage = r.PaymentPreimage
	ql.Description = r.Description
	return ql
}
func qlToRoute(ql Getroute_ql) (glightning.Route, error) {
	var route glightning.Route
	var h glightning.RouteHop
	var err error
	for _, hop := range ql.Getroute {
	        h, err = qlToRouteHop(hop)
		if err != nil {
			return route, err
		}
		route.Hops = append(route.Hops, h)
	}
	return route, err
}
func qlToRouteHop(ql RouteHop_ql) (glightning.RouteHop, error) {
	var err error
	var hop glightning.RouteHop
	hop.Id = ql.Id
	hop.ShortChannelId = ql.ShortChannelId
	hop.MilliSatoshi, err = strconv.ParseUint(ql.MilliSatoshi, 10, 64)
	if err != nil {
		return hop, err
	}
	hop.Delay = ql.Delay
	return hop, err
}


//waitanyinvoice
func completedInvoiceToql(i glightning.CompletedInvoice) CompletedInvoice_ql {
        var ql CompletedInvoice_ql
	ql.Label = i.Label
	ql.Bolt11 = i.Bolt11
	ql.PaymentHash = i.PaymentHash
	ql.Status = i.Status
	ql.Description = i.Description
	ql.PayIndex = i.PayIndex
	ql.MilliSatoshi = strconv.FormatUint(i.MilliSatoshi, 10)
	ql.MilliSatoshiReceived = strconv.FormatUint(i.MilliSatoshiReceived, 10)
	ql.PaidAt = strconv.FormatUint(i.PaidAt, 10)
	ql.ExpiresAt = strconv.FormatUint(i.ExpiresAt, 10)
	return ql
}
//waitanyinvoice ^^

