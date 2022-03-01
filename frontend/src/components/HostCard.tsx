import { useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";
import Tooltip from "@mui/material/Tooltip";
import Snackbar from "@mui/material/Snackbar";
import Alert from "@mui/material/Alert";
import Box from "@mui/material/Box";
import Chip from "@mui/material/Chip";
import ReactTimeAgo from "react-time-ago";
import {
  useArchiveHostMutation,
  useBootstrapHostMutation,
  useStartHostMutation,
  useStopHostMutation,
} from "../store/api";
import { Host } from "../api/models/Host";
import { actions as snackbarActions } from "../store/snackbarSlice";
import { useAppDispatch } from "../store/config";

interface HostCardProps {
  host: Host;
}

const HostCard: React.FC<HostCardProps> = ({ host }) => {
  const dispatch = useAppDispatch();
  const [archiveHost] = useArchiveHostMutation();
  const [bootstrapHost, { isLoading: isBootstrappingHost }] = useBootstrapHostMutation();
  const [startHost, { isLoading: isStartingHost }] = useStartHostMutation();

  return (
    <Grid item xs={12} md={4} lg={4}>
      <Paper
        sx={{
          p: 2,
          display: "flex",
          flexDirection: "column",
          height: 240,
        }}
      >
        <Tooltip title={`Click to copy: ${host.hostId}`}>
          <Typography
            component="p"
            variant="h6"
            noWrap
            onClick={async () => {
              await navigator.clipboard.writeText(host.hostId);
              dispatch(
                snackbarActions.addNotification({
                  key: new Date().getTime() + Math.random(),
                  variant: "success",
                  message: "Peer ID copied to clipboard!",
                })
              );
            }}
            sx={{ cursor: "pointer" }}
            gutterBottom
          >
            {host.hostId}
          </Typography>
        </Tooltip>
        <Typography component="h2" variant="h4" color="primary">
          {host.name}
        </Typography>
        <Tooltip title={new Date(host.createdAt).toLocaleString()}>
          <Typography color="text.secondary" noWrap>
            Created <ReactTimeAgo date={new Date(host.createdAt)} />
          </Typography>
        </Tooltip>
        <Stack direction="row" spacing={1} mt={1}>
          {host.startedAt && (
            <Tooltip title={new Date(host.startedAt).toLocaleString()}>
              <Chip label="Running" color="success" variant="filled" />
            </Tooltip>
          )}
          {!host.startedAt && (
            <Tooltip title={"Click to start"}>
              <Chip
                label="Not Running"
                color="warning"
                variant="outlined"
                disabled={isStartingHost}
                onClick={async () => {
                  await startHost(host.hostId);
                }}
              />
            </Tooltip>
          )}
          {host.bootstrappedAt && (
            <Tooltip title={new Date(host.createdAt).toLocaleString()}>
              <Chip label="Bootstrapped" color="success" variant="filled" />
            </Tooltip>
          )}
          {host.startedAt && !host.bootstrappedAt && (
            <Tooltip title={"Click to bootstrap"}>
              <Chip
                label="Not Bootstrapped"
                color="warning"
                variant="outlined"
                disabled={isBootstrappingHost}
                onClick={async () => {
                  await bootstrapHost(host.hostId);
                }}
              />
            </Tooltip>
          )}
        </Stack>
        <Box sx={{ flex: 1 }}></Box>
        <Stack direction="row" spacing={1} justifyContent="space-between" alignItems="center">
          <Button color="primary" variant="contained" component={RouterLink} to={`/hosts/${host.hostId}`}>
            Details
          </Button>
          <Button color="error" variant="outlined" onClick={() => archiveHost(host.hostId)}>
            Archive
          </Button>
        </Stack>
      </Paper>
    </Grid>
  );
};

export default HostCard;
