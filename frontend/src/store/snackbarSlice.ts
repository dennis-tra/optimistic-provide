import { Component } from "react";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { SnackbarAction, SnackbarKey, TransitionCloseHandler, TransitionHandler, VariantType } from "notistack";

export interface Notification {
  message: string;
  key: SnackbarKey;
  variant: VariantType;
  action?: SnackbarAction;
  dismissed?: boolean;
  onClose?: TransitionCloseHandler;
}

export interface SnackbarState {
  notifications: Notification[];
}

const snackbarSlice = createSlice({
  name: "snackbar",
  initialState: { notifications: [] } as SnackbarState,
  reducers: {
    addNotification(state, action: PayloadAction<Notification>) {
      state.notifications = [...state.notifications, action.payload];
    },
    dismissNotification(state, action: PayloadAction<SnackbarKey>) {
      state.notifications = state.notifications.map((notification) =>
        notification.key === action.payload ? { ...notification, dismissed: true } : notification
      );
    },
    removeNotification(state, action: PayloadAction<SnackbarKey>) {
      state.notifications = state.notifications.filter((n) => n.key !== action.payload);
    },
  },
});

export const { reducer, actions } = snackbarSlice;
