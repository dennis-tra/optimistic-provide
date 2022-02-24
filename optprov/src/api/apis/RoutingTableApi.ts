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


import * as runtime from '../runtime';
import {
    ErrorResponse,
    ErrorResponseFromJSON,
    ErrorResponseToJSON,
    RoutingTable,
    RoutingTableFromJSON,
    RoutingTableToJSON,
    RoutingTableDetails,
    RoutingTableDetailsFromJSON,
    RoutingTableDetailsToJSON,
    RoutingTablePeer,
    RoutingTablePeerFromJSON,
    RoutingTablePeerToJSON,
    RoutingTableUpdate,
    RoutingTableUpdateFromJSON,
    RoutingTableUpdateToJSON,
} from '../models';

export interface GetCurrentRoutingTableRequest {
    hostId: string;
}

export interface GetRoutingTableRequest {
    hostId: string;
    routingTableId: number;
}

export interface GetRoutingTablesRequest {
    hostId: string;
}

export interface ListenRoutingTableRequest {
    hostId: string;
}

export interface RoutingTableCreateRequest {
    hostId: string;
}

/**
 * 
 */
export class RoutingTableApi extends runtime.BaseAPI {

    /**
     * Returns the current routing table with its entries for the given peer.
     * Returns the current routing table.
     */
    async getCurrentRoutingTableRaw(requestParameters: GetCurrentRoutingTableRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<RoutingTablePeer>>> {
        if (requestParameters.hostId === null || requestParameters.hostId === undefined) {
            throw new runtime.RequiredError('hostId','Required parameter requestParameters.hostId was null or undefined when calling getCurrentRoutingTable.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/{hostId}/routing-table`.replace(`{${"hostId"}}`, encodeURIComponent(String(requestParameters.hostId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(RoutingTablePeerFromJSON));
    }

    /**
     * Returns the current routing table with its entries for the given peer.
     * Returns the current routing table.
     */
    async getCurrentRoutingTable(requestParameters: GetCurrentRoutingTableRequest, initOverrides?: RequestInit): Promise<Array<RoutingTablePeer>> {
        const response = await this.getCurrentRoutingTableRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Returns a single routing table with its entries
     * Returns a single routing table.
     */
    async getRoutingTableRaw(requestParameters: GetRoutingTableRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<RoutingTableDetails>>> {
        if (requestParameters.hostId === null || requestParameters.hostId === undefined) {
            throw new runtime.RequiredError('hostId','Required parameter requestParameters.hostId was null or undefined when calling getRoutingTable.');
        }

        if (requestParameters.routingTableId === null || requestParameters.routingTableId === undefined) {
            throw new runtime.RequiredError('routingTableId','Required parameter requestParameters.routingTableId was null or undefined when calling getRoutingTable.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/{hostId}/routing-tables/{routingTableId}`.replace(`{${"hostId"}}`, encodeURIComponent(String(requestParameters.hostId))).replace(`{${"routingTableId"}}`, encodeURIComponent(String(requestParameters.routingTableId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(RoutingTableDetailsFromJSON));
    }

    /**
     * Returns a single routing table with its entries
     * Returns a single routing table.
     */
    async getRoutingTable(requestParameters: GetRoutingTableRequest, initOverrides?: RequestInit): Promise<Array<RoutingTableDetails>> {
        const response = await this.getRoutingTableRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Lists all saved routing tables for a specific host.
     * Lists all saved routing tables.
     */
    async getRoutingTablesRaw(requestParameters: GetRoutingTablesRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<RoutingTable>>> {
        if (requestParameters.hostId === null || requestParameters.hostId === undefined) {
            throw new runtime.RequiredError('hostId','Required parameter requestParameters.hostId was null or undefined when calling getRoutingTables.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/{hostId}/routing-tables`.replace(`{${"hostId"}}`, encodeURIComponent(String(requestParameters.hostId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(RoutingTableFromJSON));
    }

    /**
     * Lists all saved routing tables for a specific host.
     * Lists all saved routing tables.
     */
    async getRoutingTables(requestParameters: GetRoutingTablesRequest, initOverrides?: RequestInit): Promise<Array<RoutingTable>> {
        const response = await this.getRoutingTablesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Returns routing table updates as they occur for the given host. Prepend is a full update
     * Subscribe to real time updates of the routing table.
     */
    async listenRoutingTableRaw(requestParameters: ListenRoutingTableRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<RoutingTableUpdate>> {
        if (requestParameters.hostId === null || requestParameters.hostId === undefined) {
            throw new runtime.RequiredError('hostId','Required parameter requestParameters.hostId was null or undefined when calling listenRoutingTable.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/{hostId}/routing-tables/listen`.replace(`{${"hostId"}}`, encodeURIComponent(String(requestParameters.hostId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RoutingTableUpdateFromJSON(jsonValue));
    }

    /**
     * Returns routing table updates as they occur for the given host. Prepend is a full update
     * Subscribe to real time updates of the routing table.
     */
    async listenRoutingTable(requestParameters: ListenRoutingTableRequest, initOverrides?: RequestInit): Promise<RoutingTableUpdate> {
        const response = await this.listenRoutingTableRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Saves a current routing table snapshot of the given host.
     * Saves a current routing table snapshot of the given host.
     */
    async routingTableCreateRaw(requestParameters: RoutingTableCreateRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<RoutingTable>> {
        if (requestParameters.hostId === null || requestParameters.hostId === undefined) {
            throw new runtime.RequiredError('hostId','Required parameter requestParameters.hostId was null or undefined when calling routingTableCreate.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/{hostId}/routing-tables`.replace(`{${"hostId"}}`, encodeURIComponent(String(requestParameters.hostId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RoutingTableFromJSON(jsonValue));
    }

    /**
     * Saves a current routing table snapshot of the given host.
     * Saves a current routing table snapshot of the given host.
     */
    async routingTableCreate(requestParameters: RoutingTableCreateRequest, initOverrides?: RequestInit): Promise<RoutingTable> {
        const response = await this.routingTableCreateRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
