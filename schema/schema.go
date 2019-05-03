package schema


import (
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-graphql/auth"
)


func BuildSchema() graphql.Schema {
	queryFields := graphql.Fields {
		"feerates": &graphql.Field {
			Type: feeRateEstimateType,
			Description: "Return feerate estimates, either satoshi-per-kw ({style} perkw) or satoshi-per-kb ({style} perkb).",
			Args: graphql.FieldConfigArgument {
				"style": &graphql.ArgumentConfig {
					Type: graphql.String,
					Description: "either perkw for satoshi-per-kw or perkb for satoshi-per-kb",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_feerates, authLevels, p)
			},
		},
		"getinfo": &graphql.Field {
			Type:  nodeinfoType,
			Description: "Get my node info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_getinfo, authLevels, p)
			},
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_listinvoices, authLevels, p)
			},
		},
		"listnodes": &graphql.Field {
			Type: graphql.NewList(nodeType),
			Description: "Get a list of all nodes seen in network though channels and node announcement messages",
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					DefaultValue: "",
					Description: "Id for listnodes query. '' is all nodes.",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_listnodes, authLevels, p)
			},
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.NoAuth}
				return auth.AuthWrapper(r_listpeers, authLevels, p)
			},
		},
	}
	mutationFields := graphql.Fields {
		"connect": &graphql.Field {
			Type: graphql.String,
			Description: "Connect to {id} at {host} (which can end in ':port' if not default). {id} can also be of the form id@host",
			Args: graphql.FieldConfigArgument {
				"id": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "id of peer",
				},
				"host": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
					Description: "address of peer. It can be tor, ipv4 or ipv6",
				},
				"port": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.Int),
					Description: "port of peer",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.Peers}
				return auth.AuthWrapper(r_connect, authLevels, p)
			},
		},
		"pay": &graphql.Field {
			Type: paymentSuccessType,
			Description: "Pay via bolt11 as argument",
			Args: graphql.FieldConfigArgument {
				"bolt11": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String), //non null means that argument is required
					Description: "full bolt11 invoice to pay to",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var authLevels []auth.AuthLevel = []auth.AuthLevel{auth.Funds}
				return auth.AuthWrapper(r_pay, authLevels, p)
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}
	mutations := graphql.ObjectConfig{Name: "Mutation", Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(mutations)}
	schema, _ := graphql.NewSchema(schemaConfig)
        return schema
}




