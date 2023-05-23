// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Astra
{
    public static class GetSecureConnectBundleUrl
    {
        /// <summary>
        /// `astra.getSecureConnectBundleUrl` provides a datasource that generates a temporary secure connect bundle URL. This URL lasts five minutes. Secure connect bundles are used to connect to Astra using cql cassandra drivers. See the [docs](https://docs.datastax.com/en/astra/docs/connecting-to-database.html) for more information on how to connect.
        /// </summary>
        public static Task<GetSecureConnectBundleUrlResult> InvokeAsync(GetSecureConnectBundleUrlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecureConnectBundleUrlResult>("astra:index/getSecureConnectBundleUrl:getSecureConnectBundleUrl", args ?? new GetSecureConnectBundleUrlArgs(), options.WithDefaults());

        /// <summary>
        /// `astra.getSecureConnectBundleUrl` provides a datasource that generates a temporary secure connect bundle URL. This URL lasts five minutes. Secure connect bundles are used to connect to Astra using cql cassandra drivers. See the [docs](https://docs.datastax.com/en/astra/docs/connecting-to-database.html) for more information on how to connect.
        /// </summary>
        public static Output<GetSecureConnectBundleUrlResult> Invoke(GetSecureConnectBundleUrlInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecureConnectBundleUrlResult>("astra:index/getSecureConnectBundleUrl:getSecureConnectBundleUrl", args ?? new GetSecureConnectBundleUrlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecureConnectBundleUrlArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Astra database.
        /// </summary>
        [Input("databaseId", required: true)]
        public string DatabaseId { get; set; } = null!;

        /// <summary>
        /// The ID of the Astra datacenter. If omitted, all bundles will be fetched.
        /// </summary>
        [Input("datacenterId")]
        public string? DatacenterId { get; set; }

        public GetSecureConnectBundleUrlArgs()
        {
        }
        public static new GetSecureConnectBundleUrlArgs Empty => new GetSecureConnectBundleUrlArgs();
    }

    public sealed class GetSecureConnectBundleUrlInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Astra database.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        /// <summary>
        /// The ID of the Astra datacenter. If omitted, all bundles will be fetched.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        public GetSecureConnectBundleUrlInvokeArgs()
        {
        }
        public static new GetSecureConnectBundleUrlInvokeArgs Empty => new GetSecureConnectBundleUrlInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecureConnectBundleUrlResult
    {
        /// <summary>
        /// The ID of the Astra database.
        /// </summary>
        public readonly string DatabaseId;
        /// <summary>
        /// The ID of the Astra datacenter. If omitted, all bundles will be fetched.
        /// </summary>
        public readonly string? DatacenterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Secure Connect Bundle info
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSecureConnectBundleUrlSecureBundleResult> SecureBundles;

        [OutputConstructor]
        private GetSecureConnectBundleUrlResult(
            string databaseId,

            string? datacenterId,

            string id,

            ImmutableArray<Outputs.GetSecureConnectBundleUrlSecureBundleResult> secureBundles)
        {
            DatabaseId = databaseId;
            DatacenterId = datacenterId;
            Id = id;
            SecureBundles = secureBundles;
        }
    }
}
