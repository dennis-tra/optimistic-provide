import React from "react";
import ReactDOM from "react-dom";
import RootPage from "./pages/RootPage";
import CssBaseline from "@mui/material/CssBaseline";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import HostPage from "./pages/HostPage";
import HostsPage from "./pages/HostsPage";
import { themeOptions } from "./theme";
import { ThemeProvider, createTheme } from "@mui/material";
import TimeAgo from "javascript-time-ago";

import en from "javascript-time-ago/locale/en.json";
import ru from "javascript-time-ago/locale/ru.json";

TimeAgo.addDefaultLocale(en);
TimeAgo.addLocale(ru);

ReactDOM.render(
  <React.StrictMode>
    <CssBaseline />
    <ThemeProvider theme={createTheme(themeOptions)}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<RootPage />} />
          <Route path="/hosts" element={<HostsPage />} />
          <Route path="/hosts/:hostId" element={<HostPage />} />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
