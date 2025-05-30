// apps/frontend/app/transaction/page.tsx
"use client";

import { useState } from "react";
import { api } from "@/lib/api";

export default function TransactionPage() {
  const [sender, setSender] = useState("");
  const [recipient, setRecipient] = useState("");
  const [amount, setAmount] = useState<number>(0);
  const [status, setStatus] = useState("");

  const sendTransaction = async () => {
    try {
      await api.post("/transaction", {
        sender,
        recipient,
        amount,
      });
      setStatus("✅ Transaction sent");
    } catch {
      setStatus("❌ Failed to send");
    }
  };

  return (
    <main className="p-8 max-w-xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">💸 Send Transaction</h1>
      <input
        type="text"
        placeholder="Sender Address"
        value={sender}
        onChange={(e) => setSender(e.target.value)}
        className="w-full p-2 mb-2 border rounded"
      />
      <input
        type="text"
        placeholder="Recipient Address"
        value={recipient}
        onChange={(e) => setRecipient(e.target.value)}
        className="w-full p-2 mb-2 border rounded"
      />
      <input
        type="number"
        placeholder="Amount"
        value={amount}
        onChange={(e) => setAmount(parseFloat(e.target.value))}
        className="w-full p-2 mb-4 border rounded"
      />
      <button
        onClick={sendTransaction}
        className="bg-blue-600 text-white px-4 py-2 rounded"
      >
        Send
      </button>
      {status && <p className="mt-4">{status}</p>}
    </main>
  );
}
