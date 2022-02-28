import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import { ResponsiveContainer, BarChart, Bar, CartesianGrid, XAxis, YAxis, Tooltip } from "recharts";
import { Host } from "../../api/models/Host";
import {
  useListenRoutingTableQuery,
  useSaveRoutingTableMutation,
  useLazyGetCurrentRoutingTablePeersQuery,
} from "../../store/api";
import { Button, CircularProgress, Stack, Alert, Snackbar } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { useAppSelector } from "../../store/config";
import { selectHistogramData } from "../../store/bucketsSlice";
import { useState } from "react";

interface HostDetailsRoutingTableCardProps {
  host: Host;
}

const HostDetailsRoutingTableCard: React.FC<HostDetailsRoutingTableCardProps> = ({ host }) => {
  const navigate = useNavigate();
  const [snackbarOpen, setSnackbarOpen] = useState<string | null>(null);
  const [saveRoutingTable, { isLoading: isSavingRoutingTable }] = useSaveRoutingTableMutation();
  const [getCurrentRoutingTablePeers, { isLoading: isLoadingCurrentRoutingTablePeers }] =
    useLazyGetCurrentRoutingTablePeersQuery();
  const bucketData = useAppSelector(selectHistogramData(host.hostId));
  const { isLoading } = useListenRoutingTableQuery(host.hostId);

  if (isLoading) {
    return <CircularProgress />;
  }

  const handleClick = (data: any, index: number) => {
    navigate(`/hosts/${host.hostId}/routing-tables?bucket=${data.bucket}`);
  };

  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSnackbarOpen(null);
  };

  return (
    <Paper
      sx={{
        p: 2,
        display: "flex",
        flexDirection: "column",
        height: 390,
      }}
    >
      <Typography component="h2" variant="h6" color="primary" gutterBottom>
        Routing Table
      </Typography>
      <ResponsiveContainer>
        <BarChart width={150} height={40} data={bucketData}>
          <CartesianGrid strokeDasharray="3 3" />
          <Tooltip />
          <XAxis dataKey="bucket" />
          <YAxis />
          <Bar dataKey="level" fill="#82ca9d" onClick={handleClick} />
        </BarChart>
      </ResponsiveContainer>
      <Stack direction="row" spacing={2}>
        <Button
          variant="outlined"
          onClick={() => saveRoutingTable(host.hostId).then(() => setSnackbarOpen("Saved routing table snapshot"))}
          disabled={isSavingRoutingTable}
        >
          Save Snapshot
        </Button>
        <Button
          variant="outlined"
          onClick={() =>
            getCurrentRoutingTablePeers(host.hostId).then(() => setSnackbarOpen("Refreshed routing table"))
          }
          disabled={isLoadingCurrentRoutingTablePeers}
        >
          Refetch
        </Button>
      </Stack>
      <Snackbar open={!!snackbarOpen} autoHideDuration={3000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success" sx={{ width: "100%" }}>
          {snackbarOpen}
        </Alert>
      </Snackbar>
    </Paper>
  );
};

export default HostDetailsRoutingTableCard;
