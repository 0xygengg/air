// apps/frontend/app/wallet/page.tsx
"use client";

import { useState } from "react";
import { api } from "@/lib/api";

export default function WalletPage() {
  const [wallet, setWallet] = useState<{ address: string } | null>(null);

  const generateWallet = async () => {
    const res = await api.get("/wallet/create");
    setWallet(res.data);
  };

  return (
    <main className="p-8 max-w-xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">🔐 Generate Wallet</h1>
      <button
        onClick={generateWallet}
        className="bg-black text-white px-4 py-2 rounded"
      >
        Generate
      </button>
      {wallet && (
        <div className="mt-4 break-words">
          <p><b>Public Address:</b> {wallet.address}</p>
        </div>
      )}
    </main>
  );
}
