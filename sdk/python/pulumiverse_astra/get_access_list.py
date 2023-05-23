# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetAccessListResult',
    'AwaitableGetAccessListResult',
    'get_access_list',
    'get_access_list_output',
]

@pulumi.output_type
class GetAccessListResult:
    """
    A collection of values returned by getAccessList.
    """
    def __init__(__self__, addresses=None, database_id=None, enabled=None, id=None):
        if addresses and not isinstance(addresses, list):
            raise TypeError("Expected argument 'addresses' to be a list")
        pulumi.set(__self__, "addresses", addresses)
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence['outputs.GetAccessListAddressResult']:
        """
        Addresses in the access list.
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> str:
        """
        The ID of the Astra database.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        The Access list is enabled or disabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetAccessListResult(GetAccessListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessListResult(
            addresses=self.addresses,
            database_id=self.database_id,
            enabled=self.enabled,
            id=self.id)


def get_access_list(database_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessListResult:
    """
    `AccessList` provides a datasource that lists the access lists for an Astra database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astra as astra

    dev = astra.get_access_list(database_id="8d356587-73b3-430a-9c0e-d780332e2afb")
    ```


    :param str database_id: The ID of the Astra database.
    """
    __args__ = dict()
    __args__['databaseId'] = database_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('astra:index/getAccessList:getAccessList', __args__, opts=opts, typ=GetAccessListResult).value

    return AwaitableGetAccessListResult(
        addresses=__ret__.addresses,
        database_id=__ret__.database_id,
        enabled=__ret__.enabled,
        id=__ret__.id)


@_utilities.lift_output_func(get_access_list)
def get_access_list_output(database_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccessListResult]:
    """
    `AccessList` provides a datasource that lists the access lists for an Astra database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_astra as astra

    dev = astra.get_access_list(database_id="8d356587-73b3-430a-9c0e-d780332e2afb")
    ```


    :param str database_id: The ID of the Astra database.
    """
    ...
