import { useParams, Link as RouterLink } from "react-router-dom";
import HostDetailsLayout from "../layouts/HostDetailsLayout";
import CircularProgress from "@mui/material/CircularProgress";
import ReactTimeAgo from "react-time-ago";
import {
  Grid,
  Paper,
  List,
  ListItemButton,
  Box,
  ListItemText,
  Chip,
  Typography,
  Stack,
  Button,
  Tooltip,
  LinearProgress,
} from "@mui/material";
import { useGetRetrievalsQuery, useLazyGetRetrievalsQuery, useStartRetrievalMutation } from "../store/api";
import { useState } from "react";
import TextField from "@mui/material/TextField";

const RetrievalsPage: React.FC = (props) => {
  const { hostId } = useParams();

  const [cid, setContentId] = useState("");

  const { data, isLoading } = useGetRetrievalsQuery(hostId!);
  const [startRetrieval, { isLoading: isStartingRetrieval }] = useStartRetrievalMutation();
  const [getRetrievals, { isLoading: isGetRetrievalsLoading }] = useLazyGetRetrievalsQuery();
  if (isLoading) {
    return (
      <HostDetailsLayout hostId={hostId!} title="Retrieval Operations">
        <CircularProgress />
      </HostDetailsLayout>
    );
  }

  const retrievals = data!;

  return (
    <HostDetailsLayout hostId={hostId!} title="Retrieval Operations">
      <Grid item xs={12} md={6} lg={4}>
        <Paper
          sx={{
            p: 2,
            display: "flex",
            flexDirection: "column",
            minHeight: 240,
          }}
        >
          <Typography component="h2" variant="h6" color="primary" gutterBottom>
            New Retrieval
          </Typography>
          <Typography variant="caption" gutterBottom>
            Enter a CID to find provider records for
          </Typography>
          <TextField
            id="outlined-basic"
            label="CID"
            variant="outlined"
            value={cid}
            onChange={(event) => setContentId(event.target.value)}
          />
          <Box flex={1} />
          <Button
            variant="outlined"
            onClick={() => {
              startRetrieval({ hostId: hostId!, body: { contentId: cid, count: 0 } });
              setContentId("");
            }}
            disabled={isStartingRetrieval || !cid}
          >
            Start Retrieval
          </Button>
        </Paper>
      </Grid>
      <Grid item xs={12} md={12} lg={12}>
        {isGetRetrievalsLoading && <LinearProgress />}
        <Paper
          sx={{
            p: 2,
            display: "flex",
            flexDirection: "column",
          }}
        >
          <Typography component="h2" variant="h6" color="primary" gutterBottom>
            <Stack direction="row" justifyContent="space-between">
              <span>Retrieval Operations</span>
              <Stack direction="row" justifyContent="space-between" spacing={2}>
                <Button variant="outlined" onClick={() => getRetrievals(hostId!)} disabled={isGetRetrievalsLoading}>
                  Refresh
                </Button>
              </Stack>
            </Stack>
          </Typography>
          <List component="nav" dense>
            {retrievals.length == 0 && "This host did not retrieve any content."}
            {retrievals.map((retrieval) => (
              <ListItemButton
                key={retrieval.retrievalId}
                component={RouterLink}
                to={`/hosts/${hostId}/retrievals/${retrieval.retrievalId}`}
              >
                <ListItemText
                  primary={
                    <Stack direction="row" spacing={2}>
                      <span>
                        <Typography sx={{ fontFamily: "Monospace" }}>CID: {retrieval.contentId}</Typography>
                      </span>
                      {retrieval.endedAt === null && (
                        <Chip label="Ongoging" color="success" variant="outlined" size="small" />
                      )}
                      {retrieval.endedAt !== null && retrieval.error === null && (
                        <Tooltip
                          children={<Chip label="Success" color="success" variant="filled" size="small" />}
                          title={`Ended at ${new Date(retrieval.endedAt).toLocaleString()}`}
                        />
                      )}
                      {retrieval.endedAt !== null && retrieval.error !== null && (
                        <Tooltip
                          children={<Chip label="Error" color="success" variant="filled" size="small" />}
                          title={`Error: ${retrieval.error}`}
                        />
                      )}
                    </Stack>
                  }
                  secondary={
                    <>
                      <span>Started </span>
                      <ReactTimeAgo date={new Date(retrieval.startedAt)} />
                    </>
                  }
                />
              </ListItemButton>
            ))}
          </List>
        </Paper>
      </Grid>
    </HostDetailsLayout>
  );
};

export default RetrievalsPage;
