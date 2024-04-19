import { useState } from "react";
import {
  Box,
  Card,
  CardContent,
  Avatar,
  Typography,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
  CardActionArea,
  Stack,
  Grid,
} from "@mui/material";
import PersonIcon from "@mui/icons-material/Person";

import { AccountCardListProps, Expense } from "../../types/sharedTypes";
import { formatNumber } from "../../utils/utils";

const AccountCardList = ({ expenses }: AccountCardListProps) => {
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

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <Stack
      direction={"column"}
      spacing={1}
      sx={{ width: { xs: "100%", md: 800 }, bgcolor: "white" }}
    >
      {expenses.map((entry, index) => (
        <Card sx={{ height: 100 }}>
          <CardActionArea onClick={() => handleClickOpen(entry)}>
            <CardContent>
              <Grid container>
                <Grid xs={2} my={"auto"}>
                  <PersonIcon sx={{ fontSize: 40 }} />
                </Grid>
                <Grid xs={6} style={{ textAlign: "left" }}>
                  <Typography variant="h6">{entry.Name}</Typography>
                  <Typography
                    variant="caption"
                    color="text.secondary"
                    display="block"
                  >
                    {entry.Date}
                  </Typography>
                  <Typography
                    variant="caption"
                    color="text.secondary"
                    display="block"
                  >
                    {entry.Payer} 支付
                  </Typography>
                </Grid>
                <Grid
                  xs={4}
                  display={"flex"}
                  direction={"column"}
                  textAlign={"right"}
                  justifyContent={"center"}
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
                </Grid>
              </Grid>
            </CardContent>
          </CardActionArea>
        </Card>
      ))}
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>詳細資料</DialogTitle>
        <DialogContent>
          <DialogContentText>{selectedData.Name}</DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary">
            關閉
          </Button>
        </DialogActions>
      </Dialog>
    </Stack>
  );
};

export default AccountCardList;
