package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/littledevel/tf-provider/vmprovider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vmprovider.Provider,
	})
}
