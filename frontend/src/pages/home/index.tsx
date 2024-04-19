import { Box } from "@mui/material";
import { AccountCardList, AccountPieChart } from "../../components/Account";
import { accountData, pieData } from "../../assets/accountData";
const home = () => {
  return (
    <>
      <Box
        sx={{
          display: "flex",
          flexDirection: { xs: "column", sm: "row" },
          justifyContent: "center",
          alignItems: "center",
          borderRadius: "10px",
          bgcolor: "#3F51B5",
          width: "auto",
          height: "auto",
          mb: 1,
        }}
      >
        <AccountPieChart accountData={accountData} pieData={pieData} />
      </Box>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <AccountCardList expenses={accountData.expenses} />
      </Box>
    </>
  );
};

export default home;
