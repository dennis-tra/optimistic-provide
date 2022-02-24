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
 * @interface RetrievalRequest
 */
export interface RetrievalRequest {
    /**
     * 
     * @type {string}
     * @memberof RetrievalRequest
     */
    contentId: string;
    /**
     * Number of providers to find until the query stops. 0 indicates that the query will run until it completes.
     * @type {number}
     * @memberof RetrievalRequest
     */
    count: number;
}

export function RetrievalRequestFromJSON(json: any): RetrievalRequest {
    return RetrievalRequestFromJSONTyped(json, false);
}

export function RetrievalRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RetrievalRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'contentId': json['contentId'],
        'count': json['count'],
    };
}

export function RetrievalRequestToJSON(value?: RetrievalRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'contentId': value.contentId,
        'count': value.count,
    };
}

