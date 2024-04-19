import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
} from "@mui/material";
import { AccountDialogProps } from "../../types/AccountDialogTypes";
const EditDialog: React.FC<AccountDialogProps> = ({ open, setOpen }) => {
  const handleClose = () => {
    setOpen(false);
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      aria-labelledby="alert-dialog-title-edit"
      aria-describedby="alert-dialog-description-edit"
    >
      <DialogTitle id="alert-dialog-title-edit">詳細資料</DialogTitle>
      <DialogContent id="alert-dialog-description-edit">
        <DialogContentText>編輯群組</DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose} color="primary">
          關閉
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default EditDialog;
