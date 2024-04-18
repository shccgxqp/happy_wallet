import { Box } from "@mui/material";
import { Outlet } from "react-router-dom";
import Sidebar from "../components/sidebar";
import StickyFooter from "../components/sticky-footer";

const Layout = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        height: "100vh",
      }}
    >
      <Box
        sx={{
          backgroundColor: "#10141F",
          display: "flex",
          flexDirection: {
            xs: "column",
            lg: "row",
          },
          color: "white",
          padding: 3,
          gap: 3,
          overflowY: "auto",
          flex: "1",
        }}
      >
        <Sidebar />
        <Box
          sx={{
            width: "100%",
            overflowY: "auto",
            borderRadius: 2,
          }}
        >
          <Outlet />
        </Box>
      </Box>
      <StickyFooter />
    </div>
  );
};

export default Layout;
