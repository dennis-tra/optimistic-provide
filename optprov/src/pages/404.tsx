import { Link } from "react-router-dom";
import RootLayout from "../layouts/RootLayout";
import { Container, Stack, Typography } from "@mui/material";

const NotFoundPage: React.FC = () => {
  return (
    <RootLayout>
      <Typography variant="h3" gutterBottom>
        404 Not Found
      </Typography>
      <Link to={`/hosts`}>Go Back</Link>
    </RootLayout>
  );
};

export default NotFoundPage;
