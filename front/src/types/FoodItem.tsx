export type FoodItem = {
  id: string;
  name: string;
  quantity: number;
  unit: string;
  expiryDate?: string;
  allergens?: string[];
  macros?: {
    energy_kcal: number;
    protein_g: number;
    fat_g: number;
    carbohydrate_g: number;
    fiber_g: number;
    sugar_g: number;
    salt_g: number;
  };
};
