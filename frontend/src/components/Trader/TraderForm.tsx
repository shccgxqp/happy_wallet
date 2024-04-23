import React from "react";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
import Grid from "@mui/material/Grid";
import Input from "@mui/material/Input";
import OutlinedInput from "@mui/material/OutlinedInput";
import { AccountTrader, members } from "../../types/AccountDialogTypes";

interface AddTraderProps {
  addTrader: AccountTrader;
  setAddTrader: (value: AccountTrader) => void;
  members: members[];
  handleAmountChange: (value: string) => void;
}

const TraderForm: React.FC<AddTraderProps> = ({
  addTrader,
  setAddTrader,
  members,
  handleAmountChange,
}) => {
  return (
    <React.Fragment>
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
    </React.Fragment>
  );
};

export default TraderForm;
