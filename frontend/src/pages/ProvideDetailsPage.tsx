import { useParams, Link as RouterLink } from "react-router-dom";
import HostDetailsLayout from "../layouts/HostDetailsLayout";
import CircularProgress from "@mui/material/CircularProgress";
import ReactTimeAgo from "react-time-ago";
import {
  Grid,
  Paper,
  List,
  ListItemButton,
  Breadcrumbs,
  Link,
  ListItemText,
  Chip,
  Typography,
  Stack,
  Button,
  Tooltip,
  LinearProgress,
} from "@mui/material";
import { useGetProvideQuery, useLazyGetProvidesQuery, useStartProvideMutation } from "../store/api";

const ProvideDetailsPage: React.FC = (props) => {
  const { hostId, provideId } = useParams();

  const { data, isLoading } = useGetProvideQuery({ hostId: hostId!, provideId: provideId! });
  if (isLoading) {
    return (
      <HostDetailsLayout hostId={hostId!} title="Provide Operations">
        <CircularProgress />
      </HostDetailsLayout>
    );
  }

  const provide = data!;

  return (
    <HostDetailsLayout hostId={hostId!} title="Provide Operations">
      <Grid item xs={12} md={12} lg={12}>
        Data {hostId} {provideId}
        <h2>CID: {provide.contentId}</h2>
      </Grid>
    </HostDetailsLayout>
  );
};

export default ProvideDetailsPage;
