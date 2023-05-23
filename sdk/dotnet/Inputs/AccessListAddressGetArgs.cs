// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Astra.Inputs
{

    public sealed class AccessListAddressGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP Address/CIDR group that should have access
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// Description for the IP Address/CIDR group
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable/disable this IP Address/CIDR group's access
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public AccessListAddressGetArgs()
        {
        }
        public static new AccessListAddressGetArgs Empty => new AccessListAddressGetArgs();
    }
}
