// Need to use the React-specific entry point to import createApi
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { Host, CreateHostRequest, RoutingTablePeer, RoutingTableUpdate } from "../api";
import { actions as bucketsActions } from "./bucketsSlice";

// Define a service using a base URL and expected endpoints
export const optprovApi = createApi({
  reducerPath: "optprovApi",
  baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:7000" }),
  tagTypes: ["Host", "RoutingTable"],
  endpoints: (builder) => ({
    saveRoutingTable: builder.mutation<Host, string>({
      query: (hostId) => ({ url: `hosts/${hostId}/routing-tables/`, method: "POST" }),
      invalidatesTags: [{ type: "RoutingTable", id: "LIST" }],
    }),
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
    getCurrentRoutingTablePeers: builder.query<RoutingTablePeer[], string>({
      query: (hostId) => `hosts/${hostId}/routing-table`,
      onCacheEntryAdded: async (hostId, { cacheDataLoaded, cacheEntryRemoved, dispatch }) => {
        const fullUpdate = await cacheDataLoaded;
        dispatch(bucketsActions.replace({ hostId, peers: fullUpdate.data }));
      },
      providesTags: (result, error, arg) => [{ type: "RoutingTable", id: arg }],
    }),
    listenRoutingTable: builder.query<RoutingTablePeer[], string>({
      query: (hostId) => `hosts/${hostId}/routing-table`,
      onCacheEntryAdded: async (hostId, { cacheDataLoaded, cacheEntryRemoved, dispatch }) => {
        const ws = new WebSocket(`ws://localhost:7000/hosts/${hostId}/routing-tables/listen`);
        try {
          const fullUpdate = await cacheDataLoaded;
          dispatch(bucketsActions.replace({ hostId, peers: fullUpdate.data }));
          ws.onmessage = (event: MessageEvent) => {
            const data = JSON.parse(event.data) as RoutingTableUpdate;
            switch (data.type) {
              case "PEER_ADDED":
                const peer = data.update as RoutingTablePeer;
                dispatch(bucketsActions.addPeer({ hostId, peer }));
                break;
              case "PEER_REMOVED":
                const peerId = data.update as string;
                dispatch(bucketsActions.removePeer({ hostId, peerId }));
                break;
              case "FULL":
                const peers = data.update as RoutingTablePeer[];
                dispatch(bucketsActions.replace({ hostId, peers }));
                break;
            }
          };
        } catch (err) {
          console.error(err);
        } finally {
          await cacheEntryRemoved;
          ws.close();
        }
      },
      providesTags: (result, error, arg) => [{ type: "RoutingTable", id: arg }],
    }),
  }),
});

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const {
  useSaveRoutingTableMutation,
  useGetHostQuery,
  useGetHostsQuery,
  useCreateHostMutation,
  useDeleteHostMutation,
  useBootstrapHostMutation,
  useListenRoutingTableQuery,
  useGetCurrentRoutingTablePeersQuery,
  useLazyGetCurrentRoutingTablePeersQuery,
} = optprovApi;
