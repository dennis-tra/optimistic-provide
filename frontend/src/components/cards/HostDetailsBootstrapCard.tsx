import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import Tooltip from "@mui/material/Tooltip";
import Box from "@mui/material/Box";
import Chip from "@mui/material/Chip";
import LinearProgress from "@mui/material/LinearProgress";
import CardActions from "@mui/material/CardActions";
import Button from "@mui/material/Button";
import ReactTimeAgo from "react-time-ago";
import { Host } from "../../api";
import { useBootstrapHostMutation } from "../../store/api";

interface HostDetailsBootstrapCardProps {
  host: Host;
}

const HostDetailsBootstrapCard: React.FC<HostDetailsBootstrapCardProps> = ({ host }) => {
  const [bootstrapHost, { isLoading }] = useBootstrapHostMutation(host.hostId);
  return (
    <Card sx={{ display: "flex", justifyContent: "space-between", flexDirection: "column" }}>
      {isLoading ? <LinearProgress /> : null}
      <CardContent>
        <Typography component="h2" variant="h6" color="primary" gutterBottom>
          Status{" "}
          {host.bootstrappedAt ? (
            <Chip label="Bootstrapped" color="success" size="small" />
          ) : (
            <Chip label="Not Bootstrapped" color="warning" size="small" variant="outlined" />
          )}
        </Typography>
        <Typography variant="h6">Bootstrapped At</Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          {host.bootstrappedAt ? (
            <Tooltip title={<ReactTimeAgo date={new Date(host.bootstrappedAt)} />}>
              <span>{new Date(host.bootstrappedAt).toLocaleString()}</span>
            </Tooltip>
          ) : (
            "n.a."
          )}
        </Typography>
        {host.bootstrappedAt ? null : (
          <CardActions>
            <Button onClick={() => bootstrapHost(host.hostId)} variant="contained" disabled={isLoading}>
              Bootstrap
            </Button>
          </CardActions>
        )}
      </CardContent>
    </Card>
  );
};

export default HostDetailsBootstrapCard;
