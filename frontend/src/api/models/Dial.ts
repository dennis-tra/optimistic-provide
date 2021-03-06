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
 * @interface Dial
 */
export interface Dial {
    /**
     * 
     * @type {number}
     * @memberof Dial
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof Dial
     */
    remoteId: string;
    /**
     * 
     * @type {string}
     * @memberof Dial
     */
    transport: string;
    /**
     * 
     * @type {string}
     * @memberof Dial
     */
    multiAddress: string;
    /**
     * 
     * @type {string}
     * @memberof Dial
     */
    startedAt: string;
    /**
     * 
     * @type {string}
     * @memberof Dial
     */
    endedAt: string;
    /**
     * 
     * @type {number}
     * @memberof Dial
     */
    durationInS: number;
    /**
     * 
     * @type {string}
     * @memberof Dial
     */
    error: string | null;
}

export function DialFromJSON(json: any): Dial {
    return DialFromJSONTyped(json, false);
}

export function DialFromJSONTyped(json: any, ignoreDiscriminator: boolean): Dial {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'remoteId': json['remoteId'],
        'transport': json['transport'],
        'multiAddress': json['multiAddress'],
        'startedAt': json['startedAt'],
        'endedAt': json['endedAt'],
        'durationInS': json['durationInS'],
        'error': json['error'],
    };
}

export function DialToJSON(value?: Dial | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'remoteId': value.remoteId,
        'transport': value.transport,
        'multiAddress': value.multiAddress,
        'startedAt': value.startedAt,
        'endedAt': value.endedAt,
        'durationInS': value.durationInS,
        'error': value.error,
    };
}

