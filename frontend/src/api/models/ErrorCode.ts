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
export enum ErrorCode {
    Internal = 'INTERNAL',
    HostNotFound = 'HOST_NOT_FOUND',
    MalformedRequest = 'MALFORMED_REQUEST',
    PeerNotFound = 'PEER_NOT_FOUND',
    RoutingTableNotFound = 'ROUTING_TABLE_NOT_FOUND',
    MalformedPeerId = 'MALFORMED_PEER_ID',
    SavingRoutingTable = 'SAVING_ROUTING_TABLE',
    HostStopped = 'HOST_STOPPED'
}

export function ErrorCodeFromJSON(json: any): ErrorCode {
    return ErrorCodeFromJSONTyped(json, false);
}

export function ErrorCodeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ErrorCode {
    return json as ErrorCode;
}

export function ErrorCodeToJSON(value?: ErrorCode | null): any {
    return value as any;
}

