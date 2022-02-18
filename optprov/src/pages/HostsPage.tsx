import { useEffect, useState } from "react";
import { OptProvClient } from "../api";
import Button from "@mui/material/Button";
import { Host } from "../models/Host";
import { useParams, useNavigate } from "react-router-dom";
import Stack from "@mui/material/Stack";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemText from "@mui/material/ListItemText";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemButton from "@mui/material/ListItemButton";
import IconButton from "@mui/material/IconButton";
import Grid from "@mui/material/Grid";
import DeleteIcon from "@mui/icons-material/Delete";
import FolderIcon from "@mui/icons-material/Folder";
import RootLayout from "../layouts/RootLayout";

function HostsPage() {
  const navigate = useNavigate();
  const [hosts, setHosts] = useState<Host[]>([]);

  const client = new OptProvClient("http://localhost:7000/v1");

  useEffect(() => {
    client.hostsList().then(setHosts);
  }, []);

  return (
    <RootLayout>
      <h1>Hosts</h1>
      <Grid container direction="row" justifyContent="flex-end" alignItems="center">
        <Button
          onClick={() => {
            client.hostsList().then(setHosts);
          }}
        >
          Refresh
        </Button>
      </Grid>
      <Stack spacing={2}>
        <List>
          {hosts.map((host) => (
            <ListItem
              key={host.host_id}
              disablePadding
              secondaryAction={
                <IconButton
                  color="error"
                  edge="end"
                  onClick={async () => {
                    await client.deleteHost(host.host_id);
                    client.hostsList().then(setHosts);
                  }}
                >
                  <DeleteIcon />
                </IconButton>
              }
            >
              <ListItemButton role={undefined} onClick={() => navigate(`/hosts/${host.host_id}`)} dense>
                <ListItemIcon>
                  <FolderIcon />
                </ListItemIcon>
                <ListItemText primary={host.host_id} secondary={"Created at " + host.created_at.toLocaleString()} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
        <Button
          variant="outlined"
          onClick={async () => {
            await client.hostsCreate();
            client.hostsList().then(setHosts);
          }}
        >
          Create Host
        </Button>
      </Stack>
    </RootLayout>
  );
}

export default HostsPage;
