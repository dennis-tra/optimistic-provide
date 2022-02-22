import { useState } from "react";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";
import Tooltip from "@mui/material/Tooltip";
import Snackbar from "@mui/material/Snackbar";
import Alert from "@mui/material/Alert";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import { Host, HostsApi } from "../api";

interface NewHostCardProps {
  reload: () => Promise<void>;
}

const client = new HostsApi();

const NewHostCard: React.FC<NewHostCardProps> = ({ reload }) => {
  const [hostName, setHostName] = useState("");

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    if (hostName === "") {
      return;
    }
    await client.createHost({ createHostRequest: { name: hostName } });
    setHostName("");
    await reload();
  };

  return (
    <Grid item xs={12} md={4} lg={4}>
      <form action="new-host-form">
        <Paper
          sx={{
            p: 2,
            display: "flex",
            flexDirection: "column",
            height: 240,
            gap: 2,
          }}
        >
          <Typography component="p" variant="h4">
            New Host
          </Typography>
          <TextField
            id="outlined-basic"
            label="Host Name"
            variant="outlined"
            value={hostName}
            onChange={(event) => setHostName(event.target.value)}
          />
          <Box sx={{ flex: 1 }}></Box>
          <Button type="submit" variant="contained" onClick={handleSubmit}>
            Create
          </Button>
        </Paper>
      </form>
    </Grid>
  );
};

export default NewHostCard;
