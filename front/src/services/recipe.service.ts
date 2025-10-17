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
  data: RecipeData;
  error?: string;
};

export async function askRecipe(payload: RecipeRequest): Promise<RecipeResponse> {
  console.log("askRecipe", payload);
  // Simulation: délai 2s puis renvoi d'une chaîne
  await new Promise((resolve) => setTimeout(resolve, 2000));
  return {
    status: "success",
    message: "item created successfully",
    data: {
      title: "Pâte à la crème fraîche et lardons fumés",
      description:
        "Recette simple de pâte à la crème fraîche et aux lardons fumés, parfaite pour une ambiance conviviale.",
      ingredients: [
        { name: "Pâtes", quantity: "300 g" },
        { name: "Lardons fumés", quantity: "150 g" },
        { name: "Oignon", quantity: "1 pièce" },
        { name: "Persil frais", quantity: "5 g" },
      ],
      steps: [
        "Faire bouillir l'eau salée dans une grande casserole. Ajoutez les lardons fumés et cuisez-les pendant environ 3 minutes.",
        "Maisonnez l'oignon émincé et le faites revenir dans une poêle à feu doux jusqu'à ce qu'il soit tendre. Ajoutez la crème fraîche liquide.",
        "Rajoutez les lardons fumés cuits dans la poêle et mélangez bien pour incorporer le goût du lardons à la crème fraîche.",
        "En même temps, faites cuire les pâtes selon les indications sur le paquet. Une fois cuites, versez-les dans la poêle avec la sauce et mélangez bien.",
        "Terminez en hachant le persil frais finement et en l'ajoutant à la préparation.",
      ],
      estimated_time: "30 minutes",
      difficulty: "facile",
    },
  };
}