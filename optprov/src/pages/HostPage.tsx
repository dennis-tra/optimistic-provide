import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import RootLayout from "../layouts/RootLayout";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemText from "@mui/material/ListItemText";
import IconButton from "@mui/material/IconButton";
import CloudIcon from "@mui/icons-material/Cloud";
import RoutingTable from "../components/RoutingTable";
import { HostsApi, Host } from "../api";

const HostPage: React.FC = (props) => {
  const { hostId } = useParams();
  if (!hostId) {
    return <div>No hostId</div>;
  }

  const [host, setHost] = useState<Host | null>(null);

  const client = new HostsApi();
  useEffect(() => {
    client.getHost({ hostId }).then(setHost);
  }, []);

  if (!host) {
    return <div>Loading...</div>;
  }

  return (
    <RootLayout>
      <h1>Host</h1>
      <h2>General</h2>
      <List>
        <ListItem>
          <ListItemText primary="PeerID" secondary={host.hostId} />
        </ListItem>
        <ListItem>
          <ListItemText primary="Created at" secondary={host.createdAt} />
        </ListItem>
        <ListItem
          secondaryAction={
            host.bootstrappedAt ? null : (
              <IconButton edge="end" onClick={async () => client.bootstrapHost({ hostId }).then(setHost)}>
                <CloudIcon />
              </IconButton>
            )
          }
        >
          <ListItemText primary="Bootstrapped at" secondary={host.bootstrappedAt || "n.a."} />
        </ListItem>
      </List>
      <h2>Routing Table</h2>
      <RoutingTable host={host} />
    </RootLayout>
  );
};

export default HostPage;
