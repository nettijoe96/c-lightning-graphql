package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-api/lightning"
	"strconv"
)

func r_getinfo(p graphql.ResolveParams) (interface{}, error) {
        l := lightning.GetGlobalLightning()
	r, err := l.GetInfo()
	var nodeinfo *NodeInfo = &NodeInfo{}
	nodeinfo.Id = r.Id
	nodeinfo.Alias = r.Alias
	nodeinfo.Color = r.Color
	nodeinfo.PeerCount = r.PeerCount
	nodeinfo.PendingChannelCount = r.PendingChannelCount
        nodeinfo.ActiveChannelCount = r.ActiveChannelCount
	nodeinfo.InactiveChannelCount = r.InactiveChannelCount
	nodeinfo.Addresses = r.Addresses
        nodeinfo.Binding = r.Binding
	nodeinfo.Version = r.Version
	nodeinfo.Blockheight = r.Blockheight
	nodeinfo.Network = r.Network
	nodeinfo.FeesCollectedMilliSatoshis = strconv.FormatUint(r.FeesCollectedMilliSatoshis, 10)
        return nodeinfo, err
}

