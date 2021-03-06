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
 * @interface ProvideMeasurementConfiguration
 */
export interface ProvideMeasurementConfiguration {
    /**
     * 
     * @type {number}
     * @memberof ProvideMeasurementConfiguration
     */
    iterations: number;
    /**
     * Multi-Query parallelism.
     * @type {number}
     * @memberof ProvideMeasurementConfiguration
     */
    concurrency?: number;
    /**
     * 
     * @type {ProvideType}
     * @memberof ProvideMeasurementConfiguration
     */
    provideType: ProvideType;
}

export function ProvideMeasurementConfigurationFromJSON(json: any): ProvideMeasurementConfiguration {
    return ProvideMeasurementConfigurationFromJSONTyped(json, false);
}

export function ProvideMeasurementConfigurationFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProvideMeasurementConfiguration {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'iterations': json['iterations'],
        'concurrency': !exists(json, 'concurrency') ? undefined : json['concurrency'],
        'provideType': ProvideTypeFromJSON(json['provideType']),
    };
}

export function ProvideMeasurementConfigurationToJSON(value?: ProvideMeasurementConfiguration | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'iterations': value.iterations,
        'concurrency': value.concurrency,
        'provideType': ProvideTypeToJSON(value.provideType),
    };
}

