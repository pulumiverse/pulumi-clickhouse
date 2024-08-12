// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clickhouse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-clickhouse/sdk/go/clickhouse/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-clickhouse/sdk/go/clickhouse"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clickhouse.NewService(ctx, "service", &clickhouse.ServiceArgs{
//				CloudProvider:      pulumi.String("aws"),
//				IdleScaling:        pulumi.Bool(true),
//				IdleTimeoutMinutes: pulumi.Int(5),
//				IpAccesses: clickhouse.ServiceIpAccessArray{
//					&clickhouse.ServiceIpAccessArgs{
//						Description: pulumi.String("Test IP"),
//						Source:      pulumi.String("192.168.2.63"),
//					},
//				},
//				MaxTotalMemoryGb: pulumi.Int(360),
//				MinTotalMemoryGb: pulumi.Int(24),
//				PasswordHash:     pulumi.String("n4bQgYhMfWWaL+qgxVrQFaO/TxsrC4Is0V1sFbDwCgg="),
//				Region:           pulumi.String("us-east-1"),
//				Tier:             pulumi.String("production"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Services can be imported by specifying the UUID.
//
// ```sh
// $ pulumi import clickhouse:index/service:Service example xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// ```
type Service struct {
	pulumi.CustomResourceState

	// Cloud provider ('aws', 'gcp', or 'azure') in which the service is deployed in.
	CloudProvider pulumi.StringOutput `pulumi:"cloudProvider"`
	// Double SHA1 hash of password for connecting with the MySQL protocol. Cannot be specified if `password` is specified.
	DoubleSha1PasswordHash pulumi.StringPtrOutput `pulumi:"doubleSha1PasswordHash"`
	// Custom role identifier arn
	EncryptionAssumedRoleIdentifier pulumi.StringPtrOutput `pulumi:"encryptionAssumedRoleIdentifier"`
	// Custom encryption key arn
	EncryptionKey pulumi.StringPtrOutput `pulumi:"encryptionKey"`
	// List of public endpoints.
	Endpoints ServiceEndpointArrayOutput `pulumi:"endpoints"`
	// IAM role used for accessing objects in s3.
	IamRole pulumi.StringOutput `pulumi:"iamRole"`
	// When set to true the service is allowed to scale down to zero when idle.
	IdleScaling pulumi.BoolPtrOutput `pulumi:"idleScaling"`
	// Set minimum idling timeout (in minutes). Must be greater than or equal to 5 minutes. Must be set if idleScaling is enabled
	IdleTimeoutMinutes pulumi.IntPtrOutput `pulumi:"idleTimeoutMinutes"`
	// List of IP addresses allowed to access the service.
	IpAccesses  ServiceIpAccessArrayOutput `pulumi:"ipAccesses"`
	LastUpdated pulumi.StringOutput        `pulumi:"lastUpdated"`
	// Maximum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and lower than 360 for non paid services or 720 for paid services.
	MaxTotalMemoryGb pulumi.IntPtrOutput `pulumi:"maxTotalMemoryGb"`
	// Minimum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and greater than 24.
	MinTotalMemoryGb pulumi.IntPtrOutput `pulumi:"minTotalMemoryGb"`
	// User defined identifier for the service.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for the default user. One of either `password` or `passwordHash` must be specified.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// SHA256 hash of password for the default user. One of either `password` or `passwordHash` must be specified.
	PasswordHash pulumi.StringPtrOutput `pulumi:"passwordHash"`
	// Service config for private endpoints
	//
	// Deprecated: Please use the `PrivateEndpoint.getConfig` data source instead.
	PrivateEndpointConfig ServicePrivateEndpointConfigOutput `pulumi:"privateEndpointConfig"`
	// List of private endpoint IDs
	PrivateEndpointIds pulumi.StringArrayOutput `pulumi:"privateEndpointIds"`
	// Region within the cloud provider in which the service is deployed in.
	Region pulumi.StringOutput `pulumi:"region"`
	// Tier of the service: 'development', 'production'. Production services scale, Development are fixed size.
	Tier pulumi.StringOutput `pulumi:"tier"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudProvider == nil {
		return nil, errors.New("invalid value for required argument 'CloudProvider'")
	}
	if args.IpAccesses == nil {
		return nil, errors.New("invalid value for required argument 'IpAccesses'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	if args.DoubleSha1PasswordHash != nil {
		args.DoubleSha1PasswordHash = pulumi.ToSecret(args.DoubleSha1PasswordHash).(pulumi.StringPtrInput)
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.PasswordHash != nil {
		args.PasswordHash = pulumi.ToSecret(args.PasswordHash).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"doubleSha1PasswordHash",
		"password",
		"passwordHash",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("clickhouse:index/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("clickhouse:index/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// Cloud provider ('aws', 'gcp', or 'azure') in which the service is deployed in.
	CloudProvider *string `pulumi:"cloudProvider"`
	// Double SHA1 hash of password for connecting with the MySQL protocol. Cannot be specified if `password` is specified.
	DoubleSha1PasswordHash *string `pulumi:"doubleSha1PasswordHash"`
	// Custom role identifier arn
	EncryptionAssumedRoleIdentifier *string `pulumi:"encryptionAssumedRoleIdentifier"`
	// Custom encryption key arn
	EncryptionKey *string `pulumi:"encryptionKey"`
	// List of public endpoints.
	Endpoints []ServiceEndpoint `pulumi:"endpoints"`
	// IAM role used for accessing objects in s3.
	IamRole *string `pulumi:"iamRole"`
	// When set to true the service is allowed to scale down to zero when idle.
	IdleScaling *bool `pulumi:"idleScaling"`
	// Set minimum idling timeout (in minutes). Must be greater than or equal to 5 minutes. Must be set if idleScaling is enabled
	IdleTimeoutMinutes *int `pulumi:"idleTimeoutMinutes"`
	// List of IP addresses allowed to access the service.
	IpAccesses  []ServiceIpAccess `pulumi:"ipAccesses"`
	LastUpdated *string           `pulumi:"lastUpdated"`
	// Maximum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and lower than 360 for non paid services or 720 for paid services.
	MaxTotalMemoryGb *int `pulumi:"maxTotalMemoryGb"`
	// Minimum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and greater than 24.
	MinTotalMemoryGb *int `pulumi:"minTotalMemoryGb"`
	// User defined identifier for the service.
	Name *string `pulumi:"name"`
	// Password for the default user. One of either `password` or `passwordHash` must be specified.
	Password *string `pulumi:"password"`
	// SHA256 hash of password for the default user. One of either `password` or `passwordHash` must be specified.
	PasswordHash *string `pulumi:"passwordHash"`
	// Service config for private endpoints
	//
	// Deprecated: Please use the `PrivateEndpoint.getConfig` data source instead.
	PrivateEndpointConfig *ServicePrivateEndpointConfig `pulumi:"privateEndpointConfig"`
	// List of private endpoint IDs
	PrivateEndpointIds []string `pulumi:"privateEndpointIds"`
	// Region within the cloud provider in which the service is deployed in.
	Region *string `pulumi:"region"`
	// Tier of the service: 'development', 'production'. Production services scale, Development are fixed size.
	Tier *string `pulumi:"tier"`
}

type ServiceState struct {
	// Cloud provider ('aws', 'gcp', or 'azure') in which the service is deployed in.
	CloudProvider pulumi.StringPtrInput
	// Double SHA1 hash of password for connecting with the MySQL protocol. Cannot be specified if `password` is specified.
	DoubleSha1PasswordHash pulumi.StringPtrInput
	// Custom role identifier arn
	EncryptionAssumedRoleIdentifier pulumi.StringPtrInput
	// Custom encryption key arn
	EncryptionKey pulumi.StringPtrInput
	// List of public endpoints.
	Endpoints ServiceEndpointArrayInput
	// IAM role used for accessing objects in s3.
	IamRole pulumi.StringPtrInput
	// When set to true the service is allowed to scale down to zero when idle.
	IdleScaling pulumi.BoolPtrInput
	// Set minimum idling timeout (in minutes). Must be greater than or equal to 5 minutes. Must be set if idleScaling is enabled
	IdleTimeoutMinutes pulumi.IntPtrInput
	// List of IP addresses allowed to access the service.
	IpAccesses  ServiceIpAccessArrayInput
	LastUpdated pulumi.StringPtrInput
	// Maximum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and lower than 360 for non paid services or 720 for paid services.
	MaxTotalMemoryGb pulumi.IntPtrInput
	// Minimum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and greater than 24.
	MinTotalMemoryGb pulumi.IntPtrInput
	// User defined identifier for the service.
	Name pulumi.StringPtrInput
	// Password for the default user. One of either `password` or `passwordHash` must be specified.
	Password pulumi.StringPtrInput
	// SHA256 hash of password for the default user. One of either `password` or `passwordHash` must be specified.
	PasswordHash pulumi.StringPtrInput
	// Service config for private endpoints
	//
	// Deprecated: Please use the `PrivateEndpoint.getConfig` data source instead.
	PrivateEndpointConfig ServicePrivateEndpointConfigPtrInput
	// List of private endpoint IDs
	PrivateEndpointIds pulumi.StringArrayInput
	// Region within the cloud provider in which the service is deployed in.
	Region pulumi.StringPtrInput
	// Tier of the service: 'development', 'production'. Production services scale, Development are fixed size.
	Tier pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Cloud provider ('aws', 'gcp', or 'azure') in which the service is deployed in.
	CloudProvider string `pulumi:"cloudProvider"`
	// Double SHA1 hash of password for connecting with the MySQL protocol. Cannot be specified if `password` is specified.
	DoubleSha1PasswordHash *string `pulumi:"doubleSha1PasswordHash"`
	// Custom role identifier arn
	EncryptionAssumedRoleIdentifier *string `pulumi:"encryptionAssumedRoleIdentifier"`
	// Custom encryption key arn
	EncryptionKey *string `pulumi:"encryptionKey"`
	// When set to true the service is allowed to scale down to zero when idle.
	IdleScaling *bool `pulumi:"idleScaling"`
	// Set minimum idling timeout (in minutes). Must be greater than or equal to 5 minutes. Must be set if idleScaling is enabled
	IdleTimeoutMinutes *int `pulumi:"idleTimeoutMinutes"`
	// List of IP addresses allowed to access the service.
	IpAccesses []ServiceIpAccess `pulumi:"ipAccesses"`
	// Maximum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and lower than 360 for non paid services or 720 for paid services.
	MaxTotalMemoryGb *int `pulumi:"maxTotalMemoryGb"`
	// Minimum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and greater than 24.
	MinTotalMemoryGb *int `pulumi:"minTotalMemoryGb"`
	// User defined identifier for the service.
	Name *string `pulumi:"name"`
	// Password for the default user. One of either `password` or `passwordHash` must be specified.
	Password *string `pulumi:"password"`
	// SHA256 hash of password for the default user. One of either `password` or `passwordHash` must be specified.
	PasswordHash *string `pulumi:"passwordHash"`
	// List of private endpoint IDs
	PrivateEndpointIds []string `pulumi:"privateEndpointIds"`
	// Region within the cloud provider in which the service is deployed in.
	Region string `pulumi:"region"`
	// Tier of the service: 'development', 'production'. Production services scale, Development are fixed size.
	Tier string `pulumi:"tier"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Cloud provider ('aws', 'gcp', or 'azure') in which the service is deployed in.
	CloudProvider pulumi.StringInput
	// Double SHA1 hash of password for connecting with the MySQL protocol. Cannot be specified if `password` is specified.
	DoubleSha1PasswordHash pulumi.StringPtrInput
	// Custom role identifier arn
	EncryptionAssumedRoleIdentifier pulumi.StringPtrInput
	// Custom encryption key arn
	EncryptionKey pulumi.StringPtrInput
	// When set to true the service is allowed to scale down to zero when idle.
	IdleScaling pulumi.BoolPtrInput
	// Set minimum idling timeout (in minutes). Must be greater than or equal to 5 minutes. Must be set if idleScaling is enabled
	IdleTimeoutMinutes pulumi.IntPtrInput
	// List of IP addresses allowed to access the service.
	IpAccesses ServiceIpAccessArrayInput
	// Maximum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and lower than 360 for non paid services or 720 for paid services.
	MaxTotalMemoryGb pulumi.IntPtrInput
	// Minimum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and greater than 24.
	MinTotalMemoryGb pulumi.IntPtrInput
	// User defined identifier for the service.
	Name pulumi.StringPtrInput
	// Password for the default user. One of either `password` or `passwordHash` must be specified.
	Password pulumi.StringPtrInput
	// SHA256 hash of password for the default user. One of either `password` or `passwordHash` must be specified.
	PasswordHash pulumi.StringPtrInput
	// List of private endpoint IDs
	PrivateEndpointIds pulumi.StringArrayInput
	// Region within the cloud provider in which the service is deployed in.
	Region pulumi.StringInput
	// Tier of the service: 'development', 'production'. Production services scale, Development are fixed size.
	Tier pulumi.StringInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//	ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//	ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// Cloud provider ('aws', 'gcp', or 'azure') in which the service is deployed in.
func (o ServiceOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.CloudProvider }).(pulumi.StringOutput)
}

// Double SHA1 hash of password for connecting with the MySQL protocol. Cannot be specified if `password` is specified.
func (o ServiceOutput) DoubleSha1PasswordHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.DoubleSha1PasswordHash }).(pulumi.StringPtrOutput)
}

// Custom role identifier arn
func (o ServiceOutput) EncryptionAssumedRoleIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.EncryptionAssumedRoleIdentifier }).(pulumi.StringPtrOutput)
}

// Custom encryption key arn
func (o ServiceOutput) EncryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.EncryptionKey }).(pulumi.StringPtrOutput)
}

// List of public endpoints.
func (o ServiceOutput) Endpoints() ServiceEndpointArrayOutput {
	return o.ApplyT(func(v *Service) ServiceEndpointArrayOutput { return v.Endpoints }).(ServiceEndpointArrayOutput)
}

// IAM role used for accessing objects in s3.
func (o ServiceOutput) IamRole() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.IamRole }).(pulumi.StringOutput)
}

// When set to true the service is allowed to scale down to zero when idle.
func (o ServiceOutput) IdleScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolPtrOutput { return v.IdleScaling }).(pulumi.BoolPtrOutput)
}

// Set minimum idling timeout (in minutes). Must be greater than or equal to 5 minutes. Must be set if idleScaling is enabled
func (o ServiceOutput) IdleTimeoutMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.IdleTimeoutMinutes }).(pulumi.IntPtrOutput)
}

// List of IP addresses allowed to access the service.
func (o ServiceOutput) IpAccesses() ServiceIpAccessArrayOutput {
	return o.ApplyT(func(v *Service) ServiceIpAccessArrayOutput { return v.IpAccesses }).(ServiceIpAccessArrayOutput)
}

func (o ServiceOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.LastUpdated }).(pulumi.StringOutput)
}

// Maximum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and lower than 360 for non paid services or 720 for paid services.
func (o ServiceOutput) MaxTotalMemoryGb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.MaxTotalMemoryGb }).(pulumi.IntPtrOutput)
}

// Minimum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and greater than 24.
func (o ServiceOutput) MinTotalMemoryGb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.MinTotalMemoryGb }).(pulumi.IntPtrOutput)
}

// User defined identifier for the service.
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for the default user. One of either `password` or `passwordHash` must be specified.
func (o ServiceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// SHA256 hash of password for the default user. One of either `password` or `passwordHash` must be specified.
func (o ServiceOutput) PasswordHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.PasswordHash }).(pulumi.StringPtrOutput)
}

// Service config for private endpoints
//
// Deprecated: Please use the `PrivateEndpoint.getConfig` data source instead.
func (o ServiceOutput) PrivateEndpointConfig() ServicePrivateEndpointConfigOutput {
	return o.ApplyT(func(v *Service) ServicePrivateEndpointConfigOutput { return v.PrivateEndpointConfig }).(ServicePrivateEndpointConfigOutput)
}

// List of private endpoint IDs
func (o ServiceOutput) PrivateEndpointIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Service) pulumi.StringArrayOutput { return v.PrivateEndpointIds }).(pulumi.StringArrayOutput)
}

// Region within the cloud provider in which the service is deployed in.
func (o ServiceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Tier of the service: 'development', 'production'. Production services scale, Development are fixed size.
func (o ServiceOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Tier }).(pulumi.StringOutput)
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
