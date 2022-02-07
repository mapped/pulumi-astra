module github.com/mapped/pulumi-astra/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019194827-62530c6537a4
	github.com/mitchellh/mapstructure v1.4.1 => github.com/mitchellh/mapstructure v1.1.2
)

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.13.0
	github.com/pulumi/pulumi/sdk/v3 v3.19.0
	github.com/vavsab/terraform-provider-astra/v2 v2.0.7
)
