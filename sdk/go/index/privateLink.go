// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package index

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `PrivateLink` provides a private link resource. Private Link is a private network endpoint that can be created to connect from your vpc to Astra without using a publicly routable IP address. `PrivateLink` resources are associated with a database id. Once the privateLink resource is created in Astra it must be linked to an endpoint within your vpc, use `PrivateLinkEndpoint` to do this.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-index/sdk/go/index"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := astra.NewPrivateLink(ctx, "example", &astra.PrivateLinkArgs{
// 			AllowedPrincipals: pulumi.StringArray{
// 				pulumi.String("arn:aws:iam::111708290731:user/sebastian.estevez"),
// 			},
// 			DatabaseId:   pulumi.String("a6bc9c26-e7ce-424f-84c7-0a00afb12588"),
// 			DatacenterId: pulumi.String("a6bc9c26-e7ce-424f-84c7-0a00afb12588-1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ```sh
//  $ pulumi import astra:index/privateLink:PrivateLink example a6bc9c26-e7ce-424f-84c7-0a00afb12588/datacenter/a6bc9c26-e7ce-424f-84c7-0a00afb12588/serviceNames/svc-name-here
// ```
type PrivateLink struct {
	pulumi.CustomResourceState

	// List of service principals to apply to the Private Link (i.e. arn:aws:iam::123456789012:role/admin).
	AllowedPrincipals pulumi.StringArrayOutput `pulumi:"allowedPrincipals"`
	// Astra database where private link will be enabled.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// Astra datacenter in the region where the private link will be created.
	DatacenterId pulumi.StringOutput `pulumi:"datacenterId"`
	// Name of the endpoint service for private link generated by the cloud provider.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewPrivateLink registers a new resource with the given unique name, arguments, and options.
func NewPrivateLink(ctx *pulumi.Context,
	name string, args *PrivateLinkArgs, opts ...pulumi.ResourceOption) (*PrivateLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowedPrincipals == nil {
		return nil, errors.New("invalid value for required argument 'AllowedPrincipals'")
	}
	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	if args.DatacenterId == nil {
		return nil, errors.New("invalid value for required argument 'DatacenterId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PrivateLink
	err := ctx.RegisterResource("astra:index/privateLink:PrivateLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLink gets an existing PrivateLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkState, opts ...pulumi.ResourceOption) (*PrivateLink, error) {
	var resource PrivateLink
	err := ctx.ReadResource("astra:index/privateLink:PrivateLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLink resources.
type privateLinkState struct {
	// List of service principals to apply to the Private Link (i.e. arn:aws:iam::123456789012:role/admin).
	AllowedPrincipals []string `pulumi:"allowedPrincipals"`
	// Astra database where private link will be enabled.
	DatabaseId *string `pulumi:"databaseId"`
	// Astra datacenter in the region where the private link will be created.
	DatacenterId *string `pulumi:"datacenterId"`
	// Name of the endpoint service for private link generated by the cloud provider.
	ServiceName *string `pulumi:"serviceName"`
}

type PrivateLinkState struct {
	// List of service principals to apply to the Private Link (i.e. arn:aws:iam::123456789012:role/admin).
	AllowedPrincipals pulumi.StringArrayInput
	// Astra database where private link will be enabled.
	DatabaseId pulumi.StringPtrInput
	// Astra datacenter in the region where the private link will be created.
	DatacenterId pulumi.StringPtrInput
	// Name of the endpoint service for private link generated by the cloud provider.
	ServiceName pulumi.StringPtrInput
}

func (PrivateLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkState)(nil)).Elem()
}

type privateLinkArgs struct {
	// List of service principals to apply to the Private Link (i.e. arn:aws:iam::123456789012:role/admin).
	AllowedPrincipals []string `pulumi:"allowedPrincipals"`
	// Astra database where private link will be enabled.
	DatabaseId string `pulumi:"databaseId"`
	// Astra datacenter in the region where the private link will be created.
	DatacenterId string `pulumi:"datacenterId"`
}

// The set of arguments for constructing a PrivateLink resource.
type PrivateLinkArgs struct {
	// List of service principals to apply to the Private Link (i.e. arn:aws:iam::123456789012:role/admin).
	AllowedPrincipals pulumi.StringArrayInput
	// Astra database where private link will be enabled.
	DatabaseId pulumi.StringInput
	// Astra datacenter in the region where the private link will be created.
	DatacenterId pulumi.StringInput
}

func (PrivateLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkArgs)(nil)).Elem()
}

type PrivateLinkInput interface {
	pulumi.Input

	ToPrivateLinkOutput() PrivateLinkOutput
	ToPrivateLinkOutputWithContext(ctx context.Context) PrivateLinkOutput
}

func (*PrivateLink) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLink)(nil)).Elem()
}

func (i *PrivateLink) ToPrivateLinkOutput() PrivateLinkOutput {
	return i.ToPrivateLinkOutputWithContext(context.Background())
}

func (i *PrivateLink) ToPrivateLinkOutputWithContext(ctx context.Context) PrivateLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkOutput)
}

// PrivateLinkArrayInput is an input type that accepts PrivateLinkArray and PrivateLinkArrayOutput values.
// You can construct a concrete instance of `PrivateLinkArrayInput` via:
//
//          PrivateLinkArray{ PrivateLinkArgs{...} }
type PrivateLinkArrayInput interface {
	pulumi.Input

	ToPrivateLinkArrayOutput() PrivateLinkArrayOutput
	ToPrivateLinkArrayOutputWithContext(context.Context) PrivateLinkArrayOutput
}

type PrivateLinkArray []PrivateLinkInput

func (PrivateLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateLink)(nil)).Elem()
}

func (i PrivateLinkArray) ToPrivateLinkArrayOutput() PrivateLinkArrayOutput {
	return i.ToPrivateLinkArrayOutputWithContext(context.Background())
}

func (i PrivateLinkArray) ToPrivateLinkArrayOutputWithContext(ctx context.Context) PrivateLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkArrayOutput)
}

// PrivateLinkMapInput is an input type that accepts PrivateLinkMap and PrivateLinkMapOutput values.
// You can construct a concrete instance of `PrivateLinkMapInput` via:
//
//          PrivateLinkMap{ "key": PrivateLinkArgs{...} }
type PrivateLinkMapInput interface {
	pulumi.Input

	ToPrivateLinkMapOutput() PrivateLinkMapOutput
	ToPrivateLinkMapOutputWithContext(context.Context) PrivateLinkMapOutput
}

type PrivateLinkMap map[string]PrivateLinkInput

func (PrivateLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateLink)(nil)).Elem()
}

func (i PrivateLinkMap) ToPrivateLinkMapOutput() PrivateLinkMapOutput {
	return i.ToPrivateLinkMapOutputWithContext(context.Background())
}

func (i PrivateLinkMap) ToPrivateLinkMapOutputWithContext(ctx context.Context) PrivateLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkMapOutput)
}

type PrivateLinkOutput struct{ *pulumi.OutputState }

func (PrivateLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLink)(nil)).Elem()
}

func (o PrivateLinkOutput) ToPrivateLinkOutput() PrivateLinkOutput {
	return o
}

func (o PrivateLinkOutput) ToPrivateLinkOutputWithContext(ctx context.Context) PrivateLinkOutput {
	return o
}

// List of service principals to apply to the Private Link (i.e. arn:aws:iam::123456789012:role/admin).
func (o PrivateLinkOutput) AllowedPrincipals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateLink) pulumi.StringArrayOutput { return v.AllowedPrincipals }).(pulumi.StringArrayOutput)
}

// Astra database where private link will be enabled.
func (o PrivateLinkOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLink) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// Astra datacenter in the region where the private link will be created.
func (o PrivateLinkOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLink) pulumi.StringOutput { return v.DatacenterId }).(pulumi.StringOutput)
}

// Name of the endpoint service for private link generated by the cloud provider.
func (o PrivateLinkOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLink) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type PrivateLinkArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateLink)(nil)).Elem()
}

func (o PrivateLinkArrayOutput) ToPrivateLinkArrayOutput() PrivateLinkArrayOutput {
	return o
}

func (o PrivateLinkArrayOutput) ToPrivateLinkArrayOutputWithContext(ctx context.Context) PrivateLinkArrayOutput {
	return o
}

func (o PrivateLinkArrayOutput) Index(i pulumi.IntInput) PrivateLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateLink {
		return vs[0].([]*PrivateLink)[vs[1].(int)]
	}).(PrivateLinkOutput)
}

type PrivateLinkMapOutput struct{ *pulumi.OutputState }

func (PrivateLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateLink)(nil)).Elem()
}

func (o PrivateLinkMapOutput) ToPrivateLinkMapOutput() PrivateLinkMapOutput {
	return o
}

func (o PrivateLinkMapOutput) ToPrivateLinkMapOutputWithContext(ctx context.Context) PrivateLinkMapOutput {
	return o
}

func (o PrivateLinkMapOutput) MapIndex(k pulumi.StringInput) PrivateLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateLink {
		return vs[0].(map[string]*PrivateLink)[vs[1].(string)]
	}).(PrivateLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkInput)(nil)).Elem(), &PrivateLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkArrayInput)(nil)).Elem(), PrivateLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkMapInput)(nil)).Elem(), PrivateLinkMap{})
	pulumi.RegisterOutputType(PrivateLinkOutput{})
	pulumi.RegisterOutputType(PrivateLinkArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkMapOutput{})
}
