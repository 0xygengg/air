// apps/frontend/lib/api.ts
import axios from "axios";

export const api = axios.create({
  baseURL: typeof window === "undefined"
    ? "http://localhost:8080"
    : "http://127.0.0.1:8080", // Fixes localhost on Windows
});
