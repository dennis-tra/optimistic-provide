import { useTheme } from "@mui/material/styles";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import { ResponsiveContainer, BarChart, Bar, CartesianGrid, XAxis, YAxis, Tooltip } from "recharts";
import { Host, RoutingTablePeer } from "../../api";
import { useGetRoutingTablePeersQuery } from "../../store/api";
import { CircularProgress } from "@mui/material";
import { useNavigate } from "react-router-dom";

interface RoutingTable {
  [key: number]: RoutingTablePeer[];
}

const buckets = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18];

const newRoutingTable = (): RoutingTable => {
  const routingTable: RoutingTable = {};
  for (const bucket in buckets) {
    routingTable[bucket] = [];
  }
  return routingTable;
};

interface HostDetailsRoutingTableCardProps {
  host: Host;
}

const HostDetailsRoutingTableCard: React.FC<HostDetailsRoutingTableCardProps> = ({ host }) => {
  const navigate = useNavigate();
  const { data, isLoading } = useGetRoutingTablePeersQuery(host.hostId);

  if (isLoading || !data) {
    return <CircularProgress />;
  }

  const newRT = newRoutingTable();
  for (const peer of data!) {
    newRT[peer.bucket].push(peer);
  }
  const bucketLevels = buckets.map((bucket) => ({ bucket: bucket, level: newRT[bucket].length }));

  const handleClick = (data: any, index: number) => {
    navigate(`/hosts/${host.hostId}/routing-tables?bucket=${data.bucket}`);
  };

  return (
    <Paper
      sx={{
        p: 2,
        display: "flex",
        flexDirection: "column",
        height: 300,
      }}
    >
      <Typography component="h2" variant="h6" color="primary" gutterBottom>
        Routing Table
      </Typography>
      <ResponsiveContainer>
        <BarChart width={150} height={40} data={bucketLevels}>
          <CartesianGrid strokeDasharray="3 3" />
          <Tooltip />
          <XAxis dataKey="bucket" />
          <YAxis />
          <Bar dataKey="level" fill="#82ca9d" onClick={handleClick} />
        </BarChart>
      </ResponsiveContainer>
    </Paper>
  );
};

export default HostDetailsRoutingTableCard;
