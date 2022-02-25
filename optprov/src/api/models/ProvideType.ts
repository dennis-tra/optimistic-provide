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
export enum ProvideType {
    SingleQuery = 'SINGLE_QUERY',
    MultiQuery = 'MULTI_QUERY'
}

export function ProvideTypeFromJSON(json: any): ProvideType {
    return ProvideTypeFromJSONTyped(json, false);
}

export function ProvideTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProvideType {
    return json as ProvideType;
}

export function ProvideTypeToJSON(value?: ProvideType | null): any {
    return value as any;
}

