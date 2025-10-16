import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Plus, Trash2, Info } from "lucide-react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { toast } from "sonner";

type FoodItem = {
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

const Frigo = () => {
  const [foodItems, setFoodItems] = useState<FoodItem[]>([]);
  const [newItem, setNewItem] = useState({
    name: "",
    quantity: "",
    unit: "g",
    expiryDate: "",
  });
  const [selectedItem, setSelectedItem] = useState<FoodItem | null>(null);
  const [showMacrosDialog, setShowMacrosDialog] = useState(false);

  const allergenColors: Record<string, string> = {
    gluten: "bg-allergen-gluten",
    lactose: "bg-allergen-dairy",
    noix: "bg-allergen-nuts",
    soja: "bg-allergen-soy",
    oeuf: "bg-allergen-egg",
  };

  const handleAddItem = () => {
    if (!newItem.name || !newItem.quantity) {
      toast.error("Veuillez remplir tous les champs obligatoires");
      return;
    }

    const item: FoodItem = {
      id: Date.now().toString(),
      name: newItem.name,
      quantity: parseFloat(newItem.quantity),
      unit: newItem.unit,
      expiryDate: newItem.expiryDate || undefined,
      allergens: ["gluten", "lactose"],
      macros: {
        energy_kcal: 250,
        protein_g: 5.2,
        fat_g: 3.1,
        carbohydrate_g: 45.5,
        fiber_g: 2.8,
        sugar_g: 8.5,
        salt_g: 0.5,
      },
    };

    setFoodItems([...foodItems, item]);
    setNewItem({ name: "", quantity: "", unit: "g", expiryDate: "" });
    toast.success(`${item.name} ajouté au frigo !`);
  };

  const handleDeleteItem = (id: string) => {
    const item = foodItems.find((i) => i.id === id);
    setFoodItems(foodItems.filter((item) => item.id !== id));
    toast.success(`${item?.name} retiré du frigo`);
  };

  const handleShowMacros = (item: FoodItem) => {
    setSelectedItem(item);
    setShowMacrosDialog(true);
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-background to-cream p-8">
      <div className="max-w-6xl mx-auto space-y-8 animate-fade-in">
        <div>
          <h1 className="text-4xl font-bold text-foreground mb-2">Mon Frigo 🧊</h1>
          <p className="text-muted-foreground">
            Gérez vos aliments et leurs dates d'expiration
          </p>
        </div>

        {/* Formulaire d'ajout */}
        <Card className="p-6 shadow-medium border-2 border-primary/20">
          <h2 className="text-xl font-semibold text-foreground mb-4">
            Ajouter un aliment
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4">
            <Input
              placeholder="Nom de l'aliment"
              value={newItem.name}
              onChange={(e) => setNewItem({ ...newItem, name: e.target.value })}
              className="border-input focus:border-primary"
            />
            <Input
              type="number"
              placeholder="Quantité"
              value={newItem.quantity}
              onChange={(e) =>
                setNewItem({ ...newItem, quantity: e.target.value })
              }
              className="border-input focus:border-primary"
            />
            <select
              value={newItem.unit}
              onChange={(e) => setNewItem({ ...newItem, unit: e.target.value })}
              className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus:border-primary focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
            >
              <option value="g">g</option>
              <option value="kg">kg</option>
              <option value="ml">ml</option>
              <option value="L">L</option>
              <option value="unité">unité</option>
            </select>
            <Input
              type="date"
              placeholder="Date d'expiration"
              value={newItem.expiryDate}
              onChange={(e) =>
                setNewItem({ ...newItem, expiryDate: e.target.value })
              }
              className="border-input focus:border-primary"
            />
            <Button
              onClick={handleAddItem}
              className="bg-gradient-to-r from-primary to-secondary hover:opacity-90 transition-opacity"
            >
              <Plus className="w-4 h-4 mr-2" />
              Ajouter
            </Button>
          </div>
        </Card>

        {/* Liste des aliments */}
        <div>
          <h2 className="text-2xl font-bold text-foreground mb-4">
            Mes aliments ({foodItems.length})
          </h2>
          {foodItems.length === 0 ? (
            <Card className="p-12 text-center">
              <p className="text-muted-foreground text-lg">
                Votre frigo est vide ! Ajoutez des aliments ci-dessus.
              </p>
            </Card>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {foodItems.map((item) => (
                <Card
                  key={item.id}
                  className="p-4 hover:shadow-medium transition-all duration-300 border-l-4 border-l-primary animate-slide-in"
                >
                  <div className="flex items-start justify-between mb-3">
                    <div className="flex-1">
                      <h3 className="font-semibold text-lg text-foreground">
                        {item.name}
                      </h3>
                      <p className="text-sm text-muted-foreground">
                        {item.quantity} {item.unit}
                      </p>
                      {item.expiryDate && (
                        <p className="text-xs text-destructive mt-1">
                          Expire le {new Date(item.expiryDate).toLocaleDateString()}
                        </p>
                      )}
                    </div>
                    <div className="flex gap-2">
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => handleShowMacros(item)}
                        className="h-8 w-8 hover:bg-primary/10"
                      >
                        <Info className="w-4 h-4 text-primary" />
                      </Button>
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() => handleDeleteItem(item.id)}
                        className="h-8 w-8 hover:bg-destructive/10"
                      >
                        <Trash2 className="w-4 h-4 text-destructive" />
                      </Button>
                    </div>
                  </div>
                  {item.allergens && item.allergens.length > 0 && (
                    <div className="flex flex-wrap gap-2">
                      {item.allergens.map((allergen) => (
                        <Badge
                          key={allergen}
                          className={`${
                            allergenColors[allergen] || "bg-muted"
                          } text-white text-xs`}
                        >
                          {allergen}
                        </Badge>
                      ))}
                    </div>
                  )}
                </Card>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Dialog Macronutriments */}
      <Dialog open={showMacrosDialog} onOpenChange={setShowMacrosDialog}>
        <DialogContent className="max-w-md">
          <DialogHeader>
            <DialogTitle className="text-2xl text-primary">
              {selectedItem?.name}
            </DialogTitle>
            <DialogDescription>
              Valeurs nutritionnelles pour 100g
            </DialogDescription>
          </DialogHeader>
          {selectedItem?.macros && (
            <div className="space-y-3">
              {Object.entries(selectedItem.macros).map(([key, value]) => {
                const labels: Record<string, string> = {
                  energy_kcal: "Énergie (kcal)",
                  protein_g: "Protéines (g)",
                  fat_g: "Lipides (g)",
                  carbohydrate_g: "Glucides (g)",
                  fiber_g: "Fibres (g)",
                  sugar_g: "Sucres (g)",
                  salt_g: "Sel (g)",
                };
                return (
                  <div
                    key={key}
                    className="flex justify-between items-center p-3 bg-orange-soft rounded-lg"
                  >
                    <span className="font-medium text-foreground">
                      {labels[key]}
                    </span>
                    <span className="font-bold text-primary">{value}</span>
                  </div>
                );
              })}
            </div>
          )}
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default Frigo;
