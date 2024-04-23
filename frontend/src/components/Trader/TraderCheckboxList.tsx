import React from "react";
import Stack from "@mui/material/Stack";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Checkbox from "@mui/material/Checkbox";
import PersonIcon from "@mui/icons-material/Person";

interface MemberCheckboxListProps {
  members: Array<{ name: string }>;
  addTrader: {
    amountPerRecipient: number[];
    recipients: boolean[];
  };
  handleCheckboxChange: (index: number) => void;
}

const TraderCheckboxList: React.FC<MemberCheckboxListProps> = ({
  members,
  addTrader,
  handleCheckboxChange,
}) => {
  return (
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
  );
};

export default TraderCheckboxList;
