import React from "react";
import HostsPage from "./pages/HostsPage";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import HostPage from "./pages/HostPage";
import RoutingTablePage from "./pages/RoutingTablePage";
import ProvidePage from "./pages/ProvidePage";
import ProvideDetailsPage from "./pages/ProvideDetailsPage";
import RetrievalsPage from "./pages/RetrievalsPage";
import NotFoundPage from "./pages/404";
import useNotifier from "./store/snackbarHook";
import DialsPage from "./pages/DialsPage";

function App() {
  useNotifier();

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/hosts" element={<HostsPage />} />
        <Route path="/hosts/:hostId" element={<HostPage />} />
        <Route path="/hosts/:hostId/routing-tables" element={<RoutingTablePage />} />
        <Route path="/hosts/:hostId/provides" element={<ProvidePage />} />
        <Route path="/hosts/:hostId/provides/:provideId" element={<ProvideDetailsPage />} />
        <Route path="/hosts/:hostId/provides/:provideId/dials" element={<DialsPage />} />
        <Route path="/hosts/:hostId/retrievals" element={<RetrievalsPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
