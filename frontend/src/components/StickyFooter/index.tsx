import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import Link from "@mui/material/Link";

function Copyright() {
  return (
    <Typography variant="body2" color="text.secondary">
      {"Copyright © "}
      <Link color="inherit" href="https://mui.com/">
        Wang Yuan Chen
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}

const StickyFooter = () => {
  return (
    <Box
      component="footer"
      sx={{
        py: { xs: 1, sm: 3 },
        px: { xs: 1, sm: 2 },
        backgroundColor: (theme) =>
          theme.palette.mode === "light"
            ? theme.palette.grey[200]
            : theme.palette.grey[800],
      }}
    >
      <Container maxWidth="sm">
        <Typography variant="body2">
          {" "}
          Built with React, TypeScript, and Material-UI.
        </Typography>
        <Copyright />
      </Container>
    </Box>
  );
};

export default StickyFooter;
