import { useState } from "react";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import { useCreateHostMutation } from "../store/api";
import { FormControl, InputLabel, MenuItem, Select, SelectChangeEvent } from "@mui/material";
import { NetworkType } from "../api/models/NetworkType";

interface NewHostCardProps {}

const NewHostCard: React.FC<NewHostCardProps> = ({}) => {
  const [createHost, { isLoading }] = useCreateHostMutation();
  const [hostName, setHostName] = useState("");
  const [network, setNetwork] = useState<NetworkType>(NetworkType.Ipfs);

  const handleChange = (event: SelectChangeEvent) => {
    setNetwork(event.target.value as NetworkType);
  };

  const handleSubmit = async (event: React.MouseEvent<HTMLElement>) => {
    event.preventDefault();
    try {
      await createHost({ name: hostName, network: network }).unwrap();
      setHostName("");
    } catch (err) {
      console.error("Failed to create new host: ", err);
    }
  };

  return (
    <Grid item xs={12} md={4} lg={4}>
      <form action="new-host-form">
        <Paper
          sx={{
            p: 2,
            display: "flex",
            flexDirection: "column",
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
          <FormControl>
            <InputLabel id="select-network">Network</InputLabel>
            <Select
              labelId="select-network"
              id="select-network"
              value={network}
              label="Network"
              onChange={handleChange}
            >
              <MenuItem value={NetworkType.Ipfs}>IPFS</MenuItem>
              <MenuItem value={NetworkType.Filecoin}>Filecoin</MenuItem>
              <MenuItem value={NetworkType.Polkadot}>Polkadot</MenuItem>
              <MenuItem value={NetworkType.Kusama}>Kusama</MenuItem>
            </Select>
          </FormControl>
          <Box sx={{ flex: 1 }} />
          <Button
            type="submit"
            variant="contained"
            onClick={handleSubmit}
            disabled={!hostName || !network || isLoading}
          >
            Create
          </Button>
        </Paper>
      </form>
    </Grid>
  );
};

export default NewHostCard;
