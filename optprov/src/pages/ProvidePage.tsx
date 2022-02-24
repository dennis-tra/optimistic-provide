import { useParams } from "react-router-dom";
import HostDetailsLayout from "../layouts/HostDetailsLayout";
import CircularProgress from "@mui/material/CircularProgress";

const ProvidePage: React.FC = (props) => {
  const { hostId } = useParams();
  return (
    <HostDetailsLayout hostId={hostId!}>
      <CircularProgress />
    </HostDetailsLayout>
  );
};

export default ProvidePage;
