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
import { useDeleteHostMutation, useBootstrapHostMutation } from "../store/api";
import { Host } from "../api/models/Host";

interface HostCardProps {
  host: Host;
}

const HostCard: React.FC<HostCardProps> = ({ host }) => {
  const [deleteHost] = useDeleteHostMutation();
  const [bootstrapHost] = useBootstrapHostMutation();

  const [snackbarOpen, setSnackbarOpen] = useState(false);
  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSnackbarOpen(false);
  };

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
            onClick={() => {
              navigator.clipboard.writeText(host.hostId);
              setSnackbarOpen(true);
            }}
            sx={{ cursor: "pointer" }}
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
          {host.bootstrappedAt ? (
            <Tooltip title={new Date(host.createdAt).toLocaleString()}>
              <Chip label="Bootstrapped" color="success" variant="filled" />
            </Tooltip>
          ) : (
            <Tooltip title={"Click to bootstrap"}>
              <Chip
                label="Not Bootstrapped"
                color="warning"
                variant="outlined"
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
          <Button
            color="error"
            variant="outlined"
            onClick={async () => {
              await deleteHost(host.hostId);
            }}
          >
            Delete
          </Button>
        </Stack>
      </Paper>
      <Snackbar open={snackbarOpen} autoHideDuration={3000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success" sx={{ width: "100%" }}>
          Peer ID copied to clipboard!
        </Alert>
      </Snackbar>
    </Grid>
  );
};

export default HostCard;
