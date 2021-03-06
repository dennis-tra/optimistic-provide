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
    ProvideType,
    ProvideTypeFromJSON,
    ProvideTypeFromJSONTyped,
    ProvideTypeToJSON,
} from './ProvideType';

/**
 * 
 * @export
 * @interface ProvideRequest
 */
export interface ProvideRequest {
    /**
     * 
     * @type {ProvideType}
     * @memberof ProvideRequest
     */
    type: ProvideType;
}

export function ProvideRequestFromJSON(json: any): ProvideRequest {
    return ProvideRequestFromJSONTyped(json, false);
}

export function ProvideRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProvideRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'type': ProvideTypeFromJSON(json['type']),
    };
}

export function ProvideRequestToJSON(value?: ProvideRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'type': ProvideTypeToJSON(value.type),
    };
}

