import React from "react";
import type { RecipeResponse } from "@/services/recipe.service";
import { Clock, GraduationCap } from "lucide-react";

type Props = {
  response: RecipeResponse;
};

const ChatResponse: React.FC<Props> = ({ response }) => {
  if (response.error) {
    const errorValue = response.error;
    return (
      <div className="text-red-600 font-medium">
        Erreur: {errorValue}
      </div>
    );
  }

  const { data } = response;
  if (!data) return null;

  const steps: string[] = data.steps;

  const difficultyColor = (() => {
    const level = (data.difficulty || "").toLowerCase();
    if (level.includes("facile")) return "text-green-600";
    if (level.includes("moyen")) return "text-orange-500";
    if (level.includes("difficile")) return "text-red-600";
    return "text-gray-400";
  })();

  return (
    <div className="space-y-4">
      <div className="flex items-center gap-1">
        <div className={`whitespace-nowrap inline-flex items-center ${difficultyColor}`} title={data.difficulty}>
          <GraduationCap className="w-5 h-5 sm:w-6 sm:h-6" />
        </div>
        <div className="flex-1 text-center font-semibold text-base sm:text-lg px-1 truncate">{data.title}</div>
        <div className="text-right text-xs sm:text-sm text-muted-foreground whitespace-nowrap inline-flex items-center gap-1">
          <Clock className="w-5 h-5 sm:w-6 sm:h-6" />
          {data.estimated_time}
        </div>
      </div>

      <div className="text-sm text-foreground/90">{data.description}</div>

      <div>
        <div className="font-medium mb-1">Ingrédients</div>
        <ul className="list-disc pl-5 space-y-1 text-sm">
          {data.ingredients.map((ing, idx) => (
            <li key={idx}>{`${ing.quantity} de ${ing.name}`}</li>
          ))}
        </ul>
      </div>

      <div>
        <div className="font-medium mb-1">Étapes</div>
        <div className="whitespace-pre-line text-sm leading-relaxed space-y-3">
          {steps.map((s, i) => (
            <div key={i}>{s}</div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default ChatResponse;


