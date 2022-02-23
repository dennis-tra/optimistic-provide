import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import Tooltip from "@mui/material/Tooltip";
import ReactTimeAgo from "react-time-ago";
import { Host } from "../../api";

interface HostDetailsOverviewCardProps {
  host: Host;
}

const HostDetailsOverviewCard: React.FC<HostDetailsOverviewCardProps> = ({ host }) => {
  return (
    <Card sx={{ minHeight: 275 }}>
      <CardContent>
        <Typography component="h2" variant="h6" color="primary" gutterBottom>
          Overview
        </Typography>
        <Typography variant="h6">Name</Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          {host.name}
        </Typography>
        <Typography variant="h6">Peer ID</Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary" noWrap>
          {host.hostId}
        </Typography>
        <Typography variant="h6">Created At</Typography>
        <Typography color="text.secondary" noWrap>
          <Tooltip title={<ReactTimeAgo date={new Date(host.createdAt)} />}>
            <span>{new Date(host.createdAt).toLocaleString()}</span>
          </Tooltip>
        </Typography>
      </CardContent>
    </Card>
  );
};

export default HostDetailsOverviewCard;
