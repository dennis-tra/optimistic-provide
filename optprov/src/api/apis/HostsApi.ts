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
    CreateHostRequest,
    CreateHostRequestFromJSON,
    CreateHostRequestToJSON,
    Host,
    HostFromJSON,
    HostToJSON,
} from '../models';

export interface BootstrapHostRequest {
    hostId: string;
}

export interface CreateHostOperationRequest {
    createHostRequest?: CreateHostRequest;
}

export interface DeleteHostRequest {
    hostId: string;
}

export interface GetHostRequest {
    hostId: string;
}

/**
 * 
 */
export class HostsApi extends runtime.BaseAPI {

    /**
     * Instructs the given host to connect to the canonical bootstrap peers.
     * Connect to bootstrap nodes.
     */
    async bootstrapHostRaw(requestParameters: BootstrapHostRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Host>> {
        if (requestParameters.hostId === null || requestParameters.hostId === undefined) {
            throw new runtime.RequiredError('hostId','Required parameter requestParameters.hostId was null or undefined when calling bootstrapHost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/{hostId}/bootstrap`.replace(`{${"hostId"}}`, encodeURIComponent(String(requestParameters.hostId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => HostFromJSON(jsonValue));
    }

    /**
     * Instructs the given host to connect to the canonical bootstrap peers.
     * Connect to bootstrap nodes.
     */
    async bootstrapHost(requestParameters: BootstrapHostRequest, initOverrides?: RequestInit): Promise<Host> {
        const response = await this.bootstrapHostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Starts a new libp2p host with the provided parameters.
     * Creates a new libp2p host.
     */
    async createHostRaw(requestParameters: CreateHostOperationRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Host>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/hosts/`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateHostRequestToJSON(requestParameters.createHostRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => HostFromJSON(jsonValue));
    }

    /**
     * Starts a new libp2p host with the provided parameters.
     * Creates a new libp2p host.
     */
    async createHost(requestParameters: CreateHostOperationRequest = {}, initOverrides?: RequestInit): Promise<Host> {
        const response = await this.createHostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Stops a running host and aborts all in-progress provide, refresh, etc. operations.
     * Stops a running host
     */
    async deleteHostRaw(requestParameters: DeleteHostRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.hostId === null || requestParameters.hostId === undefined) {
            throw new runtime.RequiredError('hostId','Required parameter requestParameters.hostId was null or undefined when calling deleteHost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/{hostId}/`.replace(`{${"hostId"}}`, encodeURIComponent(String(requestParameters.hostId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Stops a running host and aborts all in-progress provide, refresh, etc. operations.
     * Stops a running host
     */
    async deleteHost(requestParameters: DeleteHostRequest, initOverrides?: RequestInit): Promise<void> {
        await this.deleteHostRaw(requestParameters, initOverrides);
    }

    /**
     * Returns a single running host.
     * Returns information about a single running host.
     */
    async getHostRaw(requestParameters: GetHostRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Host>> {
        if (requestParameters.hostId === null || requestParameters.hostId === undefined) {
            throw new runtime.RequiredError('hostId','Required parameter requestParameters.hostId was null or undefined when calling getHost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/{hostId}/`.replace(`{${"hostId"}}`, encodeURIComponent(String(requestParameters.hostId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => HostFromJSON(jsonValue));
    }

    /**
     * Returns a single running host.
     * Returns information about a single running host.
     */
    async getHost(requestParameters: GetHostRequest, initOverrides?: RequestInit): Promise<Host> {
        const response = await this.getHostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Returns a list of all running libp2p hosts ordered by their creation date ascending.
     * Lists all running libp2p hosts.
     */
    async getHostsRaw(initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<Host>>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(HostFromJSON));
    }

    /**
     * Returns a list of all running libp2p hosts ordered by their creation date ascending.
     * Lists all running libp2p hosts.
     */
    async getHosts(initOverrides?: RequestInit): Promise<Array<Host>> {
        const response = await this.getHostsRaw(initOverrides);
        return await response.value();
    }

}
