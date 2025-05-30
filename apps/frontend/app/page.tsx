// apps/frontend/app/page.tsx
"use client";

import { useEffect, useState } from "react";
import { api } from "@/lib/api";

type Transaction = {
  sender: string;
  recipient: string;
  amount: number;
};

type Block = {
  timestamp: number;
  transactions: Transaction[];
  hash: string;
  prevHash: string;
};

export default function Home() {
  const [blocks, setBlocks] = useState<Block[]>([]);

  useEffect(() => {
    const fetchBlocks = async () => {
      const res = await api.get("/blocks");
      setBlocks(res.data.reverse()); // Show newest first
    };
    fetchBlocks();
  }, []);

  return (
    <main className="p-8 max-w-4xl mx-auto">
      <h1 className="text-3xl font-bold mb-6">🌐 Air Blockchain Explorer</h1>
      <div className="space-y-4">
        {blocks.map((block, idx) => (
          <div key={idx} className="p-4 border rounded-xl shadow">
            <p><b>Hash:</b> {block.hash}</p>
            <p><b>Previous:</b> {block.prevHash}</p>
            <p><b>Time:</b> {new Date(block.timestamp * 1000).toLocaleString()}</p>
            <p><b>Transactions:</b> {block.transactions?.length ?? 0}</p>
          </div>
        ))}
      </div>
    </main>
  );
}
