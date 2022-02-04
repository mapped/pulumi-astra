// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package astra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessList(ctx *pulumi.Context, args *LookupAccessListArgs, opts ...pulumi.InvokeOption) (*LookupAccessListResult, error) {
	var rv LookupAccessListResult
	err := ctx.Invoke("index:astra/getAccessList:getAccessList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessList.
type LookupAccessListArgs struct {
	DatabaseId string `pulumi:"databaseId"`
}

// A collection of values returned by getAccessList.
type LookupAccessListResult struct {
	DatabaseId string `pulumi:"databaseId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string                `pulumi:"id"`
	Results []GetAccessListResult `pulumi:"results"`
}

func LookupAccessListOutput(ctx *pulumi.Context, args LookupAccessListOutputArgs, opts ...pulumi.InvokeOption) LookupAccessListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessListResult, error) {
			args := v.(LookupAccessListArgs)
			r, err := LookupAccessList(ctx, &args, opts...)
			return *r, err
		}).(LookupAccessListResultOutput)
}

// A collection of arguments for invoking getAccessList.
type LookupAccessListOutputArgs struct {
	DatabaseId pulumi.StringInput `pulumi:"databaseId"`
}

func (LookupAccessListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessListArgs)(nil)).Elem()
}

// A collection of values returned by getAccessList.
type LookupAccessListResultOutput struct{ *pulumi.OutputState }

func (LookupAccessListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessListResult)(nil)).Elem()
}

func (o LookupAccessListResultOutput) ToLookupAccessListResultOutput() LookupAccessListResultOutput {
	return o
}

func (o LookupAccessListResultOutput) ToLookupAccessListResultOutputWithContext(ctx context.Context) LookupAccessListResultOutput {
	return o
}

func (o LookupAccessListResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessListResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccessListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessListResultOutput) Results() GetAccessListResultArrayOutput {
	return o.ApplyT(func(v LookupAccessListResult) []GetAccessListResult { return v.Results }).(GetAccessListResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessListResultOutput{})
}
