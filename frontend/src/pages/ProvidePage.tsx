import { useParams, Link as RouterLink } from "react-router-dom";
import HostDetailsLayout from "../layouts/HostDetailsLayout";
import CircularProgress from "@mui/material/CircularProgress";
import ReactTimeAgo from "react-time-ago";
import {
  Grid,
  Paper,
  List,
  ListItemButton,
  IconButton,
  ListItem,
  ListItemText,
  Chip,
  Typography,
  Stack,
  Button,
  Tooltip,
  LinearProgress,
} from "@mui/material";
import { useGetProvidesQuery, useLazyGetProvidesQuery, useStartProvideMutation } from "../store/api";
import { ProvideType } from "../api/models/ProvideType";
import ContentCopyIcon from "@mui/icons-material/ContentCopy";

const ProvidePage: React.FC = (props) => {
  const { hostId } = useParams();

  const { data, isLoading } = useGetProvidesQuery(hostId!);
  const [startProvide, { isLoading: isStartingProvide }] = useStartProvideMutation();
  const [getProvides, { isLoading: isGetProvidesLoading }] = useLazyGetProvidesQuery();
  if (isLoading) {
    return (
      <HostDetailsLayout hostId={hostId!} title="Provide Operations">
        <CircularProgress />
      </HostDetailsLayout>
    );
  }

  const provides = data!;

  return (
    <HostDetailsLayout hostId={hostId!} title="Provide Operations">
      <Grid item xs={12} md={12} lg={12}>
        {isGetProvidesLoading && <LinearProgress />}
        <Paper
          sx={{
            p: 2,
            display: "flex",
            flexDirection: "column",
          }}
        >
          <Typography component="h2" variant="h6" color="primary" gutterBottom>
            <Stack direction="row" justifyContent="space-between">
              <span>Provide Operations</span>
              <Stack direction="row" justifyContent="space-between" spacing={2}>
                <Button variant="outlined" onClick={() => getProvides(hostId!)} disabled={isGetProvidesLoading}>
                  Refresh
                </Button>
                <Button
                  variant="contained"
                  onClick={() => startProvide({ hostId: hostId!, body: { type: ProvideType.SingleQuery } })}
                  disabled={isStartingProvide}
                >
                  Start Singe Query
                </Button>
                <Button
                  variant="contained"
                  onClick={() => startProvide({ hostId: hostId!, body: { type: ProvideType.MultiQuery } })}
                  disabled={isStartingProvide}
                >
                  Start Multi Query
                </Button>
              </Stack>
            </Stack>
          </Typography>
          <List component="nav" dense>
            {provides.length == 0 && "This host did not provide any content."}
            {provides.map((provide) => (
              <ListItem
                key={provide.contentId}
                secondaryAction={
                  <Tooltip title={"Copy CID"}>
                    <IconButton edge="end" onClick={() => navigator.clipboard.writeText(provide.contentId)}>
                      <ContentCopyIcon />
                    </IconButton>
                  </Tooltip>
                }
              >
                <ListItemButton component={RouterLink} to={`/hosts/${hostId}/provides/${provide.provideId}`}>
                  <ListItemText
                    primary={
                      <Stack direction="row" spacing={2}>
                        <span>
                          <Typography sx={{ fontFamily: "Monospace" }}>CID: {provide.contentId}</Typography>
                        </span>
                        {provide.endedAt === null && (
                          <Chip label="Ongoging" color="success" variant="outlined" size="small" />
                        )}
                        {provide.endedAt !== null && provide.error === null && (
                          <Tooltip
                            children={<Chip label="Success" color="success" variant="filled" size="small" />}
                            title={`Ended at ${new Date(provide.endedAt).toLocaleString()}`}
                          />
                        )}
                        {provide.endedAt !== null && provide.error !== null && (
                          <Tooltip
                            children={<Chip label="Error" color="error" variant="filled" size="small" />}
                            title={`Error: ${provide.error}`}
                          />
                        )}
                      </Stack>
                    }
                    secondary={
                      <>
                        <span>Started </span>
                        <ReactTimeAgo date={new Date(provide.startedAt)} />
                      </>
                    }
                  />
                </ListItemButton>
              </ListItem>
            ))}
          </List>
        </Paper>
      </Grid>
    </HostDetailsLayout>
  );
};

export default ProvidePage;
