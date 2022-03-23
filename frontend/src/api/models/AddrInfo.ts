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
 * @interface AddrInfo
 */
export interface AddrInfo {
    /**
     * 
     * @type {string}
     * @memberof AddrInfo
     */
    peerID: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof AddrInfo
     */
    multiAddresses: Array<string>;
}

export function AddrInfoFromJSON(json: any): AddrInfo {
    return AddrInfoFromJSONTyped(json, false);
}

export function AddrInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddrInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'peerID': json['peerID'],
        'multiAddresses': json['multiAddresses'],
    };
}

export function AddrInfoToJSON(value?: AddrInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'peerID': value.peerID,
        'multiAddresses': value.multiAddresses,
    };
}

