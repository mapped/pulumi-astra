// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Keyspace extends pulumi.CustomResource {
    /**
     * Get an existing Keyspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyspaceState, opts?: pulumi.CustomResourceOptions): Keyspace {
        return new Keyspace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'index:astra/keyspace:Keyspace';

    /**
     * Returns true if the given object is an instance of Keyspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Keyspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Keyspace.__pulumiType;
    }

    /**
     * Astra database to create the keyspace.
     */
    public readonly databaseId!: pulumi.Output<string>;
    /**
     * Keyspace name can have up to 48 alpha-numeric characters and contain underscores; only letters and numbers are supported
     * as the first character.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Keyspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyspaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyspaceArgs | KeyspaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyspaceState | undefined;
            resourceInputs["databaseId"] = state ? state.databaseId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as KeyspaceArgs | undefined;
            if ((!args || args.databaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseId'");
            }
            resourceInputs["databaseId"] = args ? args.databaseId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Keyspace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Keyspace resources.
 */
export interface KeyspaceState {
    /**
     * Astra database to create the keyspace.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * Keyspace name can have up to 48 alpha-numeric characters and contain underscores; only letters and numbers are supported
     * as the first character.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Keyspace resource.
 */
export interface KeyspaceArgs {
    /**
     * Astra database to create the keyspace.
     */
    databaseId: pulumi.Input<string>;
    /**
     * Keyspace name can have up to 48 alpha-numeric characters and contain underscores; only letters and numbers are supported
     * as the first character.
     */
    name?: pulumi.Input<string>;
}
