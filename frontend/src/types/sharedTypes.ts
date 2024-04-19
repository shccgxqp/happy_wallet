export type SharedPerson = {
  name: string;
  price: number;
};

export type Expense = {
  Name: string;
  Amount: number;
  SharingMethod: string;
  SharedBy: SharedPerson[];
  Date: string;
  Payer: string;
};

export type AccountCardListProps = {
  expenses: Expense[];
};