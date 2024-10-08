# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('clickhouse')


class _ExportableConfig(types.ModuleType):
    @property
    def api_url(self) -> Optional[str]:
        """
        API URL of the ClickHouse OpenAPI the provider will interact with. Alternatively, can be configured using the
        `CLICKHOUSE_API_URL` environment variable. Only specify if you have a specific deployment of the ClickHouse OpenAPI you
        want to run against.
        """
        return __config__.get('apiUrl')

    @property
    def organization_id(self) -> Optional[str]:
        """
        ID of the organization the provider will create services under. Alternatively, can be configured using the
        `CLICKHOUSE_ORG_ID` environment variable.
        """
        return __config__.get('organizationId')

    @property
    def token_key(self) -> Optional[str]:
        """
        Token key of the key/secret pair. Used to authenticate with OpenAPI. Alternatively, can be configured using the
        `CLICKHOUSE_TOKEN_KEY` environment variable.
        """
        return __config__.get('tokenKey')

    @property
    def token_secret(self) -> Optional[str]:
        """
        Token secret of the key/secret pair. Used to authenticate with OpenAPI. Alternatively, can be configured using the
        `CLICKHOUSE_TOKEN_SECRET` environment variable.
        """
        return __config__.get('tokenSecret')

