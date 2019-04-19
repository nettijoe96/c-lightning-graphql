package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-api/lightning"
	"github.com/niftynei/glightning/glightning"
	"strconv"
	"errors"
)

func r_getinfo(p graphql.ResolveParams) (interface{}, error) {
        l := lightning.GetGlobalLightning()
	node, err := l.GetInfo()
	var nodeinfo *NodeInfo_ql = &NodeInfo_ql{}
        nodeToNodeInfo(node, nodeinfo)
        return nodeinfo, err
}

func r_listnodes(p graphql.ResolveParams) (interface{}, error) {
	var lstNode []glightning.Node
        var err error
        l := lightning.GetGlobalLightning()
	id, idPassed := p.Args["id"];
	if !idPassed {
		errors.New("Cannot find id in mapping.")
	} else if id.(string) == "" {
                lstNode, err = l.ListNodes()
	}else{
		lstNode, err = l.GetNode(id.(string))
	}
        return lstNode, err
}

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


