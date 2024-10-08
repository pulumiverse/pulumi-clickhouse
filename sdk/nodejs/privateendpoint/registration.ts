// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ClickHouse Cloud provides the ability to connect your services to your cloud virtual network through a feature named *Private Link*.
 *
 * You can use the *private_endpoint_registration* resource to set up the private link feature.
 *
 * Check the [docs](https://clickhouse.com/docs/en/cloud/security/private-link-overview) for more details on *private link*.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as clickhouse from "@pulumiverse/clickhouse";
 *
 * const endpoint = new clickhouse.privateendpoint.Registration("endpoint", {
 *     cloudProvider: "aws",
 *     description: "Private Link from VPC foo",
 *     privateEndpointId: "vpce-...",
 *     region: "us-west-2",
 * });
 * ```
 *
 * ## Import
 *
 * Endpoint Attachments can be imported by specifying the Cloud provider private endpoint ID
 *
 * For example for AWS you could run:
 *
 * ```sh
 * $ pulumi import clickhouse:PrivateEndpoint/registration:Registration example vpce-xxxxxx
 * ```
 */
export class Registration extends pulumi.CustomResource {
    /**
     * Get an existing Registration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistrationState, opts?: pulumi.CustomResourceOptions): Registration {
        return new Registration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'clickhouse:PrivateEndpoint/registration:Registration';

    /**
     * Returns true if the given object is an instance of Registration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registration.__pulumiType;
    }

    /**
     * Cloud provider of the private endpoint ID
     */
    public readonly cloudProvider!: pulumi.Output<string>;
    /**
     * Description of the private endpoint
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ID of the private endpoint (replaces deprecated attribute `id`)
     */
    public readonly privateEndpointId!: pulumi.Output<string>;
    /**
     * Region of the private endpoint
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a Registration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistrationArgs | RegistrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegistrationState | undefined;
            resourceInputs["cloudProvider"] = state ? state.cloudProvider : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["privateEndpointId"] = state ? state.privateEndpointId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as RegistrationArgs | undefined;
            if ((!args || args.cloudProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudProvider'");
            }
            if ((!args || args.privateEndpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateEndpointId'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["cloudProvider"] = args ? args.cloudProvider : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["privateEndpointId"] = args ? args.privateEndpointId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Registration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Registration resources.
 */
export interface RegistrationState {
    /**
     * Cloud provider of the private endpoint ID
     */
    cloudProvider?: pulumi.Input<string>;
    /**
     * Description of the private endpoint
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the private endpoint (replaces deprecated attribute `id`)
     */
    privateEndpointId?: pulumi.Input<string>;
    /**
     * Region of the private endpoint
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Registration resource.
 */
export interface RegistrationArgs {
    /**
     * Cloud provider of the private endpoint ID
     */
    cloudProvider: pulumi.Input<string>;
    /**
     * Description of the private endpoint
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the private endpoint (replaces deprecated attribute `id`)
     */
    privateEndpointId: pulumi.Input<string>;
    /**
     * Region of the private endpoint
     */
    region: pulumi.Input<string>;
}
