// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package upstash

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/upstash/pulumi-upstash/sdk/go/upstash"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := upstash.NewKafkaCluster(ctx, "exampleCluster", &upstash.KafkaClusterArgs{
//				ClusterName: pulumi.String("TerraformCluster"),
//				Multizone:   pulumi.Bool(false),
//				Region:      pulumi.String("eu-west-1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type KafkaCluster struct {
	pulumi.CustomResourceState

	// Unique Cluster ID for created cluster
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Name of the cluster
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Creation time of the cluster
	CreationTime pulumi.IntOutput `pulumi:"creationTime"`
	// Max Message Size for the cluster
	MaxMessageSize pulumi.IntOutput `pulumi:"maxMessageSize"`
	// Max Messages Per Second for the cluster
	MaxMessagesPerSecond pulumi.IntOutput `pulumi:"maxMessagesPerSecond"`
	// Max Partitions for the cluster
	MaxPartitions pulumi.IntOutput `pulumi:"maxPartitions"`
	// Max Retention Size of the cluster
	MaxRetentionSize pulumi.IntOutput `pulumi:"maxRetentionSize"`
	// Max Retention Time of the cluster
	MaxRetentionTime pulumi.IntOutput `pulumi:"maxRetentionTime"`
	// Whether cluster has multizone attribute
	Multizone pulumi.BoolPtrOutput `pulumi:"multizone"`
	// Password for the cluster
	Password pulumi.StringOutput `pulumi:"password"`
	// region of the cluster. Possible values (may change) are: "eu-west-1", "us-east-1"
	Region pulumi.StringOutput `pulumi:"region"`
	// REST Endpoint of the cluster
	RestEndpoint pulumi.StringOutput `pulumi:"restEndpoint"`
	// State, where the cluster is originated
	State pulumi.StringOutput `pulumi:"state"`
	// TCP Endpoint of the cluster
	TcpEndpoint pulumi.StringOutput `pulumi:"tcpEndpoint"`
	// Type of the cluster
	Type pulumi.StringOutput `pulumi:"type"`
	// Base64 encoded username for the cluster
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewKafkaCluster registers a new resource with the given unique name, arguments, and options.
func NewKafkaCluster(ctx *pulumi.Context,
	name string, args *KafkaClusterArgs, opts ...pulumi.ResourceOption) (*KafkaCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource KafkaCluster
	err := ctx.RegisterResource("upstash:index/kafkaCluster:KafkaCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaCluster gets an existing KafkaCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaClusterState, opts ...pulumi.ResourceOption) (*KafkaCluster, error) {
	var resource KafkaCluster
	err := ctx.ReadResource("upstash:index/kafkaCluster:KafkaCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaCluster resources.
type kafkaClusterState struct {
	// Unique Cluster ID for created cluster
	ClusterId *string `pulumi:"clusterId"`
	// Name of the cluster
	ClusterName *string `pulumi:"clusterName"`
	// Creation time of the cluster
	CreationTime *int `pulumi:"creationTime"`
	// Max Message Size for the cluster
	MaxMessageSize *int `pulumi:"maxMessageSize"`
	// Max Messages Per Second for the cluster
	MaxMessagesPerSecond *int `pulumi:"maxMessagesPerSecond"`
	// Max Partitions for the cluster
	MaxPartitions *int `pulumi:"maxPartitions"`
	// Max Retention Size of the cluster
	MaxRetentionSize *int `pulumi:"maxRetentionSize"`
	// Max Retention Time of the cluster
	MaxRetentionTime *int `pulumi:"maxRetentionTime"`
	// Whether cluster has multizone attribute
	Multizone *bool `pulumi:"multizone"`
	// Password for the cluster
	Password *string `pulumi:"password"`
	// region of the cluster. Possible values (may change) are: "eu-west-1", "us-east-1"
	Region *string `pulumi:"region"`
	// REST Endpoint of the cluster
	RestEndpoint *string `pulumi:"restEndpoint"`
	// State, where the cluster is originated
	State *string `pulumi:"state"`
	// TCP Endpoint of the cluster
	TcpEndpoint *string `pulumi:"tcpEndpoint"`
	// Type of the cluster
	Type *string `pulumi:"type"`
	// Base64 encoded username for the cluster
	Username *string `pulumi:"username"`
}

type KafkaClusterState struct {
	// Unique Cluster ID for created cluster
	ClusterId pulumi.StringPtrInput
	// Name of the cluster
	ClusterName pulumi.StringPtrInput
	// Creation time of the cluster
	CreationTime pulumi.IntPtrInput
	// Max Message Size for the cluster
	MaxMessageSize pulumi.IntPtrInput
	// Max Messages Per Second for the cluster
	MaxMessagesPerSecond pulumi.IntPtrInput
	// Max Partitions for the cluster
	MaxPartitions pulumi.IntPtrInput
	// Max Retention Size of the cluster
	MaxRetentionSize pulumi.IntPtrInput
	// Max Retention Time of the cluster
	MaxRetentionTime pulumi.IntPtrInput
	// Whether cluster has multizone attribute
	Multizone pulumi.BoolPtrInput
	// Password for the cluster
	Password pulumi.StringPtrInput
	// region of the cluster. Possible values (may change) are: "eu-west-1", "us-east-1"
	Region pulumi.StringPtrInput
	// REST Endpoint of the cluster
	RestEndpoint pulumi.StringPtrInput
	// State, where the cluster is originated
	State pulumi.StringPtrInput
	// TCP Endpoint of the cluster
	TcpEndpoint pulumi.StringPtrInput
	// Type of the cluster
	Type pulumi.StringPtrInput
	// Base64 encoded username for the cluster
	Username pulumi.StringPtrInput
}

func (KafkaClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaClusterState)(nil)).Elem()
}

type kafkaClusterArgs struct {
	// Name of the cluster
	ClusterName string `pulumi:"clusterName"`
	// Whether cluster has multizone attribute
	Multizone *bool `pulumi:"multizone"`
	// region of the cluster. Possible values (may change) are: "eu-west-1", "us-east-1"
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a KafkaCluster resource.
type KafkaClusterArgs struct {
	// Name of the cluster
	ClusterName pulumi.StringInput
	// Whether cluster has multizone attribute
	Multizone pulumi.BoolPtrInput
	// region of the cluster. Possible values (may change) are: "eu-west-1", "us-east-1"
	Region pulumi.StringInput
}

func (KafkaClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaClusterArgs)(nil)).Elem()
}

type KafkaClusterInput interface {
	pulumi.Input

	ToKafkaClusterOutput() KafkaClusterOutput
	ToKafkaClusterOutputWithContext(ctx context.Context) KafkaClusterOutput
}

func (*KafkaCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaCluster)(nil)).Elem()
}

func (i *KafkaCluster) ToKafkaClusterOutput() KafkaClusterOutput {
	return i.ToKafkaClusterOutputWithContext(context.Background())
}

func (i *KafkaCluster) ToKafkaClusterOutputWithContext(ctx context.Context) KafkaClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaClusterOutput)
}

// KafkaClusterArrayInput is an input type that accepts KafkaClusterArray and KafkaClusterArrayOutput values.
// You can construct a concrete instance of `KafkaClusterArrayInput` via:
//
//	KafkaClusterArray{ KafkaClusterArgs{...} }
type KafkaClusterArrayInput interface {
	pulumi.Input

	ToKafkaClusterArrayOutput() KafkaClusterArrayOutput
	ToKafkaClusterArrayOutputWithContext(context.Context) KafkaClusterArrayOutput
}

type KafkaClusterArray []KafkaClusterInput

func (KafkaClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaCluster)(nil)).Elem()
}

func (i KafkaClusterArray) ToKafkaClusterArrayOutput() KafkaClusterArrayOutput {
	return i.ToKafkaClusterArrayOutputWithContext(context.Background())
}

func (i KafkaClusterArray) ToKafkaClusterArrayOutputWithContext(ctx context.Context) KafkaClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaClusterArrayOutput)
}

// KafkaClusterMapInput is an input type that accepts KafkaClusterMap and KafkaClusterMapOutput values.
// You can construct a concrete instance of `KafkaClusterMapInput` via:
//
//	KafkaClusterMap{ "key": KafkaClusterArgs{...} }
type KafkaClusterMapInput interface {
	pulumi.Input

	ToKafkaClusterMapOutput() KafkaClusterMapOutput
	ToKafkaClusterMapOutputWithContext(context.Context) KafkaClusterMapOutput
}

type KafkaClusterMap map[string]KafkaClusterInput

func (KafkaClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaCluster)(nil)).Elem()
}

func (i KafkaClusterMap) ToKafkaClusterMapOutput() KafkaClusterMapOutput {
	return i.ToKafkaClusterMapOutputWithContext(context.Background())
}

func (i KafkaClusterMap) ToKafkaClusterMapOutputWithContext(ctx context.Context) KafkaClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaClusterMapOutput)
}

type KafkaClusterOutput struct{ *pulumi.OutputState }

func (KafkaClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaCluster)(nil)).Elem()
}

func (o KafkaClusterOutput) ToKafkaClusterOutput() KafkaClusterOutput {
	return o
}

func (o KafkaClusterOutput) ToKafkaClusterOutputWithContext(ctx context.Context) KafkaClusterOutput {
	return o
}

// Unique Cluster ID for created cluster
func (o KafkaClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Name of the cluster
func (o KafkaClusterOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// Creation time of the cluster
func (o KafkaClusterOutput) CreationTime() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.IntOutput { return v.CreationTime }).(pulumi.IntOutput)
}

// Max Message Size for the cluster
func (o KafkaClusterOutput) MaxMessageSize() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.IntOutput { return v.MaxMessageSize }).(pulumi.IntOutput)
}

// Max Messages Per Second for the cluster
func (o KafkaClusterOutput) MaxMessagesPerSecond() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.IntOutput { return v.MaxMessagesPerSecond }).(pulumi.IntOutput)
}

// Max Partitions for the cluster
func (o KafkaClusterOutput) MaxPartitions() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.IntOutput { return v.MaxPartitions }).(pulumi.IntOutput)
}

// Max Retention Size of the cluster
func (o KafkaClusterOutput) MaxRetentionSize() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.IntOutput { return v.MaxRetentionSize }).(pulumi.IntOutput)
}

// Max Retention Time of the cluster
func (o KafkaClusterOutput) MaxRetentionTime() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.IntOutput { return v.MaxRetentionTime }).(pulumi.IntOutput)
}

// Whether cluster has multizone attribute
func (o KafkaClusterOutput) Multizone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.BoolPtrOutput { return v.Multizone }).(pulumi.BoolPtrOutput)
}

// Password for the cluster
func (o KafkaClusterOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// region of the cluster. Possible values (may change) are: "eu-west-1", "us-east-1"
func (o KafkaClusterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// REST Endpoint of the cluster
func (o KafkaClusterOutput) RestEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.StringOutput { return v.RestEndpoint }).(pulumi.StringOutput)
}

// State, where the cluster is originated
func (o KafkaClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// TCP Endpoint of the cluster
func (o KafkaClusterOutput) TcpEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.StringOutput { return v.TcpEndpoint }).(pulumi.StringOutput)
}

// Type of the cluster
func (o KafkaClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Base64 encoded username for the cluster
func (o KafkaClusterOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaCluster) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type KafkaClusterArrayOutput struct{ *pulumi.OutputState }

func (KafkaClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaCluster)(nil)).Elem()
}

func (o KafkaClusterArrayOutput) ToKafkaClusterArrayOutput() KafkaClusterArrayOutput {
	return o
}

func (o KafkaClusterArrayOutput) ToKafkaClusterArrayOutputWithContext(ctx context.Context) KafkaClusterArrayOutput {
	return o
}

func (o KafkaClusterArrayOutput) Index(i pulumi.IntInput) KafkaClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaCluster {
		return vs[0].([]*KafkaCluster)[vs[1].(int)]
	}).(KafkaClusterOutput)
}

type KafkaClusterMapOutput struct{ *pulumi.OutputState }

func (KafkaClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaCluster)(nil)).Elem()
}

func (o KafkaClusterMapOutput) ToKafkaClusterMapOutput() KafkaClusterMapOutput {
	return o
}

func (o KafkaClusterMapOutput) ToKafkaClusterMapOutputWithContext(ctx context.Context) KafkaClusterMapOutput {
	return o
}

func (o KafkaClusterMapOutput) MapIndex(k pulumi.StringInput) KafkaClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaCluster {
		return vs[0].(map[string]*KafkaCluster)[vs[1].(string)]
	}).(KafkaClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaClusterInput)(nil)).Elem(), &KafkaCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaClusterArrayInput)(nil)).Elem(), KafkaClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaClusterMapInput)(nil)).Elem(), KafkaClusterMap{})
	pulumi.RegisterOutputType(KafkaClusterOutput{})
	pulumi.RegisterOutputType(KafkaClusterArrayOutput{})
	pulumi.RegisterOutputType(KafkaClusterMapOutput{})
}
