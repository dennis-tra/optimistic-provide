import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { Host } from "../../api/models/Host";
import { useGetProvideGraphQuery } from "../../store/api";
import { ProvideDetails } from "../../api/models/ProvideDetails";
import ProvideChart from "../ProvideChart";
import { LinearProgress } from "@mui/material";

interface ProvideChartCardProps {
  host: Host;
  provide: ProvideDetails;
}

const ProvideChartCard: React.FC<ProvideChartCardProps> = ({ provide, host }) => {
  if (!provide.endedAt) {
    return (
      <Card sx={{ minHeight: 275 }}>
        <CardContent>Still providing...</CardContent>
      </Card>
    );
  }

  const {
    data: graphData,
    isLoading: isGraphDataLoading,
    isError,
  } = useGetProvideGraphQuery({
    hostId: host.hostId,
    provideId: provide.provideId,
  });

  if (isGraphDataLoading) {
    return (
      <Card sx={{ minHeight: 275 }}>
        <CardContent>
          <LinearProgress />
          Loading Graph Data...
        </CardContent>
      </Card>
    );
  }

  if (isError) {
    return (
      <Card sx={{ minHeight: 275 }}>
        <CardContent>An error occurred..</CardContent>
      </Card>
    );
  }

  return (
    <Card sx={{ minHeight: 275 }}>
      <CardContent>
        <Typography component="h2" variant="h6" color="primary" gutterBottom>
          Provide Operation
        </Typography>
        <ProvideChart data={graphData!} />
      </CardContent>
    </Card>
  );
};

export default ProvideChartCard;
