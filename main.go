package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/vmateosd/scvmm_provider/scvmm"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: scvmm.Provider})
}
