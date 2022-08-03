package main

import (
	"context"
	"polycode-provider/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "do-2021.fr/polycode/polycode",
	})
}
