import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
} from "@mui/material";
import { AccountCartDialogProps } from "../../types/AccountDialogTypes";

const CartDialog: React.FC<AccountCartDialogProps> = ({
  open,
  setOpen,
  data,
}) => {
  const handleClose = () => {
    setOpen(false);
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      aria-labelledby="alert-dialog-title-cart"
      aria-describedby="alert-dialog-description-cart"
    >
      <DialogTitle id="alert-dialog-title-cart">詳細資料</DialogTitle>
      <DialogContent id="alert-dialog-description-cart">
        <DialogContentText>{data.Name}</DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose} color="primary">
          關閉
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default CartDialog;
