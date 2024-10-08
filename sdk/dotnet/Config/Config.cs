// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumiverse.Clickhouse
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("clickhouse");

        private static readonly __Value<string?> _apiUrl = new __Value<string?>(() => __config.Get("apiUrl"));
        /// <summary>
        /// API URL of the ClickHouse OpenAPI the provider will interact with. Alternatively, can be configured using the
        /// `CLICKHOUSE_API_URL` environment variable. Only specify if you have a specific deployment of the ClickHouse OpenAPI you
        /// want to run against.
        /// </summary>
        public static string? ApiUrl
        {
            get => _apiUrl.Get();
            set => _apiUrl.Set(value);
        }

        private static readonly __Value<string?> _organizationId = new __Value<string?>(() => __config.Get("organizationId"));
        /// <summary>
        /// ID of the organization the provider will create services under. Alternatively, can be configured using the
        /// `CLICKHOUSE_ORG_ID` environment variable.
        /// </summary>
        public static string? OrganizationId
        {
            get => _organizationId.Get();
            set => _organizationId.Set(value);
        }

        private static readonly __Value<string?> _tokenKey = new __Value<string?>(() => __config.Get("tokenKey"));
        /// <summary>
        /// Token key of the key/secret pair. Used to authenticate with OpenAPI. Alternatively, can be configured using the
        /// `CLICKHOUSE_TOKEN_KEY` environment variable.
        /// </summary>
        public static string? TokenKey
        {
            get => _tokenKey.Get();
            set => _tokenKey.Set(value);
        }

        private static readonly __Value<string?> _tokenSecret = new __Value<string?>(() => __config.Get("tokenSecret"));
        /// <summary>
        /// Token secret of the key/secret pair. Used to authenticate with OpenAPI. Alternatively, can be configured using the
        /// `CLICKHOUSE_TOKEN_SECRET` environment variable.
        /// </summary>
        public static string? TokenSecret
        {
            get => _tokenSecret.Get();
            set => _tokenSecret.Set(value);
        }

    }
}
