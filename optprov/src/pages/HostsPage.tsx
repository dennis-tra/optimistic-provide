import { useEffect, useState } from "react";
import { HostsApi, Host } from "../api";
import Button from "@mui/material/Button";
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

  const client = new HostsApi();

  useEffect(() => {
    client.getHosts().then(setHosts);
  }, []);

  return (
    <RootLayout>
      <h1>Hosts</h1>
      <Grid container direction="row" justifyContent="flex-end" alignItems="center">
        <Button
          onClick={() => {
            client.getHosts().then(setHosts);
          }}
        >
          Refresh
        </Button>
      </Grid>
      <Stack spacing={2}>
        <List>
          {hosts.map((host) => (
            <ListItem
              key={host.hostId}
              disablePadding
              secondaryAction={
                <IconButton
                  color="error"
                  edge="end"
                  onClick={async () => {
                    await client.deleteHost({ hostId: host.hostId });
                    client.getHosts().then(setHosts);
                  }}
                >
                  <DeleteIcon />
                </IconButton>
              }
            >
              <ListItemButton role={undefined} onClick={() => navigate(`/hosts/${host.hostId}`)} dense>
                <ListItemIcon>
                  <FolderIcon />
                </ListItemIcon>
                <ListItemText primary={host.hostId} secondary={"Created at " + host.createdAt.toLocaleString()} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
        <Button
          variant="outlined"
          onClick={async () => {
            await client.createHost({ createHostRequest: { name: "some name" } });
            client.getHosts().then(setHosts);
          }}
        >
          Create Host
        </Button>
      </Stack>
    </RootLayout>
  );
}

export default HostsPage;
