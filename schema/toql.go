package schema


import (
        "github.com/niftynei/glightning/glightning"
	"github.com/pkg/errors"
        "strconv"
)

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
	ql.PaymentFields = paymentFieldsToql(paymentSuccess.PaymentFields)
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


