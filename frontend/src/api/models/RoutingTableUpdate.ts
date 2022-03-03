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
import {
    RoutingTablePeer | Array | string,
    RoutingTablePeer | Array | stringFromJSON,
    RoutingTablePeer | Array | stringFromJSONTyped,
    RoutingTablePeer | Array | stringToJSON,
} from './RoutingTablePeer | Array | string';

/**
 * 
 * @export
 * @interface RoutingTableUpdate
 */
export interface RoutingTableUpdate {
    /**
     * 
     * @type {string}
     * @memberof RoutingTableUpdate
     */
    type: RoutingTableUpdateTypeEnum;
    /**
     * 
     * @type {RoutingTablePeer | Array | string}
     * @memberof RoutingTableUpdate
     */
    update: RoutingTablePeer | Array | string | null;
}

/**
* @export
* @enum {string}
*/
export enum RoutingTableUpdateTypeEnum {
    Full = 'FULL',
    PeerAdded = 'PEER_ADDED',
    PeerRemoved = 'PEER_REMOVED'
}

export function RoutingTableUpdateFromJSON(json: any): RoutingTableUpdate {
    return RoutingTableUpdateFromJSONTyped(json, false);
}

export function RoutingTableUpdateFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoutingTableUpdate {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'type': json['type'],
        'update': RoutingTablePeer | Array | stringFromJSON(json['update']),
    };
}

export function RoutingTableUpdateToJSON(value?: RoutingTableUpdate | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'type': value.type,
        'update': RoutingTablePeer | Array | stringToJSON(value.update),
    };
}
