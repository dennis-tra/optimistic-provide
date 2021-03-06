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

/**
 * 
 * @export
 * @enum {string}
 */
export enum MeasuremetType {
    Rpovide = 'RPOVIDE',
    MonitorProviderRecord = 'MONITOR_PROVIDER_RECORD'
}

export function MeasuremetTypeFromJSON(json: any): MeasuremetType {
    return MeasuremetTypeFromJSONTyped(json, false);
}

export function MeasuremetTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MeasuremetType {
    return json as MeasuremetType;
}

export function MeasuremetTypeToJSON(value?: MeasuremetType | null): any {
    return value as any;
}

