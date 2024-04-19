import { Box, Typography, Button } from "@mui/material";
import { DefaultizedPieValueType } from "@mui/x-charts";
import { PieChart, pieArcLabelClasses } from "@mui/x-charts/PieChart";
import AccountCardList from "../../components/AccountCardList";

import { accountData, pieData } from "../../assets/accountData";
import { formatNumber } from "../../utils/utils";

const home = () => {
  const TOTAL = pieData.map((item) => item.value).reduce((a, b) => a + b, 0);
  const getArcLabel = (params: DefaultizedPieValueType) => {
    const percent = params.value / TOTAL;
    return `${(percent * 100).toFixed(0)}%`;
  };

  return (
    <>
      <Box
        height={200}
        width={"auto"}
        my={1}
        display="flex"
        alignItems="center"
        justifyItems="center"
        justifyContent="center"
        borderRadius="10px"
        bgcolor={"#3F51B5"}
      >
        <Box m="20px">
          <PieChart
            series={[
              {
                arcLabel: getArcLabel,
                data: pieData,
              },
            ]}
            sx={{
              [`& .${pieArcLabelClasses.root}`]: {
                fill: "black",
                fontSize: 14,
              },
            }}
            slotProps={{
              legend: { hidden: true },
            }}
            width={300}
            height={200}
          />
        </Box>
        <Box display={"flex"} flexDirection={"column"} alignItems={"start"}>
          <Typography variant="h4" color="white">
            {accountData.title}
          </Typography>
          <Typography variant="body2" color="white">
            成員: {accountData.transactions.count}
          </Typography>
          <Typography variant="body2" color="white">
            交易次數: {accountData.transactions.tradeCount}
          </Typography>
          <Typography variant="body2" color="white">
            總支出: ${formatNumber(accountData.transactions.amount)}
          </Typography>
          <Box sx={{ marginTop: 2 }}>
            <Button variant="contained" color="success">
              新增交易
            </Button>
            <Button variant="contained">编辑群组</Button>
          </Box>
        </Box>
      </Box>
      <Box
        display={"flex"}
        flexDirection={"column"}
        justifyItems={"center"}
        alignItems={"center"}
      >
        <AccountCardList expenses={accountData.expenses} />
      </Box>
    </>
  );
};

export default home;
