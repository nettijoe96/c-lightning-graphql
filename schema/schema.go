package schema


import (
	"github.com/graphql-go/graphql"
//	"github.com/niftynei/glightning/glightning"
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
		Name: "peer channel",
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
		Name: "funding allocation",
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
		Name: "Log",
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
		Name: "htlc type",
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


func BuildSchema() graphql.Schema {
	fields := graphql.Fields {
		"getinfo": &graphql.Field {
			Type:  nodeinfoType,
			Description: "Get my node info",
			Resolve: r_getinfo,
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
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)
        return schema
}






