import React from "react";
import ReactDOM from "react-dom";
import HostsPage from "./pages/HostsPage";
import CssBaseline from "@mui/material/CssBaseline";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import HostPage from "./pages/HostPage";
import { themeOptions } from "./theme";
import { ThemeProvider, createTheme } from "@mui/material";
import TimeAgo from "javascript-time-ago";
import { store } from "./store/config";
import { Provider } from "react-redux";
import en from "javascript-time-ago/locale/en.json";
import ru from "javascript-time-ago/locale/ru.json";
import RoutingTablePage from "./pages/RoutingTablePage";
import ProvidePage from "./pages/ProvidePage";
import ProvideDetailsPage from "./pages/ProvideDetailsPage";
import RetrievalsPage from "./pages/RetrievalsPage";
import NotFoundPage from "./pages/404";

TimeAgo.addDefaultLocale(en);
TimeAgo.addLocale(ru);

ReactDOM.render(
  <React.StrictMode>
    <CssBaseline />
    <ThemeProvider theme={createTheme(themeOptions)}>
      <Provider store={store}>
        <BrowserRouter>
          <Routes>
            <Route path="/hosts" element={<HostsPage />} />
            <Route path="/hosts/:hostId" element={<HostPage />} />
            <Route path="/hosts/:hostId/routing-tables" element={<RoutingTablePage />} />
            <Route path="/hosts/:hostId/provides" element={<ProvidePage />} />
            <Route path="/hosts/:hostId/provides/:provideId" element={<ProvideDetailsPage />} />
            <Route path="/hosts/:hostId/retrievals" element={<RetrievalsPage />} />
            <Route path="*" element={<NotFoundPage />} />
          </Routes>
        </BrowserRouter>
      </Provider>
    </ThemeProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
