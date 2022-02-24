import { useTheme } from "@mui/material/styles";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import { ResponsiveContainer, BarChart, Bar, CartesianGrid, XAxis, YAxis, Tooltip } from "recharts";
import { Host } from "../../api";
import { useGetRoutingTablePeersQuery } from "../../store/api";
import { CircularProgress } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { useAppSelector } from "../../store/config";
import { selectHistogramData } from "../../store/bucketsSlice";

interface HostDetailsRoutingTableCardProps {
  host: Host;
}

const HostDetailsRoutingTableCard: React.FC<HostDetailsRoutingTableCardProps> = ({ host }) => {
  const navigate = useNavigate();
  const bucketData = useAppSelector(selectHistogramData(host.hostId));
  const { isLoading } = useGetRoutingTablePeersQuery(host.hostId);

  if (isLoading) {
    return <CircularProgress />;
  }

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
        <BarChart width={150} height={40} data={bucketData}>
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
