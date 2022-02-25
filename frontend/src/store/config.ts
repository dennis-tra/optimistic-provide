import { configureStore } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { optprovApi } from "./api";
import { reducer as bucketsReducer } from "./bucketsSlice";

export const store = configureStore({
  reducer: {
    buckets: bucketsReducer,
    [optprovApi.reducerPath]: optprovApi.reducer,
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(optprovApi.middleware),
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;

// Use throughout your app instead of plain `useDispatch` and `useSelector`
export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
