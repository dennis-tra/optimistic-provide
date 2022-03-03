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
 * @interface FindNodeCloserPeer
 */
export interface FindNodeCloserPeer {
    /**
     * 
     * @type {string}
     * @memberof FindNodeCloserPeer
     */
    peerId: string;
    /**
     * How far away is this peer from the desired hash.
     * @type {string}
     * @memberof FindNodeCloserPeer
     */
    distance: string;
    /**
     * In which bucket was this peer of the remote peers routing table
     * @type {number}
     * @memberof FindNodeCloserPeer
     */
    bucket: number;
}

export function FindNodeCloserPeerFromJSON(json: any): FindNodeCloserPeer {
    return FindNodeCloserPeerFromJSONTyped(json, false);
}

export function FindNodeCloserPeerFromJSONTyped(json: any, ignoreDiscriminator: boolean): FindNodeCloserPeer {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'peerId': json['peerId'],
        'distance': json['distance'],
        'bucket': json['bucket'],
    };
}

export function FindNodeCloserPeerToJSON(value?: FindNodeCloserPeer | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'peerId': value.peerId,
        'distance': value.distance,
        'bucket': value.bucket,
    };
}
