import { configureStore, MiddlewareAPI, isRejectedWithValue, Middleware } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { optprovApi } from "./api";
import { reducer as bucketsReducer } from "./bucketsSlice";
import { reducer as snackbarReducer, actions as snackbarActions } from "./snackbarSlice";

export const rtkQueryErrorLogger: Middleware = (api: MiddlewareAPI) => (next) => (action) => {
  if (isRejectedWithValue(action)) {
    console.warn("Action rejected:", action);
    api.dispatch(
      snackbarActions.addNotification({
        key: new Date().getTime() + Math.random(),
        variant: "error",
        message: `Payload: ${JSON.stringify(
          action.payload.data?.message || action.payload.error
        )}, Error: ${JSON.stringify(action.error)}`,
      })
    );
  }

  return next(action);
};
export const store = configureStore({
  reducer: {
    buckets: bucketsReducer,
    snackbar: snackbarReducer,
    [optprovApi.reducerPath]: optprovApi.reducer,
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(optprovApi.middleware, rtkQueryErrorLogger),
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;

// Use throughout your app instead of plain `useDispatch` and `useSelector`
export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
