import { useParams } from "react-router-dom";
import Grid from "@mui/material/Grid";
import HostDetailsLayout from "../layouts/HostDetailsLayout";
import HostDetailsRoutingTableCard from "../components/cards/HostDetailsRoutingTableCard";
import { useGetHostQuery } from "../store/api";
import CircularProgress from "@mui/material/CircularProgress";
import HostDetailsBucketPeersCard from "../components/cards/HostDetailsBucketPeersCard";

const RoutingTablesOverview: React.FC = (props) => {
  const { hostId } = useParams();

  if (!hostId) {
    return <div>No hostId</div>;
  }

  const { data, isLoading, isError, error } = useGetHostQuery(hostId);
  if (isLoading) {
    return (
      <HostDetailsLayout hostId={hostId}>
        <CircularProgress />
      </HostDetailsLayout>
    );
  }
  const host = data!;
  return (
    <HostDetailsLayout hostId={hostId}>
      <Grid item xs={12} md={12} lg={12}>
        <HostDetailsRoutingTableCard host={host} />
      </Grid>
      <Grid item xs={12} md={12} lg={12}></Grid>
    </HostDetailsLayout>
  );
};

export default RoutingTablesOverview;
