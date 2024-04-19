import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
} from "@mui/material";
import { AccountDialogProps } from "../../types/AccountDialogTypes";
const AddDialog: React.FC<AccountDialogProps> = ({ open, setOpen }) => {
  const handleClose = () => {
    setOpen(false);
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      aria-labelledby="alert-dialog-title-add"
      aria-describedby="alert-dialog-description-add"
    >
      <DialogTitle id="alert-dialog-title-add">詳細資料</DialogTitle>
      <DialogContent id="alert-dialog-description-add">
        <DialogContentText>新增交易</DialogContentText>
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
