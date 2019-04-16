package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-api/lightning"
	"github.com/niftynei/glightning/glightning"
	"strconv"
)

func r_getinfo(p graphql.ResolveParams) (interface{}, error) {
        l := lightning.GetGlobalLightning()
	node, err := l.GetInfo()
	var nodeinfo *NodeInfo_ql = &NodeInfo_ql{}
        nodeToNodeInfo(node, nodeinfo)
        return nodeinfo, err
}


func r_listnodes(p graphql.ResolveParams) (interface{}, error) {
        l := lightning.GetGlobalLightning()
	lstNode, err := l.ListNodes()
	//var lstNodeInfo_ql []NodeInfo_ql //TODO: need sizeof
	//TODO: set each element to a list with following properties as below
	//for i, n := range lstNodeInfo {
	//        var nodeinfo_ql *NodeInfo_ql = &NodeInfo_ql{}
	//	lstNodeInfo_ql = append(lstNodeInfo_ql, n)
        //}

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


