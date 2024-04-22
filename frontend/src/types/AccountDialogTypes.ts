import {Dayjs } from 'dayjs';
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
  SelectedModel: 'add'|'edit';
};

export type AccountTrader = {
  name: string,
    trader: string,
    amount: number,
    currency: string,
    uniformInvoice: string,
    paymentMethod: 'evenSplit'| 'ratioSplit'| 'amountSplit', 
    recipients: Array<boolean>,
    amountPerRecipient: Array<number>,
    date: Dayjs | null,
}
