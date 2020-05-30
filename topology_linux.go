// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package ghw

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jaypipes/ghw/pkg/context"
)

func topologyFillInfo(ctx *context.Context, info *TopologyInfo) error {
	info.Nodes = topologyNodes(ctx)
	if len(info.Nodes) == 1 {
		info.Architecture = ARCHITECTURE_SMP
	} else {
		info.Architecture = ARCHITECTURE_NUMA
	}
	return nil
}

// TopologyNodes has been deprecated in 0.2. Please use the TopologyInfo.Nodes
// attribute.
// TODO(jaypipes): Remove in 1.0.
func TopologyNodes() ([]*TopologyNode, error) {
	msg := `
The TopologyNodes() function has been DEPRECATED and will be removed in the 1.0
release of ghw. Please use the TopologyInfo.Nodes attribute.
`
	warn(msg)
	ctx := context.FromEnv()
	return topologyNodes(ctx), nil
}

func topologyNodes(ctx *context.Context) []*TopologyNode {
	nodes := make([]*TopologyNode, 0)

	files, err := ioutil.ReadDir(pathSysDevicesSystemNode(ctx))
	if err != nil {
		warn("failed to determine nodes: %s\n", err)
		return nodes
	}
	for _, file := range files {
		filename := file.Name()
		if !strings.HasPrefix(filename, "node") {
			continue
		}
		node := &TopologyNode{}
		nodeID, err := strconv.Atoi(filename[4:])
		if err != nil {
			warn("failed to determine node ID: %s\n", err)
			return nodes
		}
		node.ID = nodeID
		cores, err := coresForNode(ctx, nodeID)
		if err != nil {
			warn("failed to determine cores for node: %s\n", err)
			return nodes
		}
		node.Cores = cores
		caches, err := cachesForNode(ctx, nodeID)
		if err != nil {
			warn("failed to determine caches for node: %s\n", err)
			return nodes
		}
		node.Caches = caches
		nodes = append(nodes, node)
	}
	return nodes
}
