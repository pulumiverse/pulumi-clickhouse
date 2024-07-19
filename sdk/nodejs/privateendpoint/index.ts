// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetConfigArgs, GetConfigResult, GetConfigOutputArgs } from "./getConfig";
export const getConfig: typeof import("./getConfig").getConfig = null as any;
export const getConfigOutput: typeof import("./getConfig").getConfigOutput = null as any;
utilities.lazyLoad(exports, ["getConfig","getConfigOutput"], () => require("./getConfig"));

export { RegistrationArgs, RegistrationState } from "./registration";
export type Registration = import("./registration").Registration;
export const Registration: typeof import("./registration").Registration = null as any;
utilities.lazyLoad(exports, ["Registration"], () => require("./registration"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "clickhouse:PrivateEndpoint/registration:Registration":
                return new Registration(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("clickhouse", "PrivateEndpoint/registration", _module)