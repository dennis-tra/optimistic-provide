import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import { ResponsiveContainer, BarChart, Bar, CartesianGrid, XAxis, YAxis, Tooltip } from "recharts";
import { Host } from "../../api/models/Host";
import {
  useListenRoutingTableQuery,
  useSaveRoutingTableMutation,
  useLazyGetCurrentRoutingTablePeersQuery,
} from "../../store/api";
import { Button, CircularProgress, Stack } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "../../store/config";
import { selectHistogramData } from "../../store/bucketsSlice";
import { actions as snackbarActions } from "../../store/snackbarSlice";

interface HostDetailsRoutingTableCardProps {
  host: Host;
}

const HostDetailsRoutingTableCard: React.FC<HostDetailsRoutingTableCardProps> = ({ host }) => {
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const [saveRoutingTable, { isLoading: isSavingRoutingTable }] = useSaveRoutingTableMutation();
  const [getCurrentRoutingTablePeers, { isLoading: isLoadingCurrentRoutingTablePeers }] =
    useLazyGetCurrentRoutingTablePeersQuery();
  const bucketData = useAppSelector(selectHistogramData(host.hostId));
  const {
    isLoading,
    isError: isListenRoutingTableError,
    refetch: listenRoutingTable,
  } = useListenRoutingTableQuery(host.hostId);

  if (isLoading) {
    return <CircularProgress />;
  }

  const handleClick = (data: any) => {
    navigate(`/hosts/${host.hostId}/routing-tables?bucket=${data.bucket}`);
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
          onClick={async () => {
            await saveRoutingTable(host.hostId);
            dispatch(
              snackbarActions.addNotification({
                message: "Saved routing table snapshot",
                variant: "success",
                key: new Date().getTime() + Math.random(),
              })
            );
          }}
          disabled={isSavingRoutingTable}
        >
          Save Snapshot
        </Button>
        <Button
          variant="outlined"
          onClick={async () => {
            await getCurrentRoutingTablePeers(host.hostId);
            dispatch(
              snackbarActions.addNotification({
                message: "Refreshed routing table",
                variant: "success",
                key: new Date().getTime() + Math.random(),
              })
            );
          }}
          disabled={isLoadingCurrentRoutingTablePeers}
        >
          Refetch
        </Button>
      </Stack>
    </Paper>
  );
};

export default HostDetailsRoutingTableCard;
