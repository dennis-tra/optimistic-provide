/* tslint:disable */
/* eslint-disable */
/**
 * OptProv API
 * This is the REST API to interact and control with libp2p hosts. All responses not in the HTTP status code range [200,300) return the error object below. 
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface RoutingTableEntry
 */
export interface RoutingTableEntry {
    /**
     * 
     * @type {string}
     * @memberof RoutingTableEntry
     */
    peerId: string;
    /**
     * 
     * @type {number}
     * @memberof RoutingTableEntry
     */
    bucket: number;
    /**
     * 
     * @type {string}
     * @memberof RoutingTableEntry
     */
    addedAt: string;
    /**
     * 
     * @type {string}
     * @memberof RoutingTableEntry
     */
    connectedSince: string | null;
    /**
     * 
     * @type {string}
     * @memberof RoutingTableEntry
     */
    lastUsefulAt: string | null;
    /**
     * 
     * @type {string}
     * @memberof RoutingTableEntry
     */
    lastSuccessfulOutboundQueryAt: string;
}

export function RoutingTableEntryFromJSON(json: any): RoutingTableEntry {
    return RoutingTableEntryFromJSONTyped(json, false);
}

export function RoutingTableEntryFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoutingTableEntry {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'peerId': json['peerId'],
        'bucket': json['bucket'],
        'addedAt': json['addedAt'],
        'connectedSince': json['connectedSince'],
        'lastUsefulAt': json['lastUsefulAt'],
        'lastSuccessfulOutboundQueryAt': json['lastSuccessfulOutboundQueryAt'],
    };
}

export function RoutingTableEntryToJSON(value?: RoutingTableEntry | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'peerId': value.peerId,
        'bucket': value.bucket,
        'addedAt': value.addedAt,
        'connectedSince': value.connectedSince,
        'lastUsefulAt': value.lastUsefulAt,
        'lastSuccessfulOutboundQueryAt': value.lastSuccessfulOutboundQueryAt,
    };
}

