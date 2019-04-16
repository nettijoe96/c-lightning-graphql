package schema


import (
	"github.com/graphql-go/graphql"
	"github.com/niftynei/glightning/glightning"
)

type NodeInfo_ql struct {
	Id                         string            `json:"id"`
	Alias                      string            `json:"alias"`
	Color                      string            `json:"color"`
	PeerCount                  int               `json:"num_peers"`
	PendingChannelCount        int               `json:"num_pending_channels"`
	ActiveChannelCount         int               `json:"num_active_channels"`
	InactiveChannelCount       int               `json:"num_inactive_channels"`
	Addresses                  []glightning.Address         `json:"address"`
	Binding                    []glightning.AddressInternal `json:"binding"`
	Version                    string            `json:"version"`
	Blockheight                int               `json:"blockheight"`
	Network                    string            `json:"network"`
	FeesCollectedMilliSatoshis string            `json:"msatoshi_fees_collected"` //graphql protocol cannot handle uint64, so turn in into string
}

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


func BuildSchema() graphql.Schema {
	fields := graphql.Fields{
                "getinfo": &graphql.Field {
			Type:  nodeinfoType,
			Description: "Get my node info",
			Resolve: r_getinfo,
		},
		"listnodes": &graphql.Field {
			Type: graphql.NewList(nodeType),
			Description: "Get a list of all nodes seen in network though channels and node announcement messages",
			Resolve: r_listnodes,
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)
	return schema
}






