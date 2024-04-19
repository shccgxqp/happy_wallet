import { Link, useLocation } from "react-router-dom";
import { Box, Typography } from "@mui/material";
import homeIcon from "../../assets/icons/icon-nav-home.svg";

const navLinks = [
  {
    name: "Home",
    icon: homeIcon,
    link: "/home",
  },
  {
    name: "Login",
    icon: homeIcon,
    link: "/",
  },
  // {
  //   name: "Movies",
  //   // icon: movieIcon,
  //   link: "/movies",
  // },
  // {
  //   name: "TV Series",
  //   // icon: tvSeriesIcon,
  //   link: "/tv-series",
  // },
  // {
  //   name: "Bookmarks",
  //   // icon: bookmarkIcon,
  //   link: "/bookmarks",
  // },
];

const Sidebar: React.FC = () => {
  const { pathname } = useLocation();

  return (
    <Box
      sx={{
        backgroundColor: "#161d2f",
        padding: 2,
        borderRadius: 2,
        display: "flex",
        flexDirection: {
          xs: "row",
          lg: "column",
        },
        alignItems: "center",
        justifyContent: "space-between",
        width: {
          sm: "95%",
          lg: 200,
        },
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: {
            xs: "row",
            lg: "column",
          },
          gap: 5,
          alignItems: {
            xs: "center",
            lg: "start",
          },
          width: "100%",
        }}
      >
        <Box sx={{ display: { xs: "none", sm: "block" } }}>
          <Typography
            variant="h5"
            component="h1"
            my={2}
            fontWeight={400}
            fontSize={18}
          >
            Happy Wallet
          </Typography>
        </Box>

        <Box
          sx={{
            py: {
              xs: "0px",
              lg: "16px",
            },
            display: "flex",
            flexDirection: {
              xs: "row",
              lg: "column",
            },
            gap: 4,
          }}
        >
          {navLinks.map((item) => (
            <Link key={item.name} to={item.link}>
              <Box
                sx={{
                  display: "flex",
                  alignItems: "center",
                  gap: 2,
                  color: "white",
                  textDecoration: "none",
                }}
              >
                <img
                  src={item.icon}
                  alt={item.name}
                  style={{
                    width: "18px",
                    filter: `${
                      pathname === item.link
                        ? "invert(58%) sepia(14%) saturate(3166%) hue-rotate(215deg) brightness(91%) contrast(87%)"
                        : "invert(84%)"
                    }`,
                  }}
                />
                <Box sx={{ display: { xs: "none", md: "block" } }}>
                  <Typography>{item.name}</Typography>
                </Box>
              </Box>
            </Link>
          ))}
        </Box>
      </Box>
    </Box>
  );
};

export default Sidebar;
