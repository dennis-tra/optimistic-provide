import TableHead from "@mui/material/TableHead";
import TableContainer from "@mui/material/TableContainer";
import Table from "@mui/material/Table";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import Paper from "@mui/material/Paper";
import TableBody from "@mui/material/TableBody";
import { useEffect, useState } from "react";
import { Host } from "../models/Host";
import { RoutingTableUpdate, Convert } from "../models/RoutingTableUpdate";
import RoutingTableHistogram from "./RoutingTableHistogram";
import ReactTimeAgo from "react-time-ago";

interface RoutingTableProps {
  host: Host;
}

interface RoutingTable {
  [key: number]: RoutingTableUpdate[];
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
    const websocket = new WebSocket(`ws://localhost:7000/v1/hosts/${props.host.host_id}/routing-tables/listen`);
    websocket.onmessage = (event) => {
      const routingTableUpdate = Convert.toRoutingTableUpdate(event.data);
      const newRT = newRoutingTable();
      for (const update of routingTableUpdate) {
        newRT[update.bucket].push(update);
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
                <TableRow key={peer.peer_id} sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
                  <TableCell sx={{ fontFamily: "Monospace" }}>{peer.peer_id}</TableCell>
                  <TableCell sx={{ fontFamily: "Monospace" }}>{peer.agent_version}</TableCell>
                  <TableCell>
                    <ReactTimeAgo date={peer.added_at} />
                  </TableCell>
                  <TableCell>
                    <ReactTimeAgo date={peer.last_successful_outbound_query_at} />
                  </TableCell>
                  <TableCell>{peer.connected_at ? <ReactTimeAgo date={peer.connected_at} /> : "-"}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      )}
    </>
  );
};

export default RoutingTable;
