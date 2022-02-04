// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package astra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoles(ctx *pulumi.Context, args *LookupRolesArgs, opts ...pulumi.InvokeOption) (*LookupRolesResult, error) {
	var rv LookupRolesResult
	err := ctx.Invoke("index:astra/getRoles:getRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoles.
type LookupRolesArgs struct {
	RoleId string `pulumi:"roleId"`
}

// A collection of values returned by getRoles.
type LookupRolesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string           `pulumi:"id"`
	Results []GetRolesResult `pulumi:"results"`
	RoleId  string           `pulumi:"roleId"`
}

func LookupRolesOutput(ctx *pulumi.Context, args LookupRolesOutputArgs, opts ...pulumi.InvokeOption) LookupRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRolesResult, error) {
			args := v.(LookupRolesArgs)
			r, err := LookupRoles(ctx, &args, opts...)
			return *r, err
		}).(LookupRolesResultOutput)
}

// A collection of arguments for invoking getRoles.
type LookupRolesOutputArgs struct {
	RoleId pulumi.StringInput `pulumi:"roleId"`
}

func (LookupRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRolesArgs)(nil)).Elem()
}

// A collection of values returned by getRoles.
type LookupRolesResultOutput struct{ *pulumi.OutputState }

func (LookupRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRolesResult)(nil)).Elem()
}

func (o LookupRolesResultOutput) ToLookupRolesResultOutput() LookupRolesResultOutput {
	return o
}

func (o LookupRolesResultOutput) ToLookupRolesResultOutputWithContext(ctx context.Context) LookupRolesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRolesResultOutput) Results() GetRolesResultArrayOutput {
	return o.ApplyT(func(v LookupRolesResult) []GetRolesResult { return v.Results }).(GetRolesResultArrayOutput)
}

func (o LookupRolesResultOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRolesResult) string { return v.RoleId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRolesResultOutput{})
}
