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
import {
    ProvidePeerInfo,
    ProvidePeerInfoFromJSON,
    ProvidePeerInfoFromJSONTyped,
    ProvidePeerInfoToJSON,
} from './ProvidePeerInfo';

/**
 * 
 * @export
 * @interface ProvideGraphAllOf
 */
export interface ProvideGraphAllOf {
    /**
     * All peers + information in the order they should be plotted.
     * @type {Array<ProvidePeerInfo>}
     * @memberof ProvideGraphAllOf
     */
    peers: Array<ProvidePeerInfo>;
    /**
     * 
     * @type {Array<Dial>}
     * @memberof ProvideGraphAllOf
     */
    dials: Array<Dial>;
    /**
     * 
     * @type {Array<Connection>}
     * @memberof ProvideGraphAllOf
     */
    connections: Array<Connection>;
    /**
     * 
     * @type {Array<FindNode>}
     * @memberof ProvideGraphAllOf
     */
    findNodes: Array<FindNode>;
    /**
     * 
     * @type {Array<AddProvider>}
     * @memberof ProvideGraphAllOf
     */
    addProviders: Array<AddProvider>;
}

export function ProvideGraphAllOfFromJSON(json: any): ProvideGraphAllOf {
    return ProvideGraphAllOfFromJSONTyped(json, false);
}

export function ProvideGraphAllOfFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProvideGraphAllOf {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'peers': ((json['peers'] as Array<any>).map(ProvidePeerInfoFromJSON)),
        'dials': ((json['dials'] as Array<any>).map(DialFromJSON)),
        'connections': ((json['connections'] as Array<any>).map(ConnectionFromJSON)),
        'findNodes': ((json['findNodes'] as Array<any>).map(FindNodeFromJSON)),
        'addProviders': ((json['addProviders'] as Array<any>).map(AddProviderFromJSON)),
    };
}

export function ProvideGraphAllOfToJSON(value?: ProvideGraphAllOf | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'peers': ((value.peers as Array<any>).map(ProvidePeerInfoToJSON)),
        'dials': ((value.dials as Array<any>).map(DialToJSON)),
        'connections': ((value.connections as Array<any>).map(ConnectionToJSON)),
        'findNodes': ((value.findNodes as Array<any>).map(FindNodeToJSON)),
        'addProviders': ((value.addProviders as Array<any>).map(AddProviderToJSON)),
    };
}

