import { useState, useEffect } from "react";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import Box from "@mui/material/Box";
import RootLayout from "../layouts/RootLayout";
import HostCard from "../components/HostCard";
import { Host, HostsApi } from "../api";
import NewHostCard from "../components/NewHostCard";

function RootPage() {
  const [hosts, setHosts] = useState<Host[]>([]);

  const client = new HostsApi();

  const reload = async () => {
    const hosts = await client.getHosts();
    setHosts(hosts);
  };

  useEffect(() => {
    reload();
  }, []);

  return (
    <RootLayout>
      <Typography component="h2" variant="h2" color="primary" gutterBottom>
        Hosts
      </Typography>
      <Grid container spacing={3}>
        {hosts.map((host, idx) => (
          <HostCard key={host.hostId} host={host} idx={idx} reload={reload} />
        ))}
        <NewHostCard reload={reload} />
      </Grid>
    </RootLayout>
  );
}

export default RootPage;
