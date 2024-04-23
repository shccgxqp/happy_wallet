import { useEffect, useState } from "react";
import dayjs from "dayjs";
import {
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Typography,
} from "@mui/material";
import { DemoContainer } from "@mui/x-date-pickers/internals/demo";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";
import CloseIcon from "@mui/icons-material/Close";

import { TraderForm, TraderCheckboxList } from "../Trader";
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
      return;
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
    setAddTrader((prevTrader) => ({
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
    setAddTrader((prevTrader) => ({
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
        {/* 付款人、金額、幣別*/}
        <Typography variant="caption" color="text.secondary">
          付款人
        </Typography>
        <TraderForm
          members={members}
          addTrader={addTrader}
          setAddTrader={setAddTrader}
          handleAmountChange={handleAmountChange}
        />

        {/* 替誰付錢 */}
        <Typography
          variant="caption"
          color="text.secondary"
          display={"block"}
          mt={2}
        >
          替誰付錢
        </Typography>
        <TraderCheckboxList
          members={members}
          addTrader={addTrader}
          handleCheckboxChange={handleCheckboxChange}
        />

        {/* 資料目的名稱、統一編號 */}
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

        {/* 日期與時間 */}
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
