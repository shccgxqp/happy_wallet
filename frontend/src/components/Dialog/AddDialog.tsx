import { useEffect, useState } from "react";
import {
  Button,
  Box,
  Checkbox,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  FormControl,
  Grid,
  Input,
  Stack,
  Select,
  TextField,
  Typography,
  MenuItem,
  OutlinedInput,
} from "@mui/material";

import dayjs from "dayjs";
import { DemoContainer } from "@mui/x-date-pickers/internals/demo";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";

import PersonIcon from "@mui/icons-material/Person";
import CloseIcon from "@mui/icons-material/Close";
import {
  AccountDialogAddProps,
  AccountTrader,
} from "../../types/AccountDialogTypes";

const AddDialog: React.FC<AccountDialogAddProps> = ({
  open,
  setOpen,
  members,
  SelectedModel,
}) => {
  const [addTrader, setAddTrader] = useState<AccountTrader>({
    name: "",
    trader: "",
    amount: 0,
    currency: "TWD",
    uniformInvoice: "",
    paymentMethod: "evenSplit", //Options: evenSplit, ratioSplit, amountSplit
    recipients: new Array(members.length).fill(true),
    amountPerRecipient: new Array(members.length).fill(0),
    date: dayjs(new Date()),
  });
  useEffect(() => {
    if (SelectedModel === "edit") {
    }
  }, [SelectedModel]);

  const handleClose = () => {
    console.log(addTrader);
    setAddTrader({
      name: "",
      trader: "",
      amount: 0,
      currency: "TWD",
      uniformInvoice: "",
      paymentMethod: "evenSplit", //Options: evenSplit, ratioSplit, amountSplit
      recipients: new Array(members.length).fill(true),
      amountPerRecipient: new Array(members.length).fill(0),
      date: dayjs(new Date()),
    });
    setOpen(false);
  };

  const calculateAmountPerRecipient = (
    recipients: boolean[],
    totalAmount: number
  ) => {
    const numRecipients = recipients.filter(Boolean).length;
    return recipients.map((checked: boolean) =>
      checked ? totalAmount / numRecipients : 0
    );
  };

  const handleCheckboxChange = (index: number) => {
    const newMemberCheck = [...addTrader.recipients];
    newMemberCheck[index] = !newMemberCheck[index];
    const amountPerRecipient = calculateAmountPerRecipient(
      newMemberCheck,
      addTrader.amount
    );
    setAddTrader((prevTrader: any) => ({
      ...prevTrader,
      recipients: newMemberCheck,
      amountPerRecipient: amountPerRecipient,
    }));
  };
  const handleAmountChange = (value: string) => {
    const newAmount = parseFloat(value);
    const amountPerRecipient = calculateAmountPerRecipient(
      addTrader.recipients,
      newAmount
    );
    setAddTrader((prevTrader: any) => ({
      ...prevTrader,
      amount: parseFloat(value),
      amountPerRecipient: amountPerRecipient,
    }));
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      aria-labelledby="dialog-add-title"
      aria-describedby="dialog-add-description"
      fullWidth={true}
    >
      <DialogTitle id="dialog-add-title">新增交易</DialogTitle>
      <Button
        aria-label="close"
        onClick={handleClose}
        sx={{
          position: "absolute",
          right: 8,
          top: 8,
          color: (theme) => theme.palette.grey[500],
        }}
      >
        <CloseIcon />
      </Button>
      <DialogContent id="dialog-add-content">
        <Typography variant="caption" color="text.secondary">
          付款人
        </Typography>

        <FormControl sx={{ minWidth: 120, ml: 2 }} size="small">
          <Select
            displayEmpty
            value={addTrader.trader}
            onChange={(e) =>
              setAddTrader({ ...addTrader, trader: e.target.value })
            }
            label="dialog-add-trader"
            input={<OutlinedInput />}
            renderValue={(selected) => {
              if (selected.length === 0) {
                return <em>"請選擇付款人"</em>;
              }
              return selected;
            }}
          >
            <MenuItem disabled value="">
              <em>選擇付款人 : </em>
            </MenuItem>
            {members.map((member, index) => (
              <MenuItem key={index} value={member.name}>
                {member.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>

        <Grid container minHeight={120} alignItems="flex-end">
          <Grid item xs={10}>
            <FormControl fullWidth variant="standard">
              <Input
                id="dialog-add-amount"
                type="number"
                value={addTrader.amount}
                onChange={(e) => handleAmountChange(e.target.value)}
                inputProps={{
                  style: { textAlign: "right", alignSelf: "flex-end" },
                }}
                sx={{
                  height: 120,
                  color: "orange",
                  fontSize: 80,
                }}
              />
            </FormControl>
          </Grid>
          <Grid item xs={2}>
            <Select
              displayEmpty
              value={addTrader.currency}
              onChange={(e) =>
                setAddTrader({ ...addTrader, currency: e.target.value })
              }
              label="dialog-add-currency"
              input={<OutlinedInput />}
              sx={{ color: "orange" }}
            >
              <MenuItem value="TWD">TWD</MenuItem>
              <MenuItem value="USD">USD</MenuItem>
              <MenuItem value="EUR">EUR</MenuItem>
              <MenuItem value="JPY">JPY</MenuItem>
            </Select>
          </Grid>
        </Grid>

        <Typography
          variant="caption"
          color="text.secondary"
          display={"block"}
          mt={2}
        >
          替誰付錢
        </Typography>
        <Stack spacing={2}>
          {members.map((member, index) => (
            <Box
              key={"Box" + index}
              sx={{
                display: "flex",
                alignItems: "center",
                height: 40,
              }}
            >
              <PersonIcon sx={{ fontSize: 40 }} />
              <Box sx={{ flex: 1, textAlign: "left", ml: 1 }}>
                <Typography>{member.name}</Typography>
                <Typography variant="caption" color="text.secondary">
                  ${addTrader.amountPerRecipient[index].toFixed(2)}
                </Typography>
              </Box>
              <Checkbox
                checked={addTrader.recipients[index]}
                onChange={() => handleCheckboxChange(index)}
              />
            </Box>
          ))}
        </Stack>

        <Typography
          variant="caption"
          color="text.secondary"
          display={"block"}
          mt={2}
        >
          目的
        </Typography>
        <TextField
          id="dialog-add-name"
          label="例如 周末朋友一起吃火鍋"
          value={addTrader.name}
          variant="standard"
          onChange={(e) => {
            setAddTrader({ ...addTrader, name: e.target.value });
          }}
          sx={{ width: "100%" }}
        />
        <TextField
          id="dialog-add-uniform-invoice"
          label="統一編號"
          value={addTrader.uniformInvoice}
          variant="standard"
          onChange={(e) => {
            setAddTrader({ ...addTrader, uniformInvoice: e.target.value });
          }}
          sx={{ width: "100%", border: "none" }}
        />

        <Typography
          variant="caption"
          color="text.secondary"
          display={"block"}
          mt={2}
        >
          日期與時間
        </Typography>
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <DemoContainer components={["DateTimePicker"]}>
            <DateTimePicker
              value={addTrader.date}
              onChange={(newValue) => {
                setAddTrader({ ...addTrader, date: newValue });
              }}
            />
          </DemoContainer>
        </LocalizationProvider>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose} color="success" autoFocus>
          新增交易
        </Button>
        <Button onClick={handleClose} color="primary">
          取消交易
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default AddDialog;
