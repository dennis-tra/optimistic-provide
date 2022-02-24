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
    AddProvider,
    AddProviderFromJSON,
    AddProviderFromJSONTyped,
    AddProviderToJSON,
} from './AddProvider';
import {
    Connection,
    ConnectionFromJSON,
    ConnectionFromJSONTyped,
    ConnectionToJSON,
} from './Connection';
import {
    Dial,
    DialFromJSON,
    DialFromJSONTyped,
    DialToJSON,
} from './Dial';
import {
    FindNode,
    FindNodeFromJSON,
    FindNodeFromJSONTyped,
    FindNodeToJSON,
} from './FindNode';

/**
 * 
 * @export
 * @interface ProvideDetailsAllOf
 */
export interface ProvideDetailsAllOf {
    /**
     * 
     * @type {Array<Connection>}
     * @memberof ProvideDetailsAllOf
     */
    connections: Array<Connection>;
    /**
     * 
     * @type {Array<Dial>}
     * @memberof ProvideDetailsAllOf
     */
    dials: Array<Dial>;
    /**
     * 
     * @type {Array<FindNode>}
     * @memberof ProvideDetailsAllOf
     */
    findNodes: Array<FindNode>;
    /**
     * 
     * @type {Array<AddProvider>}
     * @memberof ProvideDetailsAllOf
     */
    addProviders: Array<AddProvider>;
}

export function ProvideDetailsAllOfFromJSON(json: any): ProvideDetailsAllOf {
    return ProvideDetailsAllOfFromJSONTyped(json, false);
}

export function ProvideDetailsAllOfFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProvideDetailsAllOf {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'connections': ((json['connections'] as Array<any>).map(ConnectionFromJSON)),
        'dials': ((json['dials'] as Array<any>).map(DialFromJSON)),
        'findNodes': ((json['findNodes'] as Array<any>).map(FindNodeFromJSON)),
        'addProviders': ((json['addProviders'] as Array<any>).map(AddProviderFromJSON)),
    };
}

export function ProvideDetailsAllOfToJSON(value?: ProvideDetailsAllOf | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'connections': ((value.connections as Array<any>).map(ConnectionToJSON)),
        'dials': ((value.dials as Array<any>).map(DialToJSON)),
        'findNodes': ((value.findNodes as Array<any>).map(FindNodeToJSON)),
        'addProviders': ((value.addProviders as Array<any>).map(AddProviderToJSON)),
    };
}
