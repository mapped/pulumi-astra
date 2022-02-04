// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accessList";
export * from "./database";
export * from "./getAccessList";
export * from "./getAstraDatabase";
export * from "./getAstraDatabases";
export * from "./getAvailableRegions";
export * from "./getKeyspace";
export * from "./getKeyspaces";
export * from "./getPrivateLinkEndpoints";
export * from "./getPrivateLinks";
export * from "./getRoles";
export * from "./getSecureConnectBundleUrl";
export * from "./keyspace";
export * from "./privateLink";
export * from "./privateLinkEndpoint";
export * from "./role";
export * from "./token";

// Import resources to register:
import { AccessList } from "./accessList";
import { Database } from "./database";
import { Keyspace } from "./keyspace";
import { PrivateLink } from "./privateLink";
import { PrivateLinkEndpoint } from "./privateLinkEndpoint";
import { Role } from "./role";
import { Token } from "./token";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "index:astra/accessList:AccessList":
                return new AccessList(name, <any>undefined, { urn })
            case "index:astra/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "index:astra/keyspace:Keyspace":
                return new Keyspace(name, <any>undefined, { urn })
            case "index:astra/privateLink:PrivateLink":
                return new PrivateLink(name, <any>undefined, { urn })
            case "index:astra/privateLinkEndpoint:PrivateLinkEndpoint":
                return new PrivateLinkEndpoint(name, <any>undefined, { urn })
            case "index:astra/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "index:astra/token:Token":
                return new Token(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("astra", "astra/accessList", _module)
pulumi.runtime.registerResourceModule("astra", "astra/database", _module)
pulumi.runtime.registerResourceModule("astra", "astra/keyspace", _module)
pulumi.runtime.registerResourceModule("astra", "astra/privateLink", _module)
pulumi.runtime.registerResourceModule("astra", "astra/privateLinkEndpoint", _module)
pulumi.runtime.registerResourceModule("astra", "astra/role", _module)
pulumi.runtime.registerResourceModule("astra", "astra/token", _module)
