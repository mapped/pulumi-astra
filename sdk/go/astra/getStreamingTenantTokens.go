// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getStreamingTenantTokens` provides a datasource that lists streaming tenant tokens.
func GetStreamingTenantTokens(ctx *pulumi.Context, args *GetStreamingTenantTokensArgs, opts ...pulumi.InvokeOption) (*GetStreamingTenantTokensResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetStreamingTenantTokensResult
	err := ctx.Invoke("astra:index/getStreamingTenantTokens:getStreamingTenantTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStreamingTenantTokens.
type GetStreamingTenantTokensArgs struct {
	// Name of the Pulsar Cluster. Format: `pulsar-<cloud provider>-<cloud region>`. Example: `pulsar-gcp-useast1`
	ClusterName string `pulumi:"clusterName"`
	// Name of the streaming tenant for which to fetch tokens.
	TenantName string `pulumi:"tenantName"`
}

// A collection of values returned by getStreamingTenantTokens.
type GetStreamingTenantTokensResult struct {
	// Name of the Pulsar Cluster. Format: `pulsar-<cloud provider>-<cloud region>`. Example: `pulsar-gcp-useast1`
	ClusterName string `pulumi:"clusterName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the streaming tenant for which to fetch tokens.
	TenantName string `pulumi:"tenantName"`
	// The list of tokens for the specified tenant.
	Tokens []GetStreamingTenantTokensToken `pulumi:"tokens"`
}

func GetStreamingTenantTokensOutput(ctx *pulumi.Context, args GetStreamingTenantTokensOutputArgs, opts ...pulumi.InvokeOption) GetStreamingTenantTokensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStreamingTenantTokensResult, error) {
			args := v.(GetStreamingTenantTokensArgs)
			r, err := GetStreamingTenantTokens(ctx, &args, opts...)
			var s GetStreamingTenantTokensResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStreamingTenantTokensResultOutput)
}

// A collection of arguments for invoking getStreamingTenantTokens.
type GetStreamingTenantTokensOutputArgs struct {
	// Name of the Pulsar Cluster. Format: `pulsar-<cloud provider>-<cloud region>`. Example: `pulsar-gcp-useast1`
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Name of the streaming tenant for which to fetch tokens.
	TenantName pulumi.StringInput `pulumi:"tenantName"`
}

func (GetStreamingTenantTokensOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStreamingTenantTokensArgs)(nil)).Elem()
}

// A collection of values returned by getStreamingTenantTokens.
type GetStreamingTenantTokensResultOutput struct{ *pulumi.OutputState }

func (GetStreamingTenantTokensResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStreamingTenantTokensResult)(nil)).Elem()
}

func (o GetStreamingTenantTokensResultOutput) ToGetStreamingTenantTokensResultOutput() GetStreamingTenantTokensResultOutput {
	return o
}

func (o GetStreamingTenantTokensResultOutput) ToGetStreamingTenantTokensResultOutputWithContext(ctx context.Context) GetStreamingTenantTokensResultOutput {
	return o
}

// Name of the Pulsar Cluster. Format: `pulsar-<cloud provider>-<cloud region>`. Example: `pulsar-gcp-useast1`
func (o GetStreamingTenantTokensResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStreamingTenantTokensResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStreamingTenantTokensResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStreamingTenantTokensResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the streaming tenant for which to fetch tokens.
func (o GetStreamingTenantTokensResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStreamingTenantTokensResult) string { return v.TenantName }).(pulumi.StringOutput)
}

// The list of tokens for the specified tenant.
func (o GetStreamingTenantTokensResultOutput) Tokens() GetStreamingTenantTokensTokenArrayOutput {
	return o.ApplyT(func(v GetStreamingTenantTokensResult) []GetStreamingTenantTokensToken { return v.Tokens }).(GetStreamingTenantTokensTokenArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStreamingTenantTokensResultOutput{})
}
