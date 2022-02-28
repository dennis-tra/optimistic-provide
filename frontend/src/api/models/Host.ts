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
 * @interface Host
 */
export interface Host {
    /**
     * 
     * @type {string}
     * @memberof Host
     */
    hostId: string;
    /**
     * 
     * @type {string}
     * @memberof Host
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof Host
     */
    startedAt: string | null;
    /**
     * 
     * @type {string}
     * @memberof Host
     */
    createdAt: string;
    /**
     * 
     * @type {string}
     * @memberof Host
     */
    bootstrappedAt: string | null;
    /**
     * 
     * @type {number}
     * @memberof Host
     */
    runningProvidesCount: number;
}

export function HostFromJSON(json: any): Host {
    return HostFromJSONTyped(json, false);
}

export function HostFromJSONTyped(json: any, ignoreDiscriminator: boolean): Host {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hostId': json['hostId'],
        'name': json['name'],
        'startedAt': json['startedAt'],
        'createdAt': json['createdAt'],
        'bootstrappedAt': json['bootstrappedAt'],
        'runningProvidesCount': json['runningProvidesCount'],
    };
}

export function HostToJSON(value?: Host | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hostId': value.hostId,
        'name': value.name,
        'startedAt': value.startedAt,
        'createdAt': value.createdAt,
        'bootstrappedAt': value.bootstrappedAt,
        'runningProvidesCount': value.runningProvidesCount,
    };
}

