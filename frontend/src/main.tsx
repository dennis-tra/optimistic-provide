import React from "react";
import ReactDOM from "react-dom";
import CssBaseline from "@mui/material/CssBaseline";
import { SnackbarProvider } from "notistack";
import { themeOptions } from "./theme";
import { ThemeProvider, createTheme } from "@mui/material";
import TimeAgo from "javascript-time-ago";
import { store } from "./store/config";
import { Provider } from "react-redux";
import en from "javascript-time-ago/locale/en.json";
import App from "./app";

TimeAgo.addDefaultLocale(en);

ReactDOM.render(
  <React.StrictMode>
    <CssBaseline />
    <ThemeProvider theme={createTheme(themeOptions)}>
      <Provider store={store}>
        <SnackbarProvider maxSnack={3}>
          <App />
        </SnackbarProvider>
      </Provider>
    </ThemeProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
