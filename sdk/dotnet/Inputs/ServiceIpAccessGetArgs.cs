// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Clickhouse.Inputs
{

    public sealed class ServiceIpAccessGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the IP address.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IP address allowed to access the service. In case you want to set the ip_access to anywhere you should set source to 0.0.0.0/0
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public ServiceIpAccessGetArgs()
        {
        }
        public static new ServiceIpAccessGetArgs Empty => new ServiceIpAccessGetArgs();
    }
}
