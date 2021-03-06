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
 * @interface GraphAllOfPeers
 */
export interface GraphAllOfPeers {
    /**
     * 
     * @type {string}
     * @memberof GraphAllOfPeers
     */
    peerId: string;
    /**
     * 
     * @type {string}
     * @memberof GraphAllOfPeers
     */
    agentVersion: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof GraphAllOfPeers
     */
    protocols: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof GraphAllOfPeers
     */
    distance: string;
    /**
     * 
     * @type {string}
     * @memberof GraphAllOfPeers
     */
    state?: GraphAllOfPeersStateEnum;
    /**
     * 
     * @type {string}
     * @memberof GraphAllOfPeers
     */
    firstInteractedAt: string;
    /**
     * 
     * @type {string}
     * @memberof GraphAllOfPeers
     */
    referredBy: string;
}

/**
* @export
* @enum {string}
*/
export enum GraphAllOfPeersStateEnum {
    Heard = 'HEARD',
    Waiting = 'WAITING',
    Queried = 'QUERIED',
    Unreachable = 'UNREACHABLE'
}

export function GraphAllOfPeersFromJSON(json: any): GraphAllOfPeers {
    return GraphAllOfPeersFromJSONTyped(json, false);
}

export function GraphAllOfPeersFromJSONTyped(json: any, ignoreDiscriminator: boolean): GraphAllOfPeers {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'peerId': json['peerId'],
        'agentVersion': json['agentVersion'],
        'protocols': json['protocols'],
        'distance': json['distance'],
        'state': !exists(json, 'state') ? undefined : json['state'],
        'firstInteractedAt': json['firstInteractedAt'],
        'referredBy': json['referredBy'],
    };
}

export function GraphAllOfPeersToJSON(value?: GraphAllOfPeers | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'peerId': value.peerId,
        'agentVersion': value.agentVersion,
        'protocols': value.protocols,
        'distance': value.distance,
        'state': value.state,
        'firstInteractedAt': value.firstInteractedAt,
        'referredBy': value.referredBy,
    };
}

