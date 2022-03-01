import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { ProvideDetails } from "../../api/models/ProvideDetails";
import { Host } from "../../api/models/Host";
import { actions as snackbarActions } from "../../store/snackbarSlice";
import { useAppDispatch } from "../../store/config";
import { Tooltip } from "@mui/material";

interface ProvideOverviewCardProps {
  host: Host;
  provide: ProvideDetails;
}

const ProvideOverviewCard: React.FC<ProvideOverviewCardProps> = ({ provide, host }) => {
  const dispatch = useAppDispatch();
  let durationInS = 0;
  if (provide.endedAt !== null) {
    const start = new Date(provide.startedAt);
    const end = new Date(provide.endedAt);
    durationInS = (end.getTime() - start.getTime()) / 1000;
  }

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
          onClick={async () => {
            await navigator.clipboard.writeText(provide.contentId);
            dispatch(
              snackbarActions.addNotification({
                key: new Date().getTime() + Math.random(),
                variant: "success",
                message: "CID copied to clipboard!",
              })
            );
          }}
          sx={{ cursor: "pointer", mb: 1.5 }}
        >
          {provide.contentId}
        </Typography>
        <Typography variant="h6">Provider Peer ID</Typography>
        <Typography
          color="text.secondary"
          noWrap
          onClick={async () => {
            await navigator.clipboard.writeText(host.hostId);
            dispatch(
              snackbarActions.addNotification({
                key: new Date().getTime() + Math.random(),
                variant: "success",
                message: "Provider Peer ID copied to clipboard!",
              })
            );
          }}
          sx={{ cursor: "pointer", mb: 1.5 }}
        >
          {host.hostId}
        </Typography>
        <Typography variant="h6">Distance</Typography>
        <Typography color="text.secondary" noWrap sx={{ mb: 1.5 }}>
          {Math.round((1e4 * 100 * parseInt(provide.distance)) / 2 ** 256) / 1e4} %
        </Typography>
        <Typography variant="h6">Duration</Typography>
        <Typography color="text.secondary" noWrap sx={{ mb: 1.5 }}>
          {durationInS === 0 ? "Still providing" : `${durationInS} s`}
        </Typography>
        <Typography variant="h6">Counts</Typography>
        <Typography color="text.secondary" noWrap sx={{ mb: 1.5 }}>
          {provide.dials.length} Dials / {provide.connections.length} Connections / {provide.findNodes.length} FIND_NODE
          / {provide.addProviders.length} ADD_PROVIDER
        </Typography>
      </CardContent>
    </Card>
  );
};

export default ProvideOverviewCard;
