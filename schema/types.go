package schema


import (
	"github.com/graphql-go/graphql"
)


//close 
var closeResultType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CloseResult",
		Fields: graphql.Fields {
			"tx": &graphql.Field {
				Type: graphql.String,
			},
			"txid": &graphql.Field {
				Type: graphql.String,
			},
			"type": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
//close ^^


//decodepay
var decodedBolt11Type = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DecodedBolt11",
		Fields: graphql.Fields {
			"currency": &graphql.Field {
				Type: graphql.String,
			},
			"createAt": &graphql.Field {
				Type: graphql.String,
			},
			"expiry": &graphql.Field {
				Type: graphql.String,
			},
			"payee": &graphql.Field {
				Type: graphql.String,
			},
			"millisatoshis": &graphql.Field {
				Type: graphql.String,
			},
			"description": &graphql.Field {
				Type: graphql.String,
			},
			"descriptionHash": &graphql.Field {
				Type: graphql.String,
			},
			"minFinalCltvExpiry": &graphql.Field {
				Type: graphql.Int,
			},
			"fallbacks": &graphql.Field {
				Type: graphql.NewList(fallbackType),
			},
			"routes": &graphql.Field {
				Type: graphql.NewList(graphql.NewList(boltRouteType)),
			},
			"extra": &graphql.Field {
				Type: graphql.NewList(boltExtraType),
			},
			"paymentHash": &graphql.Field {
				Type: graphql.String,
			},
			"signature": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
var fallbackType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Fallback",
		Fields: graphql.Fields {
			"type": &graphql.Field {
				Type: graphql.String,
			},
			"address": &graphql.Field {
				Type: graphql.String,
			},
			"hex": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
var boltRouteType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BoltRoute",
		Fields: graphql.Fields {
			"pubkey": &graphql.Field {
				Type: graphql.String,
			},
			"shortChannelId": &graphql.Field {
				Type: graphql.String,
			},
			"feeBaseMilliSatoshis": &graphql.Field {
				Type: graphql.String,
			},
			"feeProportionalMillionths": &graphql.Field {
				Type: graphql.String,
			},
			"cltvExpiryDelta": &graphql.Field {
				Type: graphql.Int,
			},
		},
	},
)
var boltExtraType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BoltExtra",
		Fields: graphql.Fields {
			"tag": &graphql.Field {
				Type: graphql.String,
			},
			"data": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
//decodepay ^^


//feerates
var feeRateEstimateType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FeeRateEstimate",
		Fields: graphql.Fields {
			"style": &graphql.Field {
				Type: graphql.String,
			},
			"details": &graphql.Field {
				Type: feeRateDetailsType,
			},
			"onchainEstimate": &graphql.Field {
				Type: onchainEstimateType,
			},
			"warning": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
var feeRateDetailsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FeeRateDetails",
		Fields: graphql.Fields {
			"urgent": &graphql.Field {
				Type: graphql.Int,
			},
			"normal": &graphql.Field {
				Type: graphql.Int,
			},
			"slow": &graphql.Field {
				Type: graphql.Int,
			},
			"minAcceptable": &graphql.Field {
				Type: graphql.Int,
			},
			"maxAcceptable": &graphql.Field {
				Type: graphql.Int,
			},
		},
	},
)
var onchainEstimateType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "OnchainEstimate",
		Fields: graphql.Fields {
			"openingChannelSatoshis": &graphql.Field {
				Type: graphql.String,
			},
			"mutualCloseSatoshis": &graphql.Field {
				Type: graphql.String,
			},
			"unilateralCloseSatoshis": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
//feerates ^^


//fundchannel 
var fundChannelResultType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FundChannelResult",
		Fields: graphql.Fields {
			"fundingTx": &graphql.Field {
				Type: graphql.String,
			},
			"fundingTxId": &graphql.Field {
				Type: graphql.String,
			},
			"channelId": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
//fundchannel ^^


//getinfo
var nodeinfoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NodeInfo",
		Fields: graphql.Fields {
			"id": &graphql.Field {
				Type: graphql.String,
			},
			"alias": &graphql.Field {
				Type: graphql.String,
			},
			"color": &graphql.Field {
				Type: graphql.String,
			},
			"peerCount": &graphql.Field {
				Type: graphql.Int,
			},
			"pendingChannelCount": &graphql.Field {
				Type: graphql.Int,
			},
			"activeChannelCount": &graphql.Field {
				Type: graphql.Int,
			},
			"InactiveChannelCount": &graphql.Field {
				Type: graphql.Int,
			},
			"addresses": &graphql.Field {
				Type: graphql.NewList(addressType),
			},
			"binding": &graphql.Field {
				Type: graphql.NewList(addressInternalType),
			},
			"blockheight": &graphql.Field {
				Type: graphql.Int,
			},
			"network": &graphql.Field {
				Type: graphql.String,
			},
			"feesCollectedMilliSatoshis": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
//getinfo ^^

//getroute
var routeHopType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "routeHop",
		Fields: graphql.Fields {
			"id": &graphql.Field {
				Type: graphql.String,
			},
			"shortChannelId": &graphql.Field {
				Type: graphql.String,
			},
			"milliSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"delay": &graphql.Field {
				Type: graphql.Int,
			},
		},
	},
)
//getroute ^^


//listchannels 
var channelType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "channel",
		Fields: graphql.Fields {
			"source": &graphql.Field {
				Type: graphql.String,
			},
			"destination": &graphql.Field {
				Type: graphql.String,
			},
			"shortChannelId": &graphql.Field {
				Type: graphql.String,
			},
			"isPublic": &graphql.Field {
				Type: graphql.Boolean,
			},
			"satoshis": &graphql.Field {
				Type: graphql.String,
			},
			"messageFlags": &graphql.Field {
				Type: graphql.Int,
			},
			"channelFlags": &graphql.Field {
				Type: graphql.Int,
			},
			"isActive": &graphql.Field {
				Type: graphql.Boolean,
			},
			"lastUpdate": &graphql.Field {
				Type: graphql.Int,
			},
			"baseFeeMillisatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"feePerMillionth": &graphql.Field {
				Type: graphql.String,
			},
			"delay": &graphql.Field {
				Type: graphql.Int,
			},
		},
	},
)
//listchannels ^^


//listinvoices 
var invoiceType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "invoice",
		Fields: graphql.Fields {
			"paymentHash": &graphql.Field {
				Type: graphql.String,
			},
			"expiresAt": &graphql.Field {
				Type: graphql.String,
			},
			"bolt11": &graphql.Field {
				Type: graphql.String,
			},
			"warningOffline": &graphql.Field {
				Type: graphql.String,
			},
			"warningCapacity": &graphql.Field {
				Type: graphql.String,
			},
			"label": &graphql.Field {
				Type: graphql.String,
			},
			"status": &graphql.Field {
				Type: graphql.String,
			},
			"description": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
//listinvoices ^^


//listnodes
var nodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Node",
		Fields: graphql.Fields {
			"id": &graphql.Field {
				Type: graphql.String,
			},
			"alias": &graphql.Field {
				Type: graphql.String,
			},
			"color": &graphql.Field {
				Type: graphql.String,
			},
			"lastTimestamp": &graphql.Field {
				Type: graphql.NewNonNull(graphql.Int),
			},
			"globalFeatures": &graphql.Field {
				Type: graphql.String,
			},
			"addresses": &graphql.Field {
				Type: graphql.NewList(addressType),
			},
		},
	},
)
var addressType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "address",
		Fields: graphql.Fields {
			"type": &graphql.Field {
				Type: graphql.String,
			},
			"address": &graphql.Field {
				Type: graphql.String,
			},
			"port": &graphql.Field {
				Type: graphql.Int,
			},
		},
	},
)
var addressInternalType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "addressInternal",
		Fields: graphql.Fields {
			"type": &graphql.Field {
				Type: graphql.String,
			},
			"address": &graphql.Field {
				Type: graphql.String,
			},
			"port": &graphql.Field {
				Type: graphql.Int,
			},
			"socket": &graphql.Field {
				Type: graphql.String,
			},
			"service": &graphql.Field {
				Type: addressType,
			},
			"name": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
//listpeers ^^


//listpeers
var peerType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "Peer",
		Fields: graphql.Fields {
			"id": &graphql.Field {
				Type: graphql.String,
			},
			"connected": &graphql.Field {
				Type: graphql.Boolean,
			},
			"netAddresses": &graphql.Field {
				Type: graphql.NewList(graphql.String),
			},
			"globalFeatures": &graphql.Field {
				Type: graphql.String,
			},
			"localFeatures": &graphql.Field {
				Type: graphql.String,
			},
			"channels": &graphql.Field {
				Type: graphql.NewList(peerChannelType),
			},
			"logs": &graphql.Field {
				Type: graphql.NewList(logType),
			},
		},
	},
)
var peerChannelType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "peerChannel",
		Fields: graphql.Fields {
			"state": &graphql.Field {
				Type: graphql.String,
			},
			"scratchTxId": &graphql.Field {
                                Type: graphql.String,
			},
			"owner": &graphql.Field {
				Type: graphql.String,
			},
			"shortChannelId": &graphql.Field {
				Type: graphql.String,
			},
			"channelDirection": &graphql.Field {
				Type: graphql.Int,
			},
			"channelId": &graphql.Field {
				Type: graphql.String,
			},
			"fundingTxId": &graphql.Field {
				Type: graphql.String,
			},
			"funding": &graphql.Field {
				Type: graphql.String,
			},
			"status": &graphql.Field {
				Type: graphql.NewList(graphql.String),
			},
			"private": &graphql.Field {
				Type: graphql.Boolean,
			},
			"fundingAllocations": &graphql.Field {
				Type: graphql.NewList(fundingAllocationType),
			},
			"milliSatoshiToUs": &graphql.Field {
				Type: graphql.String,
			},
			"milliSatoshiToUsMin": &graphql.Field {
				Type: graphql.String,
			},
			"milliSatoshiToUsMax": &graphql.Field {
				Type: graphql.String,
			},
			"milliSatoshiTotal": &graphql.Field {
				Type: graphql.String,
			},
			"dustLimitSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"maxHtlcValueInFlightMilliSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"theirChannelReserveSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"ourChannelReserveSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"spendableMilliSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"htlcMinMilliSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"theirToSelfDelay": &graphql.Field {
				Type: graphql.Int,
			},
			"ourToSelfDelay": &graphql.Field {
				Type: graphql.Int,
			},
			"maxAcceptedHtlcs": &graphql.Field {
				Type: graphql.Int,
			},
			"inPaymentsOffered": &graphql.Field {
				Type: graphql.String,
			},
			"inMilliSatoshiOffered": &graphql.Field {
				Type: graphql.String,
			},
			"inPaymentsFulfilled": &graphql.Field {
				Type: graphql.String,
			},
			"inMilliSatoshiFulfilled": &graphql.Field {
				Type: graphql.String,
			},
			"outPaymentOffered": &graphql.Field {
				Type: graphql.String,
			},
			"outMilliSatoshiOffered": &graphql.Field {
				Type: graphql.String,
			},
			"outPaymentsFulfilled": &graphql.Field {
				Type: graphql.String,
			},
			"outMilliSatoshiFulfilled": &graphql.Field {
				Type: graphql.String,
			},
			"htlcs": &graphql.Field {
				Type: graphql.NewList(htlcType),
			},
		},
	},
)
var fundingAllocationType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "fundingAllocation",
		Fields: graphql.Fields {
			"id": &graphql.Field {
				Type: graphql.String,
			},
			"msat": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
var logType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "log",
		Fields: graphql.Fields {
			"type": &graphql.Field {
				Type: graphql.String,
			},
			"time": &graphql.Field {
				Type: graphql.String,
			},
			"source": &graphql.Field {
				Type: graphql.String,
			},
			"message": &graphql.Field {
				Type: graphql.String,
			},
			"numSkipped": &graphql.Field {
				Type: graphql.Int,
			},
		},
	},
)
var htlcType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "htlc",
		Fields: graphql.Fields {
			"direction": &graphql.Field {
				Type: graphql.String,
			},
			"id": &graphql.Field {
				Type: graphql.String,
			},
			"milliSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"expiry": &graphql.Field {
				Type: graphql.String,
			},
			"paymentHash": &graphql.Field {
				Type: graphql.String,
			},
			"state": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
//listpeers ^^


//pay
var paymentSuccessType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "paymentSuccess",
		Fields: graphql.Fields {
			"paymentFields": &graphql.Field {
				Type: paymentFieldsType,
			},
			"getRouteTries": &graphql.Field {
				Type: graphql.Int,
			},
			"sendPayTries": &graphql.Field {
				Type: graphql.Int,
			},
			"route": &graphql.Field {
				Type: graphql.NewList(routeHopType),
			},
			"failures": &graphql.Field {
				Type: graphql.NewList(payFailureType),
			},
		},
	},
)
var paymentFieldsType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "paymentFields",
		Fields: graphql.Fields {
			"id": &graphql.Field {
				Type: graphql.String,
			},
			"paymentHash": &graphql.Field {
				Type: graphql.String,
			},
			"destination": &graphql.Field {
				Type: graphql.String,
			},
			"milliSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"milliSatoshiSent": &graphql.Field {
				Type: graphql.String,
			},
			"CreatedAt": &graphql.Field {
				Type: graphql.String,
			},
			"Status": &graphql.Field {
				Type: graphql.String,
			},
			"paymentPreimage": &graphql.Field {
				Type: graphql.String,
			},
			"description": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)
var payFailureType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "payFailure",
		Fields: graphql.Fields {
			"message": &graphql.Field {
				Type: graphql.String,
			},
			"type": &graphql.Field {
				Type: graphql.String,
			},
			"onionReply": &graphql.Field {
				Type: graphql.String,
			},
			"erringIndex": &graphql.Field {
				Type: graphql.Int,
			},
			"failCode": &graphql.Field {
				Type: graphql.Int,
			},
			"erringNode": &graphql.Field {
				Type: graphql.String,
			},
			"channelUpdate": &graphql.Field {
				Type: graphql.String,
			},
			"route": &graphql.Field {
				Type: graphql.NewList(routeHopType),
			},
		},
	},
)
//pay ^^


var payRequestType = graphql.NewObject(
	graphql.ObjectConfig {
	        Name: "payRequest",
		Fields: graphql.Fields {
			"bolt11": &graphql.Field {
				Type: graphql.String,
			},
			"milliSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"desc": &graphql.Field {
				Type: graphql.String,
			},
			"RiskFactor": &graphql.Field {
				Type: graphql.Float,
			},
			"maxFeePercent": &graphql.Field {
				Type: graphql.Float,
			},
			"retryFor": &graphql.Field {
				Type: graphql.Int,
			},
			"maxDelay": &graphql.Field {
				Type: graphql.Int,
			},
			"exemptFee": &graphql.Field {
				Type: graphql.Boolean,
			},
		},
	},
)


//waitanyinvoice
var completedInvoiceType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "completedInvoice",
		Fields: graphql.Fields {
			"label": &graphql.Field {
				Type: graphql.String,
			},
			"bolt11": &graphql.Field {
				Type: graphql.String,
			},
			"paymentHash": &graphql.Field {
				Type: graphql.String,
			},
			"status": &graphql.Field {
				Type: graphql.String,
			},
			"description": &graphql.Field {
				Type: graphql.String,
			},
			"payIndex": &graphql.Field {
				Type: graphql.Int,
			},
			"milliSatoshi": &graphql.Field {
				Type: graphql.String,
			},
			"milliSatoshiReceived": &graphql.Field {
				Type: graphql.String,
			},
			"paidAt": &graphql.Field {
				Type: graphql.String,
			},
			"expiresAt": &graphql.Field {
				Type: graphql.String,
			},
		},
	},
)

//waitanyinvoice ^^
