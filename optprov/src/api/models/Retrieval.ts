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
 * @interface Retrieval
 */
export interface Retrieval {
    /**
     * 
     * @type {number}
     * @memberof Retrieval
     */
    retrievalId: number;
    /**
     * 
     * @type {string}
     * @memberof Retrieval
     */
    hostId: string;
    /**
     * 
     * @type {string}
     * @memberof Retrieval
     */
    contentId: string;
    /**
     * 
     * @type {number}
     * @memberof Retrieval
     */
    initialRoutingTableId: number;
    /**
     * 
     * @type {number}
     * @memberof Retrieval
     */
    finalRoutingTableId: number | null;
    /**
     * 
     * @type {string}
     * @memberof Retrieval
     */
    startedAt: string;
    /**
     * 
     * @type {string}
     * @memberof Retrieval
     */
    endedAt: string | null;
    /**
     * 
     * @type {string}
     * @memberof Retrieval
     */
    error: string | null;
}

export function RetrievalFromJSON(json: any): Retrieval {
    return RetrievalFromJSONTyped(json, false);
}

export function RetrievalFromJSONTyped(json: any, ignoreDiscriminator: boolean): Retrieval {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'retrievalId': json['retrievalId'],
        'hostId': json['hostId'],
        'contentId': json['contentId'],
        'initialRoutingTableId': json['initialRoutingTableId'],
        'finalRoutingTableId': json['finalRoutingTableId'],
        'startedAt': json['startedAt'],
        'endedAt': json['endedAt'],
        'error': json['error'],
    };
}

export function RetrievalToJSON(value?: Retrieval | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'retrievalId': value.retrievalId,
        'hostId': value.hostId,
        'contentId': value.contentId,
        'initialRoutingTableId': value.initialRoutingTableId,
        'finalRoutingTableId': value.finalRoutingTableId,
        'startedAt': value.startedAt,
        'endedAt': value.endedAt,
        'error': value.error,
    };
}
