// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `astra.StreamingTopic` creates an Astra Streaming topic.
 */
export class StreamingTopic extends pulumi.CustomResource {
    /**
     * Get an existing StreamingTopic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamingTopicState, opts?: pulumi.CustomResourceOptions): StreamingTopic {
        return new StreamingTopic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'astra:index/streamingTopic:StreamingTopic';

    /**
     * Returns true if the given object is an instance of StreamingTopic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamingTopic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamingTopic.__pulumiType;
    }

    /**
     * Cloud provider
     */
    public readonly cloudProvider!: pulumi.Output<string>;
    /**
     * Whether or not to allow Terraform to destroy this streaming topic. Unless this field is set to false in Terraform state,
     * a `terraform destroy` or `terraform apply` command that deletes the instance will fail. Defaults to `true`.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Pulsar Namespace
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * cloud region
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Streaming tenant name.
     */
    public readonly tenantName!: pulumi.Output<string>;
    /**
     * Streaming tenant topic.
     */
    public readonly topic!: pulumi.Output<string>;

    /**
     * Create a StreamingTopic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamingTopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamingTopicArgs | StreamingTopicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StreamingTopicState | undefined;
            resourceInputs["cloudProvider"] = state ? state.cloudProvider : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantName"] = state ? state.tenantName : undefined;
            resourceInputs["topic"] = state ? state.topic : undefined;
        } else {
            const args = argsOrState as StreamingTopicArgs | undefined;
            if ((!args || args.cloudProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudProvider'");
            }
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.tenantName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantName'");
            }
            if ((!args || args.topic === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topic'");
            }
            resourceInputs["cloudProvider"] = args ? args.cloudProvider : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantName"] = args ? args.tenantName : undefined;
            resourceInputs["topic"] = args ? args.topic : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StreamingTopic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StreamingTopic resources.
 */
export interface StreamingTopicState {
    /**
     * Cloud provider
     */
    cloudProvider?: pulumi.Input<string>;
    /**
     * Whether or not to allow Terraform to destroy this streaming topic. Unless this field is set to false in Terraform state,
     * a `terraform destroy` or `terraform apply` command that deletes the instance will fail. Defaults to `true`.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Pulsar Namespace
     */
    namespace?: pulumi.Input<string>;
    /**
     * cloud region
     */
    region?: pulumi.Input<string>;
    /**
     * Streaming tenant name.
     */
    tenantName?: pulumi.Input<string>;
    /**
     * Streaming tenant topic.
     */
    topic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StreamingTopic resource.
 */
export interface StreamingTopicArgs {
    /**
     * Cloud provider
     */
    cloudProvider: pulumi.Input<string>;
    /**
     * Whether or not to allow Terraform to destroy this streaming topic. Unless this field is set to false in Terraform state,
     * a `terraform destroy` or `terraform apply` command that deletes the instance will fail. Defaults to `true`.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Pulsar Namespace
     */
    namespace: pulumi.Input<string>;
    /**
     * cloud region
     */
    region: pulumi.Input<string>;
    /**
     * Streaming tenant name.
     */
    tenantName: pulumi.Input<string>;
    /**
     * Streaming tenant topic.
     */
    topic: pulumi.Input<string>;
}
