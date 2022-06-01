package shim

import (
	"github.com/datastax/terraform-provider-astra/v2/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider(version string) *schema.Provider {
	return provider.New(version)()
}
