// Service pour les recettes

export type Item = {
  id: string;
  name: string;
  quantity: number;
  unit: string;
  exp_date?: string;
};

export type RecipeRequest = {
  items: Item[];
  people_number: number;
  allergens: string[];
};

export type RecipeResponse = {
  reponse: string;
};

export async function askRecipe(payload: RecipeRequest): Promise<RecipeResponse> {
  console.log("askRecipe", payload);
  // Simulation: délai 2s puis renvoi d'une chaîne
  await new Promise((resolve) => setTimeout(resolve, 2000));
  return { reponse: "Voici la recette: Lorep ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua." };
}