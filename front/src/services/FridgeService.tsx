// src/services/frigoService.ts
import { FoodItem } from "@/types/FoodItem";

const STORAGE_KEY = "foodItems";

// 🔹 Récupérer les aliments depuis le localStorage
export const getFoodItems = (): FoodItem[] => {
  const data = localStorage.getItem(STORAGE_KEY);
  return data ? JSON.parse(data) : [];
};

// 🔹 Sauvegarder les aliments dans le localStorage
export const saveFoodItems = (items: FoodItem[]) => {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(items));
};

// 🔹 Ajouter un aliment
export const addFoodItem = (newItem: FoodItem): FoodItem[] => {
  const items = getFoodItems();
  const updated = [...items, newItem];
  saveFoodItems(updated);
  return updated;
};

// 🔹 Supprimer un aliment
export const deleteFoodItem = (id: string): FoodItem[] => {
  const items = getFoodItems().filter((item) => item.id !== id);
  saveFoodItems(items);
  return items;
};
