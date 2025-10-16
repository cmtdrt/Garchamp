import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card } from "@/components/ui/card";
import { Send, Users } from "lucide-react";
import { toast } from "sonner";

type Message = {
  id: string;
  role: "user" | "assistant";
  content: string;
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

  const handleSendMessage = () => {
    if (!input.trim()) {
      toast.error("Veuillez écrire un message");
      return;
    }

    const userMessage: Message = {
      id: Date.now().toString(),
      role: "user",
      content: input,
    };

    setMessages([...messages, userMessage]);
    setInput("");

    // Simulation de réponse IA (à remplacer par vraie IA locale)
    setTimeout(() => {
      const aiMessage: Message = {
        id: (Date.now() + 1).toString(),
        role: "assistant",
        content: `Je vais créer une délicieuse recette pour ${numberOfPeople} personnes avec les aliments de votre frigo ! (IA à connecter)`,
      };
      setMessages((prev) => [...prev, aiMessage]);
    }, 1000);
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      handleSendMessage();
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-background to-cream p-8">
      <div className="max-w-4xl mx-auto h-[calc(100vh-4rem)] flex flex-col animate-fade-in">
        <div className="mb-6">
          <h1 className="text-4xl font-bold text-foreground mb-2">
            Créer une Recette 👨‍🍳
          </h1>
          <p className="text-muted-foreground">
            Discutez avec l'IA pour créer des recettes personnalisées
          </p>
        </div>

        {/* Nombre de personnes */}
        <Card className="p-4 mb-4 shadow-medium border-2 border-primary/20">
          <div className="flex items-center gap-4">
            <Users className="w-5 h-5 text-primary" />
            <label className="font-medium text-foreground">
              Pour combien de personnes ?
            </label>
            <Input
              type="number"
              min="1"
              max="20"
              value={numberOfPeople}
              onChange={(e) => setNumberOfPeople(e.target.value)}
              className="w-20 border-input focus:border-primary"
            />
            <span className="text-muted-foreground">
              {numberOfPeople === "1" ? "personne" : "personnes"}
            </span>
          </div>
        </Card>

        {/* Zone de chat */}
        <Card className="flex-1 flex flex-col shadow-medium border-2 border-primary/20 overflow-hidden">
          {/* Messages */}
          <div className="flex-1 overflow-y-auto p-6 space-y-4">
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
                  <p className="text-sm leading-relaxed">{message.content}</p>
                </div>
              </div>
            ))}
          </div>

          {/* Zone de saisie */}
          <div className="border-t border-border p-4 bg-card">
            <div className="flex gap-2">
              <Input
                placeholder="Demandez une recette avec vos ingrédients..."
                value={input}
                onChange={(e) => setInput(e.target.value)}
                onKeyPress={handleKeyPress}
                className="flex-1 border-input focus:border-primary"
              />
              <Button
                onClick={handleSendMessage}
                className="bg-gradient-to-r from-primary to-secondary hover:opacity-90 transition-opacity"
              >
                <Send className="w-4 h-4" />
              </Button>
            </div>
            <p className="text-xs text-muted-foreground mt-2">
              Appuyez sur Entrée pour envoyer
            </p>
          </div>
        </Card>
      </div>
    </div>
  );
};

export default Recette;
