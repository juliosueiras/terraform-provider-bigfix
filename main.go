package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-bigfix/bigfix"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		//Call provider
		ProviderFunc: bigfix.Provider})
}
