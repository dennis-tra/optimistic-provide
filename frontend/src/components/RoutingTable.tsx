import TableHead from "@mui/material/TableHead";
import TableContainer from "@mui/material/TableContainer";
import Table from "@mui/material/Table";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import Paper from "@mui/material/Paper";
import TableBody from "@mui/material/TableBody";
import { useEffect, useState } from "react";
import { Host, RoutingTablePeer, RoutingTablePeerFromJSON, RoutingTablePeerFromJSONTyped } from "../api";
import RoutingTableHistogram from "./RoutingTableHistogram";
import ReactTimeAgo from "react-time-ago";

interface RoutingTableProps {
  host: Host;
}

interface RoutingTable {
  [key: number]: RoutingTablePeer[];
}

const buckets = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15];

const newRoutingTable = (): RoutingTable => {
  const routingTable: RoutingTable = {};
  for (const bucket in buckets) {
    routingTable[bucket] = [];
  }
  return routingTable;
};

const RoutingTable: React.FC<RoutingTableProps> = (props) => {
  const [connectionState, setConnectionState] = useState<"loading" | "connected" | "error" | "closed">("loading");
  const [routingTable, setRoutingTable] = useState<RoutingTable>(newRoutingTable());
  const [selectedBucket, setSelectedBucket] = useState<number | null>(null);

  useEffect(() => {
    const websocket = new WebSocket(`ws://localhost:7000/hosts/${props.host.hostId}/routing-tables/listen`);
    websocket.onmessage = (event) => {
      const newRT = newRoutingTable();
      for (const update of JSON.parse(event.data)) {
        const routingTableUpdate = RoutingTablePeerFromJSON(update);
        newRT[update.bucket].push(routingTableUpdate);
      }
      setRoutingTable(newRT);
    };
    websocket.onopen = () => setConnectionState("connected");
    websocket.onerror = () => setConnectionState("error");
    websocket.onclose = () => setConnectionState("closed");
    return () => {
      websocket.close();
    };
  }, []);

  if (connectionState === "loading") {
    return <p>Loading...</p>;
  }

  if (connectionState === "closed") {
    return <p>Connection got closed.</p>;
  }

  const bucketLevels = buckets.map((bucket) => routingTable[bucket].length);
  return (
    <>
      <RoutingTableHistogram bucketLevels={bucketLevels} onBucketSelect={setSelectedBucket} />
      {selectedBucket !== null && (
        <>
          <h3>Bucket {selectedBucket}</h3>
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
                {routingTable[selectedBucket].map((peer) => (
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
      )}
    </>
  );
};

export default RoutingTable;
