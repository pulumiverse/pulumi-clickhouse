// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Clickhouse
{
    [ClickhouseResourceType("clickhouse:index/service:Service")]
    public partial class Service : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cloud provider ('aws', 'gcp', or 'azure') in which the service is deployed in.
        /// </summary>
        [Output("cloudProvider")]
        public Output<string> CloudProvider { get; private set; } = null!;

        /// <summary>
        /// Double SHA1 hash of password for connecting with the MySQL protocol. Cannot be specified if `password` is specified.
        /// </summary>
        [Output("doubleSha1PasswordHash")]
        public Output<string?> DoubleSha1PasswordHash { get; private set; } = null!;

        /// <summary>
        /// Custom role identifier arn
        /// </summary>
        [Output("encryptionAssumedRoleIdentifier")]
        public Output<string?> EncryptionAssumedRoleIdentifier { get; private set; } = null!;

        /// <summary>
        /// Custom encryption key arn
        /// </summary>
        [Output("encryptionKey")]
        public Output<string?> EncryptionKey { get; private set; } = null!;

        /// <summary>
        /// List of public endpoints.
        /// </summary>
        [Output("endpoints")]
        public Output<ImmutableArray<Outputs.ServiceEndpoint>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// IAM role used for accessing objects in s3.
        /// </summary>
        [Output("iamRole")]
        public Output<string> IamRole { get; private set; } = null!;

        /// <summary>
        /// When set to true the service is allowed to scale down to zero when idle. Always true for development services. Configurable only for 'production' services.
        /// </summary>
        [Output("idleScaling")]
        public Output<bool?> IdleScaling { get; private set; } = null!;

        /// <summary>
        /// Set minimum idling timeout (in minutes). Available only for 'production' services. Must be greater than or equal to 5 minutes.
        /// </summary>
        [Output("idleTimeoutMinutes")]
        public Output<int?> IdleTimeoutMinutes { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses allowed to access the service.
        /// </summary>
        [Output("ipAccesses")]
        public Output<ImmutableArray<Outputs.ServiceIpAccess>> IpAccesses { get; private set; } = null!;

        [Output("lastUpdated")]
        public Output<string> LastUpdated { get; private set; } = null!;

        /// <summary>
        /// Maximum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and lower than 360 for non paid services or 720 for paid services.
        /// </summary>
        [Output("maxTotalMemoryGb")]
        public Output<int?> MaxTotalMemoryGb { get; private set; } = null!;

        /// <summary>
        /// Minimum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and greater than 24.
        /// </summary>
        [Output("minTotalMemoryGb")]
        public Output<int?> MinTotalMemoryGb { get; private set; } = null!;

        /// <summary>
        /// User defined identifier for the service.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password for the default user. One of either `password` or `password_hash` must be specified.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// SHA256 hash of password for the default user. One of either `password` or `password_hash` must be specified.
        /// </summary>
        [Output("passwordHash")]
        public Output<string?> PasswordHash { get; private set; } = null!;

        /// <summary>
        /// Service config for private endpoints
        /// </summary>
        [Output("privateEndpointConfig")]
        public Output<Outputs.ServicePrivateEndpointConfig> PrivateEndpointConfig { get; private set; } = null!;

        /// <summary>
        /// List of private endpoint IDs
        /// </summary>
        [Output("privateEndpointIds")]
        public Output<ImmutableArray<string>> PrivateEndpointIds { get; private set; } = null!;

        /// <summary>
        /// Region within the cloud provider in which the service is deployed in.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Tier of the service: 'development', 'production'. Production services scale, Development are fixed size.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("clickhouse:index/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("clickhouse:index/service:Service", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-clickhouse",
                AdditionalSecretOutputs =
                {
                    "doubleSha1PasswordHash",
                    "password",
                    "passwordHash",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud provider ('aws', 'gcp', or 'azure') in which the service is deployed in.
        /// </summary>
        [Input("cloudProvider", required: true)]
        public Input<string> CloudProvider { get; set; } = null!;

        [Input("doubleSha1PasswordHash")]
        private Input<string>? _doubleSha1PasswordHash;

        /// <summary>
        /// Double SHA1 hash of password for connecting with the MySQL protocol. Cannot be specified if `password` is specified.
        /// </summary>
        public Input<string>? DoubleSha1PasswordHash
        {
            get => _doubleSha1PasswordHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _doubleSha1PasswordHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Custom role identifier arn
        /// </summary>
        [Input("encryptionAssumedRoleIdentifier")]
        public Input<string>? EncryptionAssumedRoleIdentifier { get; set; }

        /// <summary>
        /// Custom encryption key arn
        /// </summary>
        [Input("encryptionKey")]
        public Input<string>? EncryptionKey { get; set; }

        /// <summary>
        /// When set to true the service is allowed to scale down to zero when idle. Always true for development services. Configurable only for 'production' services.
        /// </summary>
        [Input("idleScaling")]
        public Input<bool>? IdleScaling { get; set; }

        /// <summary>
        /// Set minimum idling timeout (in minutes). Available only for 'production' services. Must be greater than or equal to 5 minutes.
        /// </summary>
        [Input("idleTimeoutMinutes")]
        public Input<int>? IdleTimeoutMinutes { get; set; }

        [Input("ipAccesses", required: true)]
        private InputList<Inputs.ServiceIpAccessArgs>? _ipAccesses;

        /// <summary>
        /// List of IP addresses allowed to access the service.
        /// </summary>
        public InputList<Inputs.ServiceIpAccessArgs> IpAccesses
        {
            get => _ipAccesses ?? (_ipAccesses = new InputList<Inputs.ServiceIpAccessArgs>());
            set => _ipAccesses = value;
        }

        /// <summary>
        /// Maximum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and lower than 360 for non paid services or 720 for paid services.
        /// </summary>
        [Input("maxTotalMemoryGb")]
        public Input<int>? MaxTotalMemoryGb { get; set; }

        /// <summary>
        /// Minimum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and greater than 24.
        /// </summary>
        [Input("minTotalMemoryGb")]
        public Input<int>? MinTotalMemoryGb { get; set; }

        /// <summary>
        /// User defined identifier for the service.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the default user. One of either `password` or `password_hash` must be specified.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("passwordHash")]
        private Input<string>? _passwordHash;

        /// <summary>
        /// SHA256 hash of password for the default user. One of either `password` or `password_hash` must be specified.
        /// </summary>
        public Input<string>? PasswordHash
        {
            get => _passwordHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwordHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateEndpointIds")]
        private InputList<string>? _privateEndpointIds;

        /// <summary>
        /// List of private endpoint IDs
        /// </summary>
        public InputList<string> PrivateEndpointIds
        {
            get => _privateEndpointIds ?? (_privateEndpointIds = new InputList<string>());
            set => _privateEndpointIds = value;
        }

        /// <summary>
        /// Region within the cloud provider in which the service is deployed in.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// Tier of the service: 'development', 'production'. Production services scale, Development are fixed size.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public ServiceArgs()
        {
        }
        public static new ServiceArgs Empty => new ServiceArgs();
    }

    public sealed class ServiceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud provider ('aws', 'gcp', or 'azure') in which the service is deployed in.
        /// </summary>
        [Input("cloudProvider")]
        public Input<string>? CloudProvider { get; set; }

        [Input("doubleSha1PasswordHash")]
        private Input<string>? _doubleSha1PasswordHash;

        /// <summary>
        /// Double SHA1 hash of password for connecting with the MySQL protocol. Cannot be specified if `password` is specified.
        /// </summary>
        public Input<string>? DoubleSha1PasswordHash
        {
            get => _doubleSha1PasswordHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _doubleSha1PasswordHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Custom role identifier arn
        /// </summary>
        [Input("encryptionAssumedRoleIdentifier")]
        public Input<string>? EncryptionAssumedRoleIdentifier { get; set; }

        /// <summary>
        /// Custom encryption key arn
        /// </summary>
        [Input("encryptionKey")]
        public Input<string>? EncryptionKey { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.ServiceEndpointGetArgs>? _endpoints;

        /// <summary>
        /// List of public endpoints.
        /// </summary>
        public InputList<Inputs.ServiceEndpointGetArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.ServiceEndpointGetArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// IAM role used for accessing objects in s3.
        /// </summary>
        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        /// <summary>
        /// When set to true the service is allowed to scale down to zero when idle. Always true for development services. Configurable only for 'production' services.
        /// </summary>
        [Input("idleScaling")]
        public Input<bool>? IdleScaling { get; set; }

        /// <summary>
        /// Set minimum idling timeout (in minutes). Available only for 'production' services. Must be greater than or equal to 5 minutes.
        /// </summary>
        [Input("idleTimeoutMinutes")]
        public Input<int>? IdleTimeoutMinutes { get; set; }

        [Input("ipAccesses")]
        private InputList<Inputs.ServiceIpAccessGetArgs>? _ipAccesses;

        /// <summary>
        /// List of IP addresses allowed to access the service.
        /// </summary>
        public InputList<Inputs.ServiceIpAccessGetArgs> IpAccesses
        {
            get => _ipAccesses ?? (_ipAccesses = new InputList<Inputs.ServiceIpAccessGetArgs>());
            set => _ipAccesses = value;
        }

        [Input("lastUpdated")]
        public Input<string>? LastUpdated { get; set; }

        /// <summary>
        /// Maximum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and lower than 360 for non paid services or 720 for paid services.
        /// </summary>
        [Input("maxTotalMemoryGb")]
        public Input<int>? MaxTotalMemoryGb { get; set; }

        /// <summary>
        /// Minimum total memory of all workers during auto-scaling in Gb. Available only for 'production' services. Must be a multiple of 12 and greater than 24.
        /// </summary>
        [Input("minTotalMemoryGb")]
        public Input<int>? MinTotalMemoryGb { get; set; }

        /// <summary>
        /// User defined identifier for the service.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the default user. One of either `password` or `password_hash` must be specified.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("passwordHash")]
        private Input<string>? _passwordHash;

        /// <summary>
        /// SHA256 hash of password for the default user. One of either `password` or `password_hash` must be specified.
        /// </summary>
        public Input<string>? PasswordHash
        {
            get => _passwordHash;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passwordHash = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Service config for private endpoints
        /// </summary>
        [Input("privateEndpointConfig")]
        public Input<Inputs.ServicePrivateEndpointConfigGetArgs>? PrivateEndpointConfig { get; set; }

        [Input("privateEndpointIds")]
        private InputList<string>? _privateEndpointIds;

        /// <summary>
        /// List of private endpoint IDs
        /// </summary>
        public InputList<string> PrivateEndpointIds
        {
            get => _privateEndpointIds ?? (_privateEndpointIds = new InputList<string>());
            set => _privateEndpointIds = value;
        }

        /// <summary>
        /// Region within the cloud provider in which the service is deployed in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Tier of the service: 'development', 'production'. Production services scale, Development are fixed size.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public ServiceState()
        {
        }
        public static new ServiceState Empty => new ServiceState();
    }
}