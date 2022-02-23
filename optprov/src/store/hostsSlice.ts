import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from "./config";

export interface HostsState {
  value: number;
}

const initialState: HostsState = {
  value: 0,
};

export const hostsSlice = createSlice({
  name: "hosts",
  initialState,
  reducers: {
    increment: (state) => {
      // Redux Toolkit allows us to write "mutating" logic in reducers. It
      // doesn't actually mutate the state because it uses the Immer library,
      // which detects changes to a "draft state" and produces a brand new
      // immutable state based off those changes
      state.value += 1;
    },
    decrement: (state) => {
      state.value -= 1;
    },
    incrementByAmount: (state, action: PayloadAction<number>) => {
      state.value += action.payload;
    },
  },
});

// Action creators are generated for each case reducer function
export const { increment, decrement, incrementByAmount } = hostsSlice.actions;

// Other code such as selectors can use the imported `RootState` type
export const selectCount = (state: RootState) => state.hosts.value;

export default hostsSlice.reducer;
