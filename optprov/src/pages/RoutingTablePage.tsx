import { useParams, useSearchParams } from "react-router-dom";
import Grid from "@mui/material/Grid";
import HostDetailsLayout from "../layouts/HostDetailsLayout";
import HostDetailsRoutingTableCard from "../components/cards/HostDetailsRoutingTableCard";
import { useGetHostQuery } from "../store/api";
import { selectBucketPeers } from "../store/bucketsSlice";
import CircularProgress from "@mui/material/CircularProgress";
import HostDetailsBucketPeersCard from "../components/cards/HostDetailsBucketPeersCard";
import { useAppSelector } from "../store/config";

const RoutingTablePage: React.FC = (props) => {
  const { hostId } = useParams();
  const [searchParams] = useSearchParams();

  if (!hostId) {
    return <div>No hostId</div>;
  }

  const bucket = parseInt(searchParams.get("bucket") || "0");
  const peers = useAppSelector(selectBucketPeers(hostId, bucket));

  const { data, isLoading, isError, error } = useGetHostQuery(hostId);
  if (isLoading) {
    return (
      <HostDetailsLayout hostId={hostId} title="Routing Table">
        <CircularProgress />
      </HostDetailsLayout>
    );
  }
  const host = data!;
  return (
    <HostDetailsLayout hostId={hostId} title="Routing Table">
      <Grid item xs={12} md={12} lg={12}>
        <HostDetailsRoutingTableCard host={host} />
      </Grid>
      <Grid item xs={12} md={12} lg={12}>
        <HostDetailsBucketPeersCard bucket={bucket} peers={peers} />
      </Grid>
    </HostDetailsLayout>
  );
};

export default RoutingTablePage;
