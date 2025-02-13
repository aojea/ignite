package cni

import (
	log "github.com/sirupsen/logrus"
	"github.com/weaveworks/ignite/pkg/network/cni"
	"github.com/weaveworks/ignite/pkg/providers"
)

func SetCNINetworkPlugin() (err error) {
	log.Trace("Initializing the CNI provider...")
	plugin, err := cni.GetCNINetworkPlugin(providers.Runtime)
	if err != nil {
		return err
	}
	providers.NetworkPlugins[plugin.Name()] = plugin
	return
}
