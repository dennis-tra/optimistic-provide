// Need to use the React-specific entry point to import createApi
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { Host, CreateHostRequest, RoutingTablePeer } from "../api";

// Define a service using a base URL and expected endpoints
export const optprovApi = createApi({
  reducerPath: "optprovApi",
  baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:7000" }),
  tagTypes: ["Host", "RoutingTable"],
  endpoints: (builder) => ({
    createHost: builder.mutation<Host, CreateHostRequest>({
      query: (body) => ({ url: `hosts/`, method: "POST", body }),
      invalidatesTags: [{ type: "Host", id: "LIST" }],
    }),
    deleteHost: builder.mutation<Host, string>({
      query: (hostId) => ({ url: `hosts/${hostId}/`, method: "DELETE" }),
      invalidatesTags: (result, error, arg) => [{ type: "Host", id: arg }],
    }),
    bootstrapHost: builder.mutation<Host, string>({
      query: (hostId) => ({ url: `hosts/${hostId}/bootstrap`, method: "POST" }),
      invalidatesTags: (result, error, arg) => [{ type: "Host", id: arg }],
    }),
    getHosts: builder.query<Host[], void>({
      query: () => `hosts/`,
      providesTags: (result) =>
        result
          ? [...result.map(({ hostId }) => ({ type: "Host", id: hostId } as const)), { type: "Host", id: "LIST" }]
          : [{ type: "Host", id: "LIST" }],
    }),
    getHost: builder.query<Host, string>({
      query: (hostId) => `hosts/${hostId}/`,
      providesTags: (result, error, arg) => [{ type: "Host", id: arg }],
    }),
    getRoutingTablePeers: builder.query<RoutingTablePeer[], string>({
      query: (hostId) => `hosts/${hostId}/routing-tables/current`,
      onCacheEntryAdded: async (arg, { updateCachedData, cacheDataLoaded, cacheEntryRemoved }) => {
        const ws = new WebSocket(`ws://localhost:7000/hosts/${arg}/routing-tables/listen`);
        try {
          await cacheDataLoaded;
          ws.onmessage = (event: MessageEvent) => {
            const data = JSON.parse(event.data);
            updateCachedData((draft) => {
              return data;
            });
          };
        } catch (err) {
          console.error(err);
        }
        await cacheEntryRemoved;
        ws.close();
      },
      providesTags: (result, error, arg) => [{ type: "RoutingTable", id: arg }],
    }),
  }),
});

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const {
  useGetHostQuery,
  useGetHostsQuery,
  useCreateHostMutation,
  useDeleteHostMutation,
  useBootstrapHostMutation,
  useGetRoutingTablePeersQuery,
} = optprovApi;
