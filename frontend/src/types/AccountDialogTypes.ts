export type AccountCartDialogProps = {
  open: boolean;
  setOpen: (value: boolean) => void;
  data: any;
};

export type AccountDialogProps = {
  open: boolean;
  setOpen: (value: boolean) => void;
};

export type AccountDialogAddProps = {
  open: boolean;  
  setOpen: (value: boolean) => void;
  members: Array<{name: string, role: string}>;
};