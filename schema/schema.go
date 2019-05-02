package schema


import (
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-graphql/auth"
)

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

var paymentSuccessType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "paymentSuccess",
		Fields: graphql.Fields {
			"paymentFields": &graphql.Field {
				Type: graphql.String,
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


func BuildSchema() graphql.Schema {
	queryFields := graphql.Fields {
		"getinfo": &graphql.Field {
			Type:  nodeinfoType,
			Description: "Get my node info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.FundsAuth}
				return auth.AuthWrapper(r_getinfo, authLevels, p)
			},
		},
		"listnodes": &graphql.Field {
			Type: graphql.NewList(nodeType),
			Description: "Get a list of all nodes seen in network though channels and node announcement messages",
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "Id for listnodes query. '' is all nodes.",
				},
			},
			Resolve: r_listnodes,
		},
                "listpeers": &graphql.Field {
			Type:  graphql.NewList(peerType),
			Description: "List peers",
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "Id for listpeers query. '' is all peers.",
				},
				"level": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "choose level of logs",
				},
			},
			Resolve: r_listpeers,
		},
		"listinvoices": &graphql.Field {
			Type: graphql.NewList(invoiceType),
			Description: "List invoices",
			Args: graphql.FieldConfigArgument {
				"label": &graphql.ArgumentConfig {
					Type: graphql.String,
					DefaultValue: "",
					Description: "list invoices. Opional label argument",
				},
			},
			Resolve: r_listinvoices,
		},
	}
	mutationFields := graphql.Fields {
		"pay": &graphql.Field {
			Type: paymentSuccessType,
			Description: "Pay via bolt11 as argument",
			Args: graphql.FieldConfigArgument {
				"bolt11": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String), //non null means that argument is required
					Description: "full bolt11 invoice to pay to",
				},
			},
			Resolve: r_pay,
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}
	mutations := graphql.ObjectConfig{Name: "Mutation", Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(mutations)}
	schema, _ := graphql.NewSchema(schemaConfig)
        return schema
}






