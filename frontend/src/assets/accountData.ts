export const accountData = {
  title: "東海岸旅遊",
  transactions: {
    group: "美國行",
    count: 3,
    tradeCount: 30,
    amount: 47000.47,
    date: "2022/01/01",
    update: "2022/01/02",
    category: "旅遊",
  },
  members: [
    { name: "jim", role: "管理員" },
    { name: "bob", role: "經理" },
    { name: "alice", role: "員工" },
  ],
  expenses: [
    {
      name: "吃飯",
      amount: 400,
      sharingMethod: "均分",
      sharedBy: [
        { name: "jim", price: 133.33 },
        { name: "bob", price: 133.33 },
        { name: "alice", price: 133.33 },
      ],
      date: "2022/01/01",
      payer: "jim",
    },
    {
      name: "購物",
      amount: 800,
      sharingMethod: "均分",
      sharedBy: [
        { name: "jim", price: 266.67 },
        { name: "bob", price: 266.67 },
        { name: "alice", price: 266.67 },
      ],
      date: "2022/01/01",
      payer: "jim",
    },
    {
      name: "門票",
      amount: 1200,
      sharingMethod: "均分",
      sharedBy: [
        { name: "jim", price: 333.33 },
        { name: "bob", price: 333.33 },
        { name: "alice", price: 333.33 },
      ],
      date: "2022/01/01",
      payer: "jim",
    },
    {
      name: "租車",
      amount: 1500,
      sharingMethod: "均分",
      sharedBy: [
        { name: "jim", price: 416.67 },
        { name: "bob", price: 416.67 },
        { name: "alice", price: 416.67 },
      ],
      date: "2022/01/01",
      payer: "jim",
    },
    {
      name: "酒店",
      amount: 3000,
      sharingMethod: "均分",
      sharedBy: [
        { name: "jim", price: 666.67 },
        { name: "bob", price: 666.67 },
        { name: "alice", price: 666.67 },
      ],
      date: "2022/01/01",
      payer: "jim",
    },
    {
      name: "紀念品",
      amount: 600,
      sharingMethod: "均分",
      sharedBy: [
        { name: "jim", price: 166.67 },
        { name: "bob", price: 166.67 },
        { name: "alice", price: 166.67 },
      ],
      date: "2022/01/01",
      payer: "jim",
    },
  ],
};

export const pieData = [
  { value: 100, label: "jim" },
  { value: 20, label: "bob" },
  { value: 50, label: "alice" },
];
