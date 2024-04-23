import { Outlet } from "react-router-dom";
import { Box } from "@mui/material";
import { SideBar } from "../../components/SideBar";
import { StickyFooter } from "../../components/StickyFooter";
import CssBaseline from "@mui/material/CssBaseline";

const Layout = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        height: "100vh",
      }}
    >
      <CssBaseline />
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
        <SideBar />
        <Box
          sx={{
            width: "100%",
            overflowY: "auto",
            borderRadius: 2,
            display: "flex",
            flexDirection: "column",
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
