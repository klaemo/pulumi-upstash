// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package upstash

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-upstash/sdk/go/upstash"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/upstash/pulumi-upstash/sdk/go/upstash"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := upstash.LookupQStashEndpoint(ctx, &GetQStashEndpointArgs{
//				EndpointId: resource.Upstash_qstash_endpoint.ExampleQstashEndpoint.Endpoint_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupQStashEndpoint(ctx *pulumi.Context, args *LookupQStashEndpointArgs, opts ...pulumi.InvokeOption) (*LookupQStashEndpointResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupQStashEndpointResult
	err := ctx.Invoke("upstash:index/getQStashEndpoint:getQStashEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQStashEndpoint.
type LookupQStashEndpointArgs struct {
	TopicId string `pulumi:"topicId"`
}

// A collection of values returned by getQStashEndpoint.
type LookupQStashEndpointResult struct {
	EndpointId string `pulumi:"endpointId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	TopicId string `pulumi:"topicId"`
	Url     string `pulumi:"url"`
}

func LookupQStashEndpointOutput(ctx *pulumi.Context, args LookupQStashEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupQStashEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQStashEndpointResult, error) {
			args := v.(LookupQStashEndpointArgs)
			r, err := LookupQStashEndpoint(ctx, &args, opts...)
			var s LookupQStashEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQStashEndpointResultOutput)
}

// A collection of arguments for invoking getQStashEndpoint.
type LookupQStashEndpointOutputArgs struct {
	TopicId pulumi.StringInput `pulumi:"topicId"`
}

func (LookupQStashEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQStashEndpointArgs)(nil)).Elem()
}

// A collection of values returned by getQStashEndpoint.
type LookupQStashEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupQStashEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQStashEndpointResult)(nil)).Elem()
}

func (o LookupQStashEndpointResultOutput) ToLookupQStashEndpointResultOutput() LookupQStashEndpointResultOutput {
	return o
}

func (o LookupQStashEndpointResultOutput) ToLookupQStashEndpointResultOutputWithContext(ctx context.Context) LookupQStashEndpointResultOutput {
	return o
}

func (o LookupQStashEndpointResultOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQStashEndpointResult) string { return v.EndpointId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQStashEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQStashEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQStashEndpointResultOutput) TopicId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQStashEndpointResult) string { return v.TopicId }).(pulumi.StringOutput)
}

func (o LookupQStashEndpointResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQStashEndpointResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQStashEndpointResultOutput{})
}
