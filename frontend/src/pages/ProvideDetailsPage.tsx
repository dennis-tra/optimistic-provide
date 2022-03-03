import { useParams } from "react-router-dom";
import HostDetailsLayout from "../layouts/HostDetailsLayout";
import LinearProgress from "@mui/material/LinearProgress";
import { Alert, Grid } from "@mui/material";
import { useGetHostQuery, useGetProvideQuery } from "../store/api";
import ProvideOverviewCard from "../components/cards/ProvideOverviewCard";

const ProvideDetailsPage: React.FC = (props) => {
  const { hostId, provideId } = useParams();
  const {
    data: hostData,
    isLoading: isHostLoading,
    isError: isGetHostError,
    error: hostError,
  } = useGetHostQuery(hostId!);
  const {
    data: provideData,
    isLoading: isProvideLoading,
    isError: isGetProvideError,
    error: provideError,
  } = useGetProvideQuery({
    hostId: hostId!,
    provideId: provideId!,
  });

  if (isHostLoading || isProvideLoading) {
    return (
      <HostDetailsLayout hostId={hostId!} title="Provide Operations">
        <LinearProgress />
      </HostDetailsLayout>
    );
  }

  if (isGetHostError || isGetProvideError) {
    return (
      <HostDetailsLayout hostId={hostId!} title="Provide Operations">
        <Alert severity="error">{JSON.stringify(hostError || provideError)}</Alert>
      </HostDetailsLayout>
    );
  }

  return (
    <HostDetailsLayout hostId={hostId!} title="Provide Operations">
      <Grid item xs={12} md={8} lg={6}>
        <ProvideOverviewCard provide={provideData!} host={hostData!} />
      </Grid>
    </HostDetailsLayout>
  );
};

export default ProvideDetailsPage;
