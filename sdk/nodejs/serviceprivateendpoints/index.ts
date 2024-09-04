// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AttachmentArgs, AttachmentState } from "./attachment";
export type Attachment = import("./attachment").Attachment;
export const Attachment: typeof import("./attachment").Attachment = null as any;
utilities.lazyLoad(exports, ["Attachment"], () => require("./attachment"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "clickhouse:ServicePrivateEndpoints/attachment:Attachment":
                return new Attachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("clickhouse", "ServicePrivateEndpoints/attachment", _module)