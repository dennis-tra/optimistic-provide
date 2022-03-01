import { useParams } from "react-router-dom";
import Grid from "@mui/material/Grid";
import CircularProgress from "@mui/material/CircularProgress";
import HostDetailsLayout from "../layouts/HostDetailsLayout";
import { useGetHostQuery } from "../store/api";
import HostDetailsOverviewCard from "../components/cards/HostDetailsOverviewCard";
import HostDetailsStatusCard from "../components/cards/HostDetailsStatusCard";
import HostDetailsRoutingTableCard from "../components/cards/HostDetailsRoutingTableCard";
import { Snackbar } from "@mui/material";

const HostPage: React.FC = (props) => {
  const { hostId } = useParams();
  if (!hostId) {
    return <div>No hostId</div>;
  }

  const { data, isLoading, isError, error } = useGetHostQuery(hostId);

  if (isError) {
    return <Snackbar></Snackbar>;
  }

  if (isLoading) {
    return (
      <HostDetailsLayout hostId={hostId} title="General">
        <CircularProgress />
      </HostDetailsLayout>
    );
  }

  const host = data!;

  return (
    <HostDetailsLayout hostId={hostId} title={`General`}>
      <Grid item xs={12} md={6} lg={4}>
        <HostDetailsOverviewCard host={host} />
      </Grid>
      <Grid item xs={12} md={6} lg={4}>
        <HostDetailsStatusCard host={host} />
      </Grid>
      <Grid item xs={12} md={12} lg={10}>
        <HostDetailsRoutingTableCard host={host} />
      </Grid>
    </HostDetailsLayout>
  );
};

export default HostPage;
