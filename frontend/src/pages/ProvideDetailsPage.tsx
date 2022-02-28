import { useParams } from "react-router-dom";
import HostDetailsLayout from "../layouts/HostDetailsLayout";
import LinearProgress from "@mui/material/LinearProgress";
import { Grid } from "@mui/material";
import { useGetHostQuery, useGetProvideQuery, useLazyGetProvidesQuery, useStartProvideMutation } from "../store/api";
import ProvideOverviewCard from "../components/cards/ProvideOverviewCard";

const ProvideDetailsPage: React.FC = (props) => {
  const { hostId, provideId } = useParams();
  const { data: hostData, isLoading: isHostLoading } = useGetHostQuery(hostId!);
  const { data: provideData, isLoading: isProvideLoading } = useGetProvideQuery({
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

  return (
    <HostDetailsLayout hostId={hostId!} title="Provide Operations">
      <Grid item xs={12} md={12} lg={12}>
        <ProvideOverviewCard provide={provideData!} host={hostData!} />
      </Grid>
    </HostDetailsLayout>
  );
};

export default ProvideDetailsPage;
