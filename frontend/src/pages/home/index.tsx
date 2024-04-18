import { Box, Typography } from "@mui/material";
import { DefaultizedPieValueType } from "@mui/x-charts";
import { PieChart, pieArcLabelClasses } from "@mui/x-charts/PieChart";
const home = () => {
  const data = [
    { value: 100, label: "jim" },
    { value: 20, label: "bob" },
    { value: 50, label: "alice" },
  ];

  const TOTAL = data.map((item) => item.value).reduce((a, b) => a + b, 0);
  const getArcLabel = (params: DefaultizedPieValueType) => {
    const percent = params.value / TOTAL;
    return `${(percent * 100).toFixed(0)}%`;
  };

  return (
    <Box
      height={200}
      width={"auto"}
      my={4}
      display="flex"
      alignItems="center"
      justifyItems="center"
      justifyContent="center"
      bgcolor={"primary.main"}
    >
      <Box m="20px">
        <PieChart
          series={[
            {
              arcLabel: getArcLabel,
              data,
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

      <Box>
        <Typography variant="h4" color="white">
          happy wallet
        </Typography>
      </Box>
    </Box>
  );
};

export default home;
