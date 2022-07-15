# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AccessListArgs', 'AccessList']

@pulumi.input_type
class AccessListArgs:
    def __init__(__self__, *,
                 addresses: pulumi.Input[Sequence[pulumi.Input['AccessListAddressArgs']]],
                 database_id: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AccessList resource.
        :param pulumi.Input[Sequence[pulumi.Input['AccessListAddressArgs']]] addresses: List of address requests that should have access to database endpoints.
        :param pulumi.Input[str] database_id: The ID of the Astra database.
        :param pulumi.Input[bool] enabled: Public access restrictions enabled or disabled
        """
        pulumi.set(__self__, "addresses", addresses)
        pulumi.set(__self__, "database_id", database_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def addresses(self) -> pulumi.Input[Sequence[pulumi.Input['AccessListAddressArgs']]]:
        """
        List of address requests that should have access to database endpoints.
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: pulumi.Input[Sequence[pulumi.Input['AccessListAddressArgs']]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Input[str]:
        """
        The ID of the Astra database.
        """
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Public access restrictions enabled or disabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class _AccessListState:
    def __init__(__self__, *,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input['AccessListAddressArgs']]]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AccessList resources.
        :param pulumi.Input[Sequence[pulumi.Input['AccessListAddressArgs']]] addresses: List of address requests that should have access to database endpoints.
        :param pulumi.Input[str] database_id: The ID of the Astra database.
        :param pulumi.Input[bool] enabled: Public access restrictions enabled or disabled
        """
        if addresses is not None:
            pulumi.set(__self__, "addresses", addresses)
        if database_id is not None:
            pulumi.set(__self__, "database_id", database_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccessListAddressArgs']]]]:
        """
        List of address requests that should have access to database endpoints.
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccessListAddressArgs']]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Astra database.
        """
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Public access restrictions enabled or disabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


class AccessList(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessListAddressArgs']]]]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        `AccessList` resource represents a database access list, used to limit the ip's / CIDR groups that have access to a database.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_astra as astra

        example = astra.AccessList("example",
            addresses=[
                astra.AccessListAddressArgs(
                    address="0.0.0.1/0",
                    enabled=True,
                ),
                astra.AccessListAddressArgs(
                    address="0.0.0.2/0",
                    enabled=True,
                ),
                astra.AccessListAddressArgs(
                    address="0.0.0.3/0",
                    enabled=True,
                ),
            ],
            database_id="a6bc9c26-e7ce-424f-84c7-0a00afb12588",
            enabled=True)
        ```

        ## Import

        # the import id should be the database_id.

        ```sh
         $ pulumi import astra:index/accessList:AccessList example a6bc9c26-e7ce-424f-84c7-0a00afb12588
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessListAddressArgs']]]] addresses: List of address requests that should have access to database endpoints.
        :param pulumi.Input[str] database_id: The ID of the Astra database.
        :param pulumi.Input[bool] enabled: Public access restrictions enabled or disabled
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessListArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `AccessList` resource represents a database access list, used to limit the ip's / CIDR groups that have access to a database.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_astra as astra

        example = astra.AccessList("example",
            addresses=[
                astra.AccessListAddressArgs(
                    address="0.0.0.1/0",
                    enabled=True,
                ),
                astra.AccessListAddressArgs(
                    address="0.0.0.2/0",
                    enabled=True,
                ),
                astra.AccessListAddressArgs(
                    address="0.0.0.3/0",
                    enabled=True,
                ),
            ],
            database_id="a6bc9c26-e7ce-424f-84c7-0a00afb12588",
            enabled=True)
        ```

        ## Import

        # the import id should be the database_id.

        ```sh
         $ pulumi import astra:index/accessList:AccessList example a6bc9c26-e7ce-424f-84c7-0a00afb12588
        ```

        :param str resource_name: The name of the resource.
        :param AccessListArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessListArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessListAddressArgs']]]]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessListArgs.__new__(AccessListArgs)

            if addresses is None and not opts.urn:
                raise TypeError("Missing required property 'addresses'")
            __props__.__dict__["addresses"] = addresses
            if database_id is None and not opts.urn:
                raise TypeError("Missing required property 'database_id'")
            __props__.__dict__["database_id"] = database_id
            __props__.__dict__["enabled"] = enabled
        super(AccessList, __self__).__init__(
            'astra:index/accessList:AccessList',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            addresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessListAddressArgs']]]]] = None,
            database_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None) -> 'AccessList':
        """
        Get an existing AccessList resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessListAddressArgs']]]] addresses: List of address requests that should have access to database endpoints.
        :param pulumi.Input[str] database_id: The ID of the Astra database.
        :param pulumi.Input[bool] enabled: Public access restrictions enabled or disabled
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessListState.__new__(_AccessListState)

        __props__.__dict__["addresses"] = addresses
        __props__.__dict__["database_id"] = database_id
        __props__.__dict__["enabled"] = enabled
        return AccessList(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def addresses(self) -> pulumi.Output[Sequence['outputs.AccessListAddress']]:
        """
        List of address requests that should have access to database endpoints.
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Output[str]:
        """
        The ID of the Astra database.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Public access restrictions enabled or disabled
        """
        return pulumi.get(self, "enabled")

