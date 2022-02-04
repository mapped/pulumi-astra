// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Astra.Astra
{
    [AstraResourceType("index:astra/privateLink:PrivateLink")]
    public partial class PrivateLink : Pulumi.CustomResource
    {
        /// <summary>
        /// List of service principals to apply to the Private Link (i.e. arn:aws:iam::123456789012:role/admin).
        /// </summary>
        [Output("allowedPrincipals")]
        public Output<ImmutableArray<string>> AllowedPrincipals { get; private set; } = null!;

        /// <summary>
        /// Astra database where private link will be enabled.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// Astra datacenter in the region where the private link will be created.
        /// </summary>
        [Output("datacenterId")]
        public Output<string> DatacenterId { get; private set; } = null!;

        /// <summary>
        /// Name of the endpoint service for private link generated by the cloud provider.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateLink(string name, PrivateLinkArgs args, CustomResourceOptions? options = null)
            : base("index:astra/privateLink:PrivateLink", name, args ?? new PrivateLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateLink(string name, Input<string> id, PrivateLinkState? state = null, CustomResourceOptions? options = null)
            : base("index:astra/privateLink:PrivateLink", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateLink Get(string name, Input<string> id, PrivateLinkState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateLink(name, id, state, options);
        }
    }

    public sealed class PrivateLinkArgs : Pulumi.ResourceArgs
    {
        [Input("allowedPrincipals", required: true)]
        private InputList<string>? _allowedPrincipals;

        /// <summary>
        /// List of service principals to apply to the Private Link (i.e. arn:aws:iam::123456789012:role/admin).
        /// </summary>
        public InputList<string> AllowedPrincipals
        {
            get => _allowedPrincipals ?? (_allowedPrincipals = new InputList<string>());
            set => _allowedPrincipals = value;
        }

        /// <summary>
        /// Astra database where private link will be enabled.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        /// <summary>
        /// Astra datacenter in the region where the private link will be created.
        /// </summary>
        [Input("datacenterId", required: true)]
        public Input<string> DatacenterId { get; set; } = null!;

        public PrivateLinkArgs()
        {
        }
    }

    public sealed class PrivateLinkState : Pulumi.ResourceArgs
    {
        [Input("allowedPrincipals")]
        private InputList<string>? _allowedPrincipals;

        /// <summary>
        /// List of service principals to apply to the Private Link (i.e. arn:aws:iam::123456789012:role/admin).
        /// </summary>
        public InputList<string> AllowedPrincipals
        {
            get => _allowedPrincipals ?? (_allowedPrincipals = new InputList<string>());
            set => _allowedPrincipals = value;
        }

        /// <summary>
        /// Astra database where private link will be enabled.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// Astra datacenter in the region where the private link will be created.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// Name of the endpoint service for private link generated by the cloud provider.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public PrivateLinkState()
        {
        }
    }
}
