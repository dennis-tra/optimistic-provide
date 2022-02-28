import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { ProvideDetails } from "../../api/models/ProvideDetails";
import { Host } from "../../api/models/Host";

interface ProvideOverviewCardProps {
  host: Host;
  provide: ProvideDetails;
}

const ProvideOverviewCard: React.FC<ProvideOverviewCardProps> = ({ provide, host }) => {
  return (
    <Card sx={{ minHeight: 275 }}>
      <CardContent>
        <Typography component="h2" variant="h6" color="primary" gutterBottom>
          Overview
        </Typography>
        <Typography variant="h6">CID</Typography>
        <Typography
          color="text.secondary"
          noWrap
          onClick={() => navigator.clipboard.writeText(provide.contentId)}
          sx={{ cursor: "pointer", mb: 1.5 }}
        >
          {provide.contentId}
        </Typography>
        <Typography variant="h6">Provider Peer ID</Typography>
        <Typography
          color="text.secondary"
          noWrap
          onClick={() => navigator.clipboard.writeText(host.hostId)}
          sx={{ cursor: "pointer", mb: 1.5 }}
        >
          {host.hostId}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default ProvideOverviewCard;
