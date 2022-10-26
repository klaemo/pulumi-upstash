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
//	"github.com/upstash/upstash-pulumi-provider/sdk/go/upstash"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := upstash.NewTeam(ctx, "exampleTeam", &upstash.TeamArgs{
//				TeamName: pulumi.String("TerraformTeam"),
//				CopyCc:   pulumi.Bool(false),
//				TeamMembers: pulumi.StringMap{
//					"X@Y.Z": pulumi.String("owner"),
//					"A@B.C": pulumi.String("dev"),
//					"E@E.F": pulumi.String("finance"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Team struct {
	pulumi.CustomResourceState

	// Whether Credit Card is copied
	CopyCc pulumi.BoolOutput `pulumi:"copyCc"`
	// Unique Cluster ID for created cluster
	TeamId pulumi.StringOutput `pulumi:"teamId"`
	// Members of the team. (Owner must be specified, which is the owner of the api key.)
	TeamMembers pulumi.StringMapOutput `pulumi:"teamMembers"`
	// Name of the team
	TeamName pulumi.StringOutput `pulumi:"teamName"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CopyCc == nil {
		return nil, errors.New("invalid value for required argument 'CopyCc'")
	}
	if args.TeamMembers == nil {
		return nil, errors.New("invalid value for required argument 'TeamMembers'")
	}
	if args.TeamName == nil {
		return nil, errors.New("invalid value for required argument 'TeamName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Team
	err := ctx.RegisterResource("upstash:index/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("upstash:index/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	// Whether Credit Card is copied
	CopyCc *bool `pulumi:"copyCc"`
	// Unique Cluster ID for created cluster
	TeamId *string `pulumi:"teamId"`
	// Members of the team. (Owner must be specified, which is the owner of the api key.)
	TeamMembers map[string]string `pulumi:"teamMembers"`
	// Name of the team
	TeamName *string `pulumi:"teamName"`
}

type TeamState struct {
	// Whether Credit Card is copied
	CopyCc pulumi.BoolPtrInput
	// Unique Cluster ID for created cluster
	TeamId pulumi.StringPtrInput
	// Members of the team. (Owner must be specified, which is the owner of the api key.)
	TeamMembers pulumi.StringMapInput
	// Name of the team
	TeamName pulumi.StringPtrInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	// Whether Credit Card is copied
	CopyCc bool `pulumi:"copyCc"`
	// Members of the team. (Owner must be specified, which is the owner of the api key.)
	TeamMembers map[string]string `pulumi:"teamMembers"`
	// Name of the team
	TeamName string `pulumi:"teamName"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	// Whether Credit Card is copied
	CopyCc pulumi.BoolInput
	// Members of the team. (Owner must be specified, which is the owner of the api key.)
	TeamMembers pulumi.StringMapInput
	// Name of the team
	TeamName pulumi.StringInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}

type TeamInput interface {
	pulumi.Input

	ToTeamOutput() TeamOutput
	ToTeamOutputWithContext(ctx context.Context) TeamOutput
}

func (*Team) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//	TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (i TeamArray) ToTeamArrayOutput() TeamArrayOutput {
	return i.ToTeamArrayOutputWithContext(context.Background())
}

func (i TeamArray) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamArrayOutput)
}

// TeamMapInput is an input type that accepts TeamMap and TeamMapOutput values.
// You can construct a concrete instance of `TeamMapInput` via:
//
//	TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

type TeamOutput struct{ *pulumi.OutputState }

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

// Whether Credit Card is copied
func (o TeamOutput) CopyCc() pulumi.BoolOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolOutput { return v.CopyCc }).(pulumi.BoolOutput)
}

// Unique Cluster ID for created cluster
func (o TeamOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

// Members of the team. (Owner must be specified, which is the owner of the api key.)
func (o TeamOutput) TeamMembers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Team) pulumi.StringMapOutput { return v.TeamMembers }).(pulumi.StringMapOutput)
}

// Name of the team
func (o TeamOutput) TeamName() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.TeamName }).(pulumi.StringOutput)
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Team {
		return vs[0].([]*Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Team {
		return vs[0].(map[string]*Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamInput)(nil)).Elem(), &Team{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamArrayInput)(nil)).Elem(), TeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMapInput)(nil)).Elem(), TeamMap{})
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}
