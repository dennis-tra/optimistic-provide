import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Divider from "@mui/material/Divider";
import Typography from "@mui/material/Typography";
import Tooltip from "@mui/material/Tooltip";
import LinearProgress from "@mui/material/LinearProgress";
import Button from "@mui/material/Button";
import ReactTimeAgo from "react-time-ago";
import { Host } from "../../api/models/Host";
import { useBootstrapHostMutation, useStartHostMutation, useStopHostMutation } from "../../store/api";

interface HostDetailsBootstrapCardProps {
  host: Host;
}

const HostDetailsStatusCard: React.FC<HostDetailsBootstrapCardProps> = ({ host }) => {
  const [bootstrapHost, { isLoading: isBootstrappingHost }] = useBootstrapHostMutation();
  const [startHost, { isLoading: isStartingHost }] = useStartHostMutation();
  const [stopHost, { isLoading: isStoppingHost }] = useStopHostMutation();

  return (
    <Card sx={{ display: "flex", justifyContent: "space-between", flexDirection: "column" }}>
      {isBootstrappingHost || isStartingHost ? <LinearProgress /> : null}
      <CardContent>
        <Typography component="h2" variant="h6" color="primary" gutterBottom>
          Status
        </Typography>
        <Typography variant="h6">Started</Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          {host.startedAt && (
            <Tooltip title={new Date(host.startedAt).toLocaleString()}>
              <span>
                <ReactTimeAgo date={new Date(host.startedAt)} />
              </span>
            </Tooltip>
          )}
          {!host.startedAt && (
            <Button onClick={() => startHost(host.hostId)} variant="contained" disabled={isStartingHost}>
              Start
            </Button>
          )}
        </Typography>
        <Typography variant="h6">Bootstrapped</Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          {host.bootstrappedAt && (
            <Tooltip title={new Date(host.bootstrappedAt).toLocaleString()}>
              <span>
                <ReactTimeAgo date={new Date(host.bootstrappedAt)} />
              </span>
            </Tooltip>
          )}
          {!host.bootstrappedAt && (
            <Button
              onClick={() => bootstrapHost(host.hostId)}
              variant="contained"
              disabled={isBootstrappingHost || !host.startedAt}
            >
              Bootstrap
            </Button>
          )}
        </Typography>
        {host.startedAt && (
          <>
            <Divider />
            <Button
              variant="outlined"
              color="primary"
              sx={{ mt: 2 }}
              onClick={() => stopHost(host.hostId)}
              disabled={isStoppingHost}
            >
              Stop
            </Button>
          </>
        )}
      </CardContent>
    </Card>
  );
};

export default HostDetailsStatusCard;
