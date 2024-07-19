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
    /// <summary>
    /// The provider type for the clickhouse package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [ClickhouseResourceType("pulumi:providers:clickhouse")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// API URL of the ClickHouse OpenAPI the provider will interact with. Alternatively, can be configured using the
        /// `CLICKHOUSE_API_URL` environment variable. Only specify if you have a specific deployment of the ClickHouse OpenAPI you
        /// want to run against.
        /// </summary>
        [Output("apiUrl")]
        public Output<string?> ApiUrl { get; private set; } = null!;

        /// <summary>
        /// ID of the organization the provider will create services under. Alternatively, can be configured using the
        /// `CLICKHOUSE_ORG_ID` environment variable.
        /// </summary>
        [Output("organizationId")]
        public Output<string?> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Token key of the key/secret pair. Used to authenticate with OpenAPI. Alternatively, can be configured using the
        /// `CLICKHOUSE_TOKEN_KEY` environment variable.
        /// </summary>
        [Output("tokenKey")]
        public Output<string?> TokenKey { get; private set; } = null!;

        /// <summary>
        /// Token secret of the key/secret pair. Used to authenticate with OpenAPI. Alternatively, can be configured using the
        /// `CLICKHOUSE_TOKEN_SECRET` environment variable.
        /// </summary>
        [Output("tokenSecret")]
        public Output<string?> TokenSecret { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("clickhouse", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
                    "tokenSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API URL of the ClickHouse OpenAPI the provider will interact with. Alternatively, can be configured using the
        /// `CLICKHOUSE_API_URL` environment variable. Only specify if you have a specific deployment of the ClickHouse OpenAPI you
        /// want to run against.
        /// </summary>
        [Input("apiUrl")]
        public Input<string>? ApiUrl { get; set; }

        /// <summary>
        /// ID of the organization the provider will create services under. Alternatively, can be configured using the
        /// `CLICKHOUSE_ORG_ID` environment variable.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Token key of the key/secret pair. Used to authenticate with OpenAPI. Alternatively, can be configured using the
        /// `CLICKHOUSE_TOKEN_KEY` environment variable.
        /// </summary>
        [Input("tokenKey")]
        public Input<string>? TokenKey { get; set; }

        [Input("tokenSecret")]
        private Input<string>? _tokenSecret;

        /// <summary>
        /// Token secret of the key/secret pair. Used to authenticate with OpenAPI. Alternatively, can be configured using the
        /// `CLICKHOUSE_TOKEN_SECRET` environment variable.
        /// </summary>
        public Input<string>? TokenSecret
        {
            get => _tokenSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tokenSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
