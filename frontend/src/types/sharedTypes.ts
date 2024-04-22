export type SharedPerson = {
  name: string;
  price: number;
};

export type Expense = {
  name: string;
  amount: number;
  sharingMethod: string;
  sharedBy: SharedPerson[];
  date: string;
  payer: string;
};

export type AccountCardListProps = {
  expenses: Expense[];
};