// src/services/frigoService.ts
import { FoodItem } from "@/types/FoodItem";

const API_URL = "http://localhost:8080/api/v1/fridge";
const STORAGE_KEY = "frigo_items";

export const fridgeService = {
  async getAll(): Promise<FoodItem[]> {
    const res = await fetch(API_URL, { headers: { Accept: "application/json" } });

    if (!res.ok) {
      const text = await res.text();
      console.error("Erreur API :", res.status, text);
      throw new Error(`Erreur API ${res.status}`);
    }

    try {
      const contentType = res.headers.get("content-type") || "";
      if (!contentType.includes("application/json")) {
        const text = await res.text();
        console.error("Réponse non JSON reçue:", text.slice(0, 200));
        return [];
      }
      const data = await res.json();

      if (data && Array.isArray(data.data)) {
        try {
          localStorage.setItem(STORAGE_KEY, JSON.stringify(data.data));
        } catch (e) {
          console.warn("Impossible d'écrire dans localStorage", e);
        }
        return data.data; 
      }

      console.error("Structure inattendue :", data);
      return [];
    } catch (err) {
      console.error("Erreur de parsing JSON :", err);
      return [];
    }
  },


  async add(item: FoodItem): Promise<Response> {
    type PostBody = {
      name: string;
      quantity: number;
      unit: string;
      exp_date?: string;
      allergens?: string[];
    };
    const unitField = (item as unknown as { unit?: string; unity?: string }).unit ?? (item as unknown as { unit?: string; unity?: string }).unity ?? "unité";
    const payload: PostBody = {
      name: item.name,
      quantity: item.quantity,
      unit: unitField,
      exp_date: item.exp_date,
      allergens: item.allergens,
    };
    const res = await fetch(API_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });
    if (res.ok) {
      try {
        const items = await fridgeService.getAll();
        localStorage.setItem(STORAGE_KEY, JSON.stringify(items));
      } catch (e) {
        console.warn("Impossible de rafraîchir le localStorage après ajout", e);
      }
    }
    return res;
  },

  async remove(id: number): Promise<void> {
    const res = await fetch(`${API_URL}/${id}`, { method: "DELETE" });
    if (!res.ok) throw new Error("Erreur lors de la suppression");
    try {
      const items = await fridgeService.getAll();
      localStorage.setItem(STORAGE_KEY, JSON.stringify(items));
    } catch (e) {
      console.warn("Impossible de rafraîchir le localStorage après suppression", e);
    }
  },

  async update(id: number, item: Partial<FoodItem>): Promise<FoodItem> {
    const res = await fetch(`${API_URL}/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item),
    });
    if (!res.ok) throw new Error("Erreur lors de la mise à jour");
    return res.json();
  },
};