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
    Provide,
    ProvideFromJSON,
    ProvideToJSON,
} from '../models';

export interface StartProvideRequest {
    hostId: string;
}

/**
 * 
 */
export class ProvidesApi extends runtime.BaseAPI {

    /**
     * Instructs the given host to generate random content and announce its CID to the network.
     * Starts providing random content.
     */
    async startProvideRaw(requestParameters: StartProvideRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Provide>> {
        if (requestParameters.hostId === null || requestParameters.hostId === undefined) {
            throw new runtime.RequiredError('hostId','Required parameter requestParameters.hostId was null or undefined when calling startProvide.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/hosts/{hostId}/provides/`.replace(`{${"hostId"}}`, encodeURIComponent(String(requestParameters.hostId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ProvideFromJSON(jsonValue));
    }

    /**
     * Instructs the given host to generate random content and announce its CID to the network.
     * Starts providing random content.
     */
    async startProvide(requestParameters: StartProvideRequest, initOverrides?: RequestInit): Promise<Provide> {
        const response = await this.startProvideRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
