// apps/frontend/app/nft/page.tsx
"use client";

import { useState } from "react";
import { api } from "@/lib/api";

export default function NFTMint() {
  const [form, setForm] = useState({
    name: "",
    description: "",
    imageUrl: "",
    owner: "",
  });
  const [status, setStatus] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const mint = async () => {
    try {
      const res = await api.post("/nft/mint", form);
      setStatus("✅ Minted: " + JSON.stringify(res.data.data));
    } catch (e) {
      setStatus("❌ Mint failed");
    }
  };

  return (
    <main className="p-8 max-w-xl mx-auto space-y-4">
      <h1 className="text-2xl font-bold">🎨 Mint NFT</h1>
      {["name", "description", "imageUrl", "owner"].map((field) => (
        <input
          key={field}
          name={field}
          placeholder={field}
          className="w-full p-2 border rounded"
          onChange={handleChange}
        />
      ))}
      <button onClick={mint} className="bg-black text-white px-4 py-2 rounded">
        Mint
      </button>
      {status && <p>{status}</p>}
    </main>
  );
}
