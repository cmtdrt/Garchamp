// Service pour les recettes

export type Item = {
  id: number;
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

export type RecipeIngredient = {
  name: string;
  quantity: string;
};

export type RecipeData = {
  title: string;
  description: string;
  ingredients: RecipeIngredient[];
  steps: string[];
  estimated_time: string;
  difficulty: string;
};

export type RecipeResponse = {
  status: string;
  message: string;
  data?: RecipeData;
  error?: string;
};

export async function askRecipe(payload: RecipeRequest): Promise<RecipeResponse> {
  try {
    const response = await fetch("http://localhost:8080/api/v1/recipe/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    const json = await response.json();

    if (!response.ok) {
      return {
        status: "error",
        message: json?.message ?? "Erreur lors de l'appel API",
        error: json?.error ?? response.statusText,
      };
    }

    return json as RecipeResponse;
  } catch (e) {
    const err = e as Error;
    return {
      status: "error",
      message: err.message ?? "Erreur réseau",
      error: err.message ?? "Erreur réseau",
    };
  }
}