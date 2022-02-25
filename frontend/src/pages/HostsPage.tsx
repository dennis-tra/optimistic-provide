import Grid from "@mui/material/Grid";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import RootLayout from "../layouts/RootLayout";
import HostCard from "../components/HostCard";
import NewHostCard from "../components/NewHostCard";
import { useGetHostsQuery } from "../store/api";

function RootPage() {
  const { data: hosts, refetch } = useGetHostsQuery();

  return (
    <RootLayout>
      <Typography component="h2" variant="h2" color="primary" gutterBottom>
        Hosts{" "}
        <Button variant="outlined" onClick={refetch}>
          Refresh
        </Button>
      </Typography>

      <Grid container spacing={3}>
        {hosts?.map((host, idx) => (
          <HostCard key={host.hostId} host={host} />
        ))}
        <NewHostCard />
      </Grid>
    </RootLayout>
  );
}

export default RootPage;
