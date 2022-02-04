// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package astra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAstraDatabases(ctx *pulumi.Context, args *LookupAstraDatabasesArgs, opts ...pulumi.InvokeOption) (*LookupAstraDatabasesResult, error) {
	var rv LookupAstraDatabasesResult
	err := ctx.Invoke("index:astra/getAstraDatabases:getAstraDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAstraDatabases.
type LookupAstraDatabasesArgs struct {
	CloudProvider *string `pulumi:"cloudProvider"`
	Status        *string `pulumi:"status"`
}

// A collection of values returned by getAstraDatabases.
type LookupAstraDatabasesResult struct {
	CloudProvider *string `pulumi:"cloudProvider"`
	// The provider-assigned unique ID for this managed resource.
	Id      string                    `pulumi:"id"`
	Results []GetAstraDatabasesResult `pulumi:"results"`
	Status  *string                   `pulumi:"status"`
}

func LookupAstraDatabasesOutput(ctx *pulumi.Context, args LookupAstraDatabasesOutputArgs, opts ...pulumi.InvokeOption) LookupAstraDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAstraDatabasesResult, error) {
			args := v.(LookupAstraDatabasesArgs)
			r, err := LookupAstraDatabases(ctx, &args, opts...)
			return *r, err
		}).(LookupAstraDatabasesResultOutput)
}

// A collection of arguments for invoking getAstraDatabases.
type LookupAstraDatabasesOutputArgs struct {
	CloudProvider pulumi.StringPtrInput `pulumi:"cloudProvider"`
	Status        pulumi.StringPtrInput `pulumi:"status"`
}

func (LookupAstraDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAstraDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getAstraDatabases.
type LookupAstraDatabasesResultOutput struct{ *pulumi.OutputState }

func (LookupAstraDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAstraDatabasesResult)(nil)).Elem()
}

func (o LookupAstraDatabasesResultOutput) ToLookupAstraDatabasesResultOutput() LookupAstraDatabasesResultOutput {
	return o
}

func (o LookupAstraDatabasesResultOutput) ToLookupAstraDatabasesResultOutputWithContext(ctx context.Context) LookupAstraDatabasesResultOutput {
	return o
}

func (o LookupAstraDatabasesResultOutput) CloudProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAstraDatabasesResult) *string { return v.CloudProvider }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAstraDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAstraDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAstraDatabasesResultOutput) Results() GetAstraDatabasesResultArrayOutput {
	return o.ApplyT(func(v LookupAstraDatabasesResult) []GetAstraDatabasesResult { return v.Results }).(GetAstraDatabasesResultArrayOutput)
}

func (o LookupAstraDatabasesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAstraDatabasesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAstraDatabasesResultOutput{})
}
