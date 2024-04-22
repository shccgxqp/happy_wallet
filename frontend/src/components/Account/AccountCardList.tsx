import { useState } from "react";
import {
  Box,
  Card,
  CardContent,
  Avatar,
  Typography,
  CardActionArea,
  Stack,
  Grid,
} from "@mui/material";
import PersonIcon from "@mui/icons-material/Person";
import { CartDialog } from "../Dialog";

import { AccountCardListProps, Expense } from "../../types/sharedTypes";
import { formatNumber } from "../../utils/utils";

const AccountCardList: React.FC<AccountCardListProps> = ({ expenses }) => {
  const [open, setOpen] = useState(false);
  const [selectedData, setSelectedData] = useState<Expense>({
    Name: "",
    Amount: 0,
    SharingMethod: "",
    SharedBy: [],
    Date: "",
    Payer: "",
  });
  console.log(expenses);

  const handleClickOpen = (data: Expense) => {
    setSelectedData(data);
    setOpen(true);
  };

  return (
    <Stack
      sx={{
        direction: "column",
        spacing: 1,
        borderRadius: "10px",
        width: { xs: "100%", md: 800 },
        bgcolor: "white",
      }}
    >
      {expenses.map((entry, index) => (
        <Card key={index} sx={{ height: 100 }}>
          <CardActionArea onClick={() => handleClickOpen(entry)}>
            <CardContent>
              <Grid container direction="row">
                <Grid item xs={2} my={"auto"}>
                  <PersonIcon sx={{ fontSize: 40 }} />
                </Grid>
                <Grid item xs={6} style={{ textAlign: "left" }}>
                  <Typography variant="h6">{entry.Name}</Typography>
                  <Typography
                    variant="caption"
                    color="text.secondary"
                    display="block"
                  >
                    {entry.Date}
                  </Typography>
                  <Typography variant="caption" color=" " display="block">
                    {entry.Payer} 支付
                  </Typography>
                </Grid>
                <Grid item xs={4}>
                  <Box
                    display={"flex"}
                    sx={{ flexDirection: "column", alignItems: "flex-end" }}
                  >
                    <Typography color={"primary"}>
                      $ {formatNumber(entry.Amount)}
                    </Typography>
                    <Box
                      display={"flex"}
                      justifyContent={"flex-end"}
                      alignItems={"center"}
                      mt={1}
                    >
                      {entry.SharedBy.map((person, index) => (
                        <Avatar
                          alt={person.name}
                          key={index}
                          sx={{ width: 24, height: 24 }}
                        >
                          {person.name[0]}
                        </Avatar>
                      ))}
                    </Box>
                  </Box>
                </Grid>
              </Grid>
            </CardContent>
          </CardActionArea>
        </Card>
      ))}
      <CartDialog open={open} setOpen={setOpen} data={selectedData} />
    </Stack>
  );
};

export default AccountCardList;
