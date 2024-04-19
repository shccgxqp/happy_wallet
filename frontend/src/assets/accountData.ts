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
        Name: "吃飯",
        Amount: 400,
        SharingMethod: "均分",
        SharedBy: [
          { name: "jim", price: 133.33 },
          { name: "bob", price: 133.33 },
          { name: "alice", price: 133.33 },
        ],
        Date: "2022/01/01",
        Payer: "jim",
      },
      {
        Name: "購物",
        Amount: 800,
        SharingMethod: "均分",
        SharedBy: [
          { name: "jim", price: 266.67 },
          { name: "bob", price: 266.67 },
          { name: "alice", price: 266.67 },
        ],
        Date: "2022/01/01",
        Payer: "jim",
      },
      {
        Name: "門票",
        Amount: 1200,
        SharingMethod: "均分",
        SharedBy: [
          { name: "jim", price: 333.33 },
          { name: "bob", price: 333.33 },
          { name: "alice", price: 333.33 },
        ],
        Date: "2022/01/01",
        Payer: "jim",
      },
      {
        Name: "租車",
        Amount: 1500,
        SharingMethod: "均分",
        SharedBy: [
          { name: "jim", price: 416.67 },
          { name: "bob", price: 416.67 },
          { name: "alice", price: 416.67 },
        ],
        Date: "2022/01/01",
        Payer: "jim",
      },
      {
        Name: "酒店",
        Amount: 3000,
        SharingMethod: "均分",
        SharedBy: [
          { name: "jim", price: 666.67 },
          { name: "bob", price: 666.67 },
          { name: "alice", price: 666.67 },
        ],
        Date: "2022/01/01",
        Payer: "jim",
      },
      {
        Name: "紀念品",
        Amount: 600,
        SharingMethod: "均分",
        SharedBy: [
          { name: "jim", price: 166.67 },
          { name: "bob", price: 166.67 },
          { name: "alice", price: 166.67 },
        ],
        Date: "2022/01/01",
        Payer: "jim",
      },
    ],
  };

  export  const pieData = [
    { value: 100, label: "jim" },
    { value: 20, label: "bob" },
    { value: 50, label: "alice" },
  ];