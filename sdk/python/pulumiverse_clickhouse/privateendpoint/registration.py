# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RegistrationArgs', 'Registration']

@pulumi.input_type
class RegistrationArgs:
    def __init__(__self__, *,
                 cloud_provider: pulumi.Input[str],
                 region: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Registration resource.
        :param pulumi.Input[str] cloud_provider: Cloud provider of the private endpoint ID
        :param pulumi.Input[str] region: Region of the private endpoint
        :param pulumi.Input[str] description: Description of the private endpoint
        """
        pulumi.set(__self__, "cloud_provider", cloud_provider)
        pulumi.set(__self__, "region", region)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> pulumi.Input[str]:
        """
        Cloud provider of the private endpoint ID
        """
        return pulumi.get(self, "cloud_provider")

    @cloud_provider.setter
    def cloud_provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloud_provider", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Region of the private endpoint
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the private endpoint
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _RegistrationState:
    def __init__(__self__, *,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Registration resources.
        :param pulumi.Input[str] cloud_provider: Cloud provider of the private endpoint ID
        :param pulumi.Input[str] description: Description of the private endpoint
        :param pulumi.Input[str] region: Region of the private endpoint
        """
        if cloud_provider is not None:
            pulumi.set(__self__, "cloud_provider", cloud_provider)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud provider of the private endpoint ID
        """
        return pulumi.get(self, "cloud_provider")

    @cloud_provider.setter
    def cloud_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_provider", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the private endpoint
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region of the private endpoint
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Registration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Registration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_provider: Cloud provider of the private endpoint ID
        :param pulumi.Input[str] description: Description of the private endpoint
        :param pulumi.Input[str] region: Region of the private endpoint
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegistrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Registration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RegistrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistrationArgs.__new__(RegistrationArgs)

            if cloud_provider is None and not opts.urn:
                raise TypeError("Missing required property 'cloud_provider'")
            __props__.__dict__["cloud_provider"] = cloud_provider
            __props__.__dict__["description"] = description
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
        super(Registration, __self__).__init__(
            'clickhouse:PrivateEndpoint/registration:Registration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cloud_provider: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Registration':
        """
        Get an existing Registration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_provider: Cloud provider of the private endpoint ID
        :param pulumi.Input[str] description: Description of the private endpoint
        :param pulumi.Input[str] region: Region of the private endpoint
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegistrationState.__new__(_RegistrationState)

        __props__.__dict__["cloud_provider"] = cloud_provider
        __props__.__dict__["description"] = description
        __props__.__dict__["region"] = region
        return Registration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> pulumi.Output[str]:
        """
        Cloud provider of the private endpoint ID
        """
        return pulumi.get(self, "cloud_provider")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the private endpoint
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Region of the private endpoint
        """
        return pulumi.get(self, "region")
