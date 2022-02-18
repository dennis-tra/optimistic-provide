import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { OptProvClient } from "../api";
import RootLayout from "../layouts/RootLayout";
import { Host } from "../models/Host";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemText from "@mui/material/ListItemText";
import IconButton from "@mui/material/IconButton";
import CloudIcon from "@mui/icons-material/Cloud";
import RoutingTable from "../components/RoutingTable";
import Button from "@mui/material/Button";

const HostPage: React.FC = (props) => {
  const { hostId } = useParams();
  if (!hostId) {
    return <div>No hostId</div>;
  }

  const [host, setHost] = useState<Host | null>(null);

  const client = new OptProvClient("http://localhost:7000/v1");
  useEffect(() => {
    client.hostGet(hostId).then(setHost);
  }, []);

  if (!host) {
    return <RootLayout></RootLayout>;
  }

  return (
    <RootLayout>
      <h1>Host</h1>
      <h2>General</h2>
      <List>
        <ListItem>
          <ListItemText primary="PeerID" secondary={host.host_id} />
        </ListItem>
        <ListItem>
          <ListItemText primary="Created at" secondary={host.created_at} />
        </ListItem>
        <ListItem
          secondaryAction={
            host.bootstrapped_at ? null : (
              <IconButton
                edge="end"
                onClick={async () => {
                  client.hostBootstrap(host.host_id).then(setHost);
                }}
              >
                <CloudIcon />
              </IconButton>
            )
          }
        >
          <ListItemText primary="Bootstrapped at" secondary={host.bootstrapped_at || "n.a."} />
        </ListItem>
      </List>
      <Button onClick={() => client.refreshRoutingTable(hostId)}>Refresh Routing Table</Button>
      <h2>Routing Table</h2>
      <RoutingTable host={host} />
    </RootLayout>
  );
};

export default HostPage;
