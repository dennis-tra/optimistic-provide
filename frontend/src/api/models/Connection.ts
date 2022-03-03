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
 * @interface Connection
 */
export interface Connection {
    /**
     * 
     * @type {number}
     * @memberof Connection
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof Connection
     */
    remoteId: string;
    /**
     * 
     * @type {string}
     * @memberof Connection
     */
    multiAddress: string;
    /**
     * 
     * @type {string}
     * @memberof Connection
     */
    startedAt: string;
    /**
     * 
     * @type {string}
     * @memberof Connection
     */
    endedAt: string;
    /**
     * 
     * @type {number}
     * @memberof Connection
     */
    durationInS: number;
}

export function ConnectionFromJSON(json: any): Connection {
    return ConnectionFromJSONTyped(json, false);
}

export function ConnectionFromJSONTyped(json: any, ignoreDiscriminator: boolean): Connection {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'remoteId': json['remoteId'],
        'multiAddress': json['multiAddress'],
        'startedAt': json['startedAt'],
        'endedAt': json['endedAt'],
        'durationInS': json['durationInS'],
    };
}

export function ConnectionToJSON(value?: Connection | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'remoteId': value.remoteId,
        'multiAddress': value.multiAddress,
        'startedAt': value.startedAt,
        'endedAt': value.endedAt,
        'durationInS': value.durationInS,
    };
}
