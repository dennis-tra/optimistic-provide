import React from "react";
import Container from "@mui/material/Container";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";

const RootLayout: React.FC = ({ children }): JSX.Element => {
  return (
    <>
      <AppBar position="static">
        <Toolbar variant="dense">
          <Typography variant="h6" color="inherit" component="div">
            DHT Provide
          </Typography>
        </Toolbar>
      </AppBar>
      <Container fixed>{children}</Container>
    </>
  );
};

export default RootLayout;
