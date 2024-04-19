import React, { useState } from "react";
import { Box, Typography, Button } from "@mui/material";
import { DefaultizedPieValueType } from "@mui/x-charts";
import { PieChart, pieArcLabelClasses } from "@mui/x-charts/PieChart"; // 请根据实际导入的圆饼图组件路径来替换这里的占位符
import { AddDialog, EditDialog } from "../Dialog";

import { formatNumber } from "../../utils/utils";

type PieChartSectionProps = {
  accountData: any;
  pieData: Array<{ label: string; value: number }>;
};

const AccountPieChartSection: React.FC<PieChartSectionProps> = ({
  accountData,
  pieData,
}) => {
  const [openAddDialog, setOpenAddDialog] = useState(false);
  const [openEditDialog, setOpenEditDialog] = useState(false);

  const TOTAL = pieData.map((item) => item.value).reduce((a, b) => a + b, 0);
  const getArcLabel = (params: DefaultizedPieValueType) => {
    const percent = params.value / TOTAL;
    return `${(percent * 100).toFixed(0)}%`;
  };

  const handleAddDialogOpen = () => {
    setOpenAddDialog(true);
  };
  const handleEditDialogOpen = () => {
    setOpenEditDialog(true);
  };
  return (
    <>
      <Box sx={{ height: { xs: 200, sm: 300 }, width: { xs: 200, sm: 300 } }}>
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
          margin={{ right: 5 }}
          slotProps={{
            legend: { hidden: true },
          }}
        />
      </Box>
      <Box display={"flex"} flexDirection={"column"} alignItems={"start"}>
        <Typography variant="h4" color="white">
          {accountData.title}
        </Typography>
        <Typography variant="body2" color="#f5f5f5" ml={1}>
          成員: {accountData.transactions.count}
        </Typography>
        <Typography variant="body2" color="#f5f5f5" ml={1}>
          交易次數: {accountData.transactions.tradeCount}
        </Typography>
        <Typography variant="body2" color="#f5f5f5" ml={1}>
          總支出: ${formatNumber(accountData.transactions.amount)}
        </Typography>
        <Box sx={{ marginTop: 2 }}>
          <Button
            variant="contained"
            color="success"
            onClick={handleAddDialogOpen}
          >
            新增交易
          </Button>
          <Button variant="contained" onClick={handleEditDialogOpen}>
            编辑群组
          </Button>
        </Box>

        <AddDialog open={openAddDialog} setOpen={setOpenAddDialog} />
        <EditDialog open={openEditDialog} setOpen={setOpenEditDialog} />
      </Box>
    </>
  );
};

export default AccountPieChartSection;
