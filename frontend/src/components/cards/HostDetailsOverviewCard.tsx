import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import Tooltip from "@mui/material/Tooltip";
import ReactTimeAgo from "react-time-ago";
import { Host } from "../../api/models/Host";
import { useAppDispatch } from "../../store/config";
import { actions as snackbarActions } from "../../store/snackbarSlice";

interface HostDetailsOverviewCardProps {
  host: Host;
}

const HostDetailsOverviewCard: React.FC<HostDetailsOverviewCardProps> = ({ host }) => {
  const dispatch = useAppDispatch();
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
        <Typography
          color="text.secondary"
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
          sx={{ cursor: "pointer", mb: 1.5 }}
        >
          {host.hostId}
        </Typography>
        <Typography variant="h6">Created At</Typography>
        <Typography color="text.secondary" noWrap>
          <Tooltip title={new Date(host.createdAt).toLocaleString()}>
            <span>
              <ReactTimeAgo date={new Date(host.createdAt)} />
            </span>
          </Tooltip>
        </Typography>
      </CardContent>
    </Card>
  );
};

export default HostDetailsOverviewCard;
