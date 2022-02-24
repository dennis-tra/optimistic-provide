// Need to use the React-specific entry point to import createApi
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RoutingTablePeer } from "../api";
import { RootState } from "./config";

interface BucketsState {
  buckets: number[];
  peers: {
    [bucket: number]: RoutingTablePeer[];
  };
}
interface BucketsStateByHost {
  [hostId: string]: BucketsState;
}

const buckets = [...Array(256).keys()];
const initialState = (): BucketsState => ({
  buckets: buckets,
  peers: buckets.reduce((prev, curr) => ({ ...prev, [curr]: [] }), {}),
});

const bucketsSlice = createSlice({
  name: "buckets",
  initialState: {} as BucketsStateByHost,
  reducers: {
    addPeer(state, action: PayloadAction<{ hostId: string; peer: RoutingTablePeer }>) {
      state[action.payload.hostId].peers[action.payload.peer.bucket].push(action.payload.peer);
    },
    removePeer(state, action: PayloadAction<{ hostId: string; peerId: string }>) {
      for (const bucket of buckets) {
        for (let i = 0; i < state[action.payload.hostId].peers[bucket].length; i++) {
          const peer = state[action.payload.hostId].peers[bucket][i];
          if (peer.peerId === action.payload.peerId) {
            state[action.payload.hostId].peers[bucket].splice(i, 1);
            return;
          }
        }
      }
    },
    replace(state, action: PayloadAction<{ hostId: string; peers: RoutingTablePeer[] }>) {
      state[action.payload.hostId] = initialState();
      for (const peer of action.payload.peers) {
        state[action.payload.hostId].peers[peer.bucket].push(peer);
      }
      return state;
    },
    reset(state, action: PayloadAction<string>) {
      delete state[action.payload];
    },
  },
});

export const selectBucketPeers =
  (hostId: string, bucket: number) =>
  (state: RootState): RoutingTablePeer[] => {
    if (!state.buckets[hostId] || bucket === undefined || bucket === null || !state.buckets[hostId].peers[bucket]) {
      return [];
    }
    return state.buckets[hostId].peers[bucket];
  };

export interface HistogramData {
  bucket: number;
  level: number;
}
export const selectHistogramData =
  (hostId: string) =>
  (state: RootState): HistogramData[] => {
    if (!state.buckets[hostId]) {
      return [...Array(16).keys()].map((bucket) => ({
        bucket: bucket,
        level: 0,
      }));
    }
    const data = state.buckets[hostId].buckets.map((bucket) => ({
      bucket: bucket,
      level: state.buckets[hostId].peers[bucket].length,
    }));
    for (let i = data.length - 1; i >= 16; i--) {
      if (data[i].level !== 0) {
        return data.slice(0, i + 1);
      }
    }
    return data.slice(0, 16);
  };

export const { reducer, actions } = bucketsSlice;
