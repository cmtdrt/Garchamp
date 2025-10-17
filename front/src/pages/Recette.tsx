import { useEffect, useRef, useState } from "react";
import { Input } from "@/components/ui/input";
import { Card } from "@/components/ui/card";
import { Users, Check } from "lucide-react";
import { askRecipe, type Item, type RecipeResponse } from "@/services/recipe.service";
import ChatResponse from "@/components/ChatResponse";
import { toast } from "sonner";

type Message = {
  id: string;
  role: "user" | "assistant";
  content: string;
  recipeResponse?: RecipeResponse;
};

const Recette = () => {
  const [messages, setMessages] = useState<Message[]>([
    {
      id: "1",
      role: "assistant",
      content:
        "Bonjour ! Je suis votre assistant culinaire. Dites-moi pour combien de personnes vous souhaitez cuisiner et je vous proposerai des recettes avec les aliments de votre frigo ! 👨‍🍳",
    },
  ]);
  const [input, setInput] = useState("");
  const [numberOfPeople, setNumberOfPeople] = useState("2");
  const [isLoading, setIsLoading] = useState(false);
  const [isTyping, setIsTyping] = useState(false);
  const [mode, setMode] = useState<"frigo" | "libre">("frigo");
  const scrollRef = useRef<HTMLDivElement>(null);
  const [showOptions, setShowOptions] = useState(true);
  const [isAllergenOpen, setIsAllergenOpen] = useState(false);
  const [selectedAllergens, setSelectedAllergens] = useState<string[]>([]);

  const ALLERGENS: { key: string; label: string; bg: string; border: string }[] = [
    { key: "Gluten", label: "Gluten", bg: "bg-yellow-400", border: "border-yellow-400" },
    { key: "Crustacés", label: "Crustacés", bg: "bg-blue-500", border: "border-blue-500" },
    { key: "Œufs", label: "Œufs", bg: "bg-orange-400", border: "border-orange-400" },
    { key: "Poissons", label: "Poissons", bg: "bg-cyan-500", border: "border-cyan-500" },
    { key: "Arachides", label: "Arachides", bg: "bg-amber-800", border: "border-amber-800" },
    { key: "Soja", label: "Soja", bg: "bg-lime-600", border: "border-lime-600" },
    { key: "Lait", label: "Lait", bg: "bg-stone-300", border: "border-stone-300" },
    { key: "Fruits à coque", label: "Fruits à coque", bg: "bg-yellow-900", border: "border-yellow-900" },
    { key: "Céleri", label: "Céleri", bg: "bg-green-500", border: "border-green-500" },
    { key: "Moutarde", label: "Moutarde", bg: "bg-yellow-500", border: "border-yellow-500" },
    { key: "Sésame", label: "Sésame", bg: "bg-black", border: "border-black" },
    { key: "Sulfites", label: "Sulfites", bg: "bg-gray-400", border: "border-gray-400" },
    { key: "Lupin", label: "Lupin", bg: "bg-violet-500", border: "border-violet-500" },
    { key: "Mollusques", label: "Mollusques", bg: "bg-blue-900", border: "border-blue-900" },
  ];

  const toggleAllergen = (key: string) => {
    setSelectedAllergens((prev) =>
      prev.includes(key) ? prev.filter((k) => k !== key) : [...prev, key]
    );
  };

  type StoredFoodItem = {
    id: string;
    name: string;
    quantity: number | string;
    unit: string;
    expiryDate?: string;
  };

  const handleSendMessage = async (selectedMode?: "frigo" | "libre") => {
    const effectiveMode = selectedMode ?? mode;
    const predefinedFrigo = "Je veux une recette à partir des aliments que j'ai dans mon frigo";
    const predefinedLibre = "Je veux une recette et tu es libre des ingrédients !";
    const predefined = effectiveMode === "frigo" ? predefinedFrigo : predefinedLibre;

    const userMessage: Message = {
      id: Date.now().toString(),
      role: "user",
      content: predefined,
    };

    setMessages([...messages, userMessage]);

    try {
      setIsLoading(true);
      setIsTyping(true);
      setShowOptions(false);

      // Récupère les items du frigo depuis localStorage (stockage simple pour relier les pages)
      let items: Item[] = [];
      if (effectiveMode === "frigo") {
        const stored = localStorage.getItem("frigo_items");
        if (stored) {
          const parsed = JSON.parse(stored) as StoredFoodItem[];
          items = parsed.map((i) => ({
            id: Number(i.id),
            name: i.name,
            quantity: typeof i.quantity === "string" ? Number(i.quantity) || 0 : i.quantity,
            unit: i.unit || "unité",
            exp_date: i.expiryDate,
          }));
        }
      }

      const payload = {
        items,
        people_number: Number(numberOfPeople) || 1,
        allergens: selectedAllergens.map((a) => a.toLowerCase()),
      };

      const res = await askRecipe(payload);
      const aiText = ""; // rendu via composant structuré par message

      const aiMessage: Message = {
        id: (Date.now() + 1).toString(),
        role: "assistant",
        content: aiText,
        recipeResponse: res,
      };
      const bonAppetitMessage: Message = {
        id: (Date.now() + 2).toString(),
        role: "assistant",
        content: "Bon appétit !",
      };
      setMessages((prev) => [...prev, aiMessage, bonAppetitMessage]);
    } catch (e: unknown) {
      const message = e instanceof Error ? e.message : "Impossible de récupérer la recette";
      toast.error(message);
    } finally {
      setIsLoading(false);
      setIsTyping(false);
      setTimeout(() => setShowOptions(true), 1000);
    }
  };

  const handleKeyPress = (_e: React.KeyboardEvent) => {
    // plus d'envoi via entrée
  };

  // Auto-scroll vers le bas à chaque nouveau message ou changement d'état
  useEffect(() => {
    const el = scrollRef.current;
    if (!el) return;
    el.scrollTo({ top: el.scrollHeight, behavior: "smooth" });
  }, [messages, isTyping]);

  // Assure d'être tout en bas au premier rendu et quand les options réapparaissent
  useEffect(() => {
    const el = scrollRef.current;
    if (!el) return;
    // petite attente pour laisser le DOM peindre
    const id = window.setTimeout(() => {
      el.scrollTo({ top: el.scrollHeight, behavior: "auto" });
    }, 0);
    return () => window.clearTimeout(id);
  }, []);

  useEffect(() => {
    const el = scrollRef.current;
    if (!el) return;
    el.scrollTo({ top: el.scrollHeight, behavior: "smooth" });
  }, [showOptions, isLoading]);

  return (
    <div className="min-h-screen bg-gradient-to-br from-background to-cream p-8">
      <div className="max-w-4xl mx-auto h-[calc(100vh-2rem)] flex flex-col animate-fade-in">
        <div className="mb-3">
          <h1 className="text-4xl font-bold text-foreground mb-2">
            Créer une Recette 👨‍🍳
          </h1>
          {/* sous-titre supprimé pour gagner en hauteur */}
          {/* Dropdown Allergènes */}
          <div className="mt-4">
            <button
              type="button"
              onClick={() => setIsAllergenOpen((v) => !v)}
              className="w-full flex items-center justify-between px-4 py-2 rounded-md border-2 border-primary/20 hover:border-primary/40 transition-colors bg-card"
            >
              <span className="font-medium">Allergènes</span>
              <span className={`transition-transform ${isAllergenOpen ? "rotate-180" : "rotate-0"}`}>
                ▼
              </span>
            </button>
            {isAllergenOpen && (
              <div className="p-3 border-2 border-t-0 border-primary/20 rounded-b-md grid grid-cols-1 sm:grid-cols-2 md:grid-cols-7 gap-2 bg-card">
                {ALLERGENS.map(({ key, label, bg, border }) => {
                  const selected = selectedAllergens.includes(key);
                  return (
                    <button
                      key={key}
                      type="button"
                      onClick={() => toggleAllergen(key)}
                      className={`inline-flex items-center gap-2 px-2 py-1 rounded-full border-2 text-xs transition-colors ${
                        selected
                          ? `${bg} text-white border-transparent`
                          : `bg-white text-foreground ${border}`
                      }`}
                    >
                      {selected && (
                        <Check className="w-3 h-3 text-white" />
                      )}
                      <span className="text-sm font-medium leading-none">{label}</span>
                    </button>
                  );
                })}
              </div>
            )}
          </div>
        </div>

        {/* Nombre de personnes */}
        <Card className="p-2 mb-3 shadow-medium border-2 border-primary/20">
          <div className="flex items-center gap-2">
            <Users className="w-4 h-4 text-primary" />
            <label className="font-medium text-sm text-foreground">
              Pour combien de personnes ?
            </label>
            <Input
              type="number"
              min="1"
              max="20"
              value={numberOfPeople}
              onChange={(e) => setNumberOfPeople(e.target.value)}
              className="w-16 border-input focus:border-primary text-sm py-1"
            />
            <span className="text-muted-foreground text-sm">
              {numberOfPeople === "1" ? "personne" : "personnes"}
            </span>
          </div>
        </Card>

        {/* Zone de chat */}
        <Card className="flex-1 flex flex-col shadow-medium border-2 border-primary/20 overflow-hidden">
          {/* Messages */}
          <div ref={scrollRef} className="flex-1 overflow-y-auto p-6 space-y-4">
            {messages.map((message) => (
              <div
                key={message.id}
                className={`flex ${
                  message.role === "user" ? "justify-end" : "justify-start"
                } animate-slide-in`}
              >
                <div
                  className={`max-w-[80%] rounded-2xl px-4 py-3 ${
                    message.role === "user"
                      ? "bg-gradient-to-r from-primary to-secondary text-primary-foreground"
                      : "bg-orange-soft text-foreground border-2 border-primary/20"
                  }`}
                >
                  {message.role === "assistant" && message.recipeResponse ? (
                    <ChatResponse response={message.recipeResponse} />
                  ) : (
                    <p className="text-sm leading-relaxed">{message.content}</p>
                  )}
                </div>
              </div>
            ))}

            {/* Bulles d'options cliquables (comme des messages utilisateur) */}
            {showOptions && !isLoading && !isTyping && (
              <div className="flex flex-col gap-3 mt-4">
                <button
                  type="button"
                  onClick={() => void handleSendMessage("frigo")}
                  className="self-end max-w-[80%] rounded-2xl px-4 py-3 bg-primary/15 text-foreground border-2 border-primary/30 shadow hover:bg-primary/15 hover:border-primary/50 transition-colors animate-slide-in"
                >
                  Je veux une recette à partir des aliments que j'ai dans mon frigo
                </button>
                <div className="self-center text-xs text-muted-foreground select-none">ou</div>
                <button
                  type="button"
                  onClick={() => void handleSendMessage("libre")}
                  className="self-end max-w-[80%] rounded-2xl px-4 py-3 bg-primary/15 text-foreground border-2 border-primary/30 shadow hover:bg-primary/15 hover:border-primary/50 transition-colors animate-slide-in"
                >
                  Je veux une recette et tu es libre des ingrédients !
                </button>
              </div>
            )}

            {isTyping && (
              <div className="flex justify-start">
                <div className="bg-orange-soft text-foreground border-2 border-primary/20 rounded-2xl px-4 py-3 inline-flex items-center gap-1">
                  <span className="inline-block w-2 h-2 rounded-full bg-foreground opacity-70 animate-bounce" />
                  <span className="inline-block w-2 h-2 rounded-full bg-foreground opacity-70 animate-bounce" style={{ animationDelay: "150ms" }} />
                  <span className="inline-block w-2 h-2 rounded-full bg-foreground opacity-70 animate-bounce" style={{ animationDelay: "300ms" }} />
                </div>
              </div>
            )}
          </div>

        </Card>
      </div>
    </div>
  );
};

export default Recette;
