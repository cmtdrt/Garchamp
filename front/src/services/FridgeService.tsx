// src/services/frigoService.ts
import { FoodItem } from "@/types/FoodItem";

const API_URL = "http://localhost:5173/api/v1/fridge";

export const fridgeService = {
  async getAll(): Promise<FoodItem[]> {
    const res = await fetch(API_URL);
    if (!res.ok) throw new Error("Erreur lors du chargement");
    return res.json();
  },

  async add(item: FoodItem): Promise<Response> {
    const res = await fetch(API_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item),
    });
    return res;
  },

  async remove(id: string): Promise<void> {
    const res = await fetch(`${API_URL}/${id}`, { method: "DELETE" });
    if (!res.ok) throw new Error("Erreur lors de la suppression");
  },

  async update(id: string, item: Partial<FoodItem>): Promise<FoodItem> {
    const res = await fetch(`${API_URL}/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item),
    });
    if (!res.ok) throw new Error("Erreur lors de la mise à jour");
    return res.json();
  },
};