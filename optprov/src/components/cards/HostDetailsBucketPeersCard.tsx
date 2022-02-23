import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import { RoutingTablePeer } from "../../api";
import TableHead from "@mui/material/TableHead";
import TableContainer from "@mui/material/TableContainer";
import Table from "@mui/material/Table";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import ReactTimeAgo from "react-time-ago";

interface HostDetailsBucketPeersCardProps {
  bucket: number;
  peers: RoutingTablePeer[];
}

const HostDetailsBucketPeersCard: React.FC<HostDetailsBucketPeersCardProps> = ({ bucket, peers }) => {
  return (
    <Paper
      sx={{
        p: 2,
        display: "flex",
        flexDirection: "column",
        height: 300,
      }}
    >
      <Typography component="h2" variant="h6" color="primary" gutterBottom>
        Routing Table
      </Typography>
      <>
        <h3>Bucket {bucket}</h3>
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
            <TableHead>
              <TableRow>
                <TableCell>Peer ID</TableCell>
                <TableCell>Agent Version</TableCell>
                <TableCell>Added at</TableCell>
                <TableCell>Last Outbound</TableCell>
                <TableCell>Connected since</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {peers.map((peer) => (
                <TableRow key={peer.peerId} sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
                  <TableCell sx={{ fontFamily: "Monospace" }}>{peer.peerId}</TableCell>
                  <TableCell sx={{ fontFamily: "Monospace" }}>{peer.agentVersion}</TableCell>
                  <TableCell>
                    <ReactTimeAgo date={new Date(peer.addedAt)} />
                  </TableCell>
                  <TableCell>
                    <ReactTimeAgo date={new Date(peer.lastSuccessfulOutboundQueryAt)} />
                  </TableCell>
                  <TableCell>
                    {peer.connectedSince ? <ReactTimeAgo date={new Date(peer.connectedSince)} /> : "-"}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>{" "}
      </>
    </Paper>
  );
};

export default HostDetailsBucketPeersCard;
