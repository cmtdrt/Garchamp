import { useState, useEffect, useRef } from "react";
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
import { fridgeService } from "@/services/FridgeService";
import { FoodItem } from "@/types/FoodItem";

// Couleurs pour tous les allergènes
const allergenColors: Record<string, string> = {
  "Gluten": "bg-yellow-400",
  "Crustacés": "bg-blue-500",
  "Œufs": "bg-orange-400",
  "Poissons": "bg-cyan-500",
  "Arachides": "bg-amber-800",
  "Soja": "bg-lime-600",
  "Lait": "bg-stone-300",
  "Fruits à coque": "bg-yellow-900",
  "Céleri": "bg-green-500",
  "Moutarde": "bg-yellow-500",
  "Sésame": "bg-black",
  "Sulfites": "bg-gray-400",
  "Lupin": "bg-violet-500",
  "Mollusques": "bg-blue-900",
};

const STORAGE_KEY = "fridge_items";

const beep = () => {
  const audioCtx = new (window.AudioContext || (window as any).webkitAudioContext)();
  const oscillator = audioCtx.createOscillator();
  oscillator.type = "square";
  oscillator.frequency.setValueAtTime(1000, audioCtx.currentTime);
  oscillator.connect(audioCtx.destination);
  oscillator.start();
  oscillator.stop(audioCtx.currentTime + 0.2);
};

const Frigo = () => {
  const [foodItems, setFoodItems] = useState<FoodItem[]>([]);
  const [newItem, setNewItem] = useState({
    name: "",
    quantity: "",
    unity: "g",
    expiryDate: "",
  });
  const [selectedItem, setSelectedItem] = useState<FoodItem | null>(null);
  const [showMacrosDialog, setShowMacrosDialog] = useState(false);
  const [temperature, setTemperature] = useState(6); // Température initiale
  const criticalTemp = 6.5; // Seuil critique
  const intervalRef = useRef<ReturnType<typeof setInterval> | null>(null);
  const [loading, setLoading] = useState(true);

  // Charger depuis localStorage puis API
  useEffect(() => {
//    const localData = localStorage.getItem(STORAGE_KEY);
//    if (localData) setFoodItems(JSON.parse(localData));

    const fetchData = async () => {
      try {
        const items = await fridgeService.getAll();
        console.log("Réponse API :", items);
        setFoodItems(items);
        localStorage.setItem(STORAGE_KEY, JSON.stringify(items));
      } catch (err) {
        console.log("Impossible de charger l'api")
        toast.error("Impossible de charger les aliments depuis l’API 😢");
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, []);

  const updateLocalStorage = (items: FoodItem[]) => {
    setFoodItems(items);
    localStorage.setItem(STORAGE_KEY, JSON.stringify(items));
  };

  // ➕ Ajouter un aliment
  const handleAddItem = async () => {
    if (!newItem.name || !newItem.quantity) {
      toast.error("Veuillez remplir tous les champs obligatoires");
      return;
    }

    const item: FoodItem = {
      name: newItem.name,
      quantity: parseInt(newItem.quantity),
      unit: newItem.unity,
      expiration_date: newItem.expiryDate || null,
    };

    try {
      console.log(item)
    const res = await fridgeService.add(item);
    console.log(res);
    if (!res.ok || (res.status !== 200 && res.status !== 201)) throw new Error("Erreur lors de l’ajout");
    // Mise à jour du localStorage / state avec la réponse de l'API
    else {
      const updatedItems = [...foodItems, item];
      setFoodItems(updatedItems);
      localStorage.setItem(STORAGE_KEY, JSON.stringify(updatedItems));
      toast.success(`${item.name} ajouté au frigo !`);
    }
    console.log("Réponse API :", res);
  } catch {
    // Si API non dispo, on ajoute localement avec un ID temporaire
    const localItem = { ...item, id: Date.now().toString() };
    const updatedItems = [...foodItems, localItem];
    setFoodItems(updatedItems);
    localStorage.setItem(STORAGE_KEY, JSON.stringify(updatedItems));
    toast.warning("Ajouté localement (API non disponible)");
  }

    setNewItem({ name: "", quantity: "", unity: "g", expiryDate: "" });
  };

  // 🗑️ Supprimer un aliment
  const handleDeleteItem = async (id: string) => {
    const updated = foodItems.filter((item) => item.id !== id);
    updateLocalStorage(updated);

    try {
      await fridgeService.remove(id);
      toast.success("Aliment supprimé ✅");
    } catch {
      toast.warning("Suppression locale (API non disponible)");
    }
  };

  // 📊 Voir les macros
  const handleShowMacros = (item: FoodItem) => {
    setSelectedItem(item);
    setShowMacrosDialog(true);
  };

  // Simulation température
  useEffect(() => {
    intervalRef.current = setInterval(() => {
      setTemperature((prev) => {
        const next = prev + 0.01;
        if (next >= criticalTemp) beep();
        return next;
      });
    }, 1000);

    return () => {
      if (intervalRef.current) clearInterval(intervalRef.current);
    };
  }, []);

  return (
    <div className="min-h-screen bg-gradient-to-br from-background to-cream p-8">
      <div className="max-w-6xl mx-auto space-y-8 animate-fade-in">
        <div>
          <h1 className="text-4xl font-bold text-foreground mb-2">Mon Frigo 🧊</h1>
          <p className="text-muted-foreground">
            Gérez vos aliments et surveillez la température !
          </p>
          <p className={`mt-2 font-bold ${temperature >= criticalTemp ? "text-red-600" : "text-foreground"}`}>
            Température : {temperature.toFixed(1)}°C {temperature >= criticalTemp && "⚠️ Trop chaud !"}
          </p>
        </div>

        {/* Formulaire ajout */}
        <Card className="p-6 shadow-medium border-2 border-primary/20">
          <h2 className="text-xl font-semibold mb-4">Ajouter un aliment</h2>
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4">
            <Input
              placeholder="Nom"
              value={newItem.name}
              onChange={(e) => setNewItem({ ...newItem, name: e.target.value })}
            />
            <Input
              type="number"
              placeholder="Quantité"
              value={newItem.quantity}
              onChange={(e) => setNewItem({ ...newItem, quantity: e.target.value })}
            />
            <select
              value={newItem.unity}
              onChange={(e) => setNewItem({ ...newItem, unity: e.target.value })}
              className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm"
            >
              <option value="g">g</option>
              <option value="kg">kg</option>
              <option value="ml">ml</option>
              <option value="L">L</option>
              <option value="unité">unité</option>
            </select>
            <Input
              type="date"
              value={newItem.expiryDate}
              onChange={(e) => setNewItem({ ...newItem, expiryDate: e.target.value })}
            />
            <Button onClick={handleAddItem}>
              <Plus className="w-4 h-4 mr-2" /> Ajouter
            </Button>
          </div>
        </Card>

        {/* Liste aliments */}
        <div>
          <h2 className="text-2xl font-bold mb-4">
            Mes aliments ({foodItems.length})
          </h2>
          {loading ? (
            <Card className="p-12 text-center">Chargement...</Card>
          ) : foodItems.length === 0 ? (
            <Card className="p-12 text-center text-muted-foreground">
              Votre frigo est vide !
            </Card>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {foodItems.map((item) => (
                <Card key={item.id} className="p-4 hover:shadow-medium transition-all border-l-4 border-l-primary">
                  <div className="flex items-start justify-between mb-3">
                    <div className="flex-1">
                      <h3 className="font-semibold text-lg">{item.name}</h3>
                      <p className="text-sm text-muted-foreground">{item.quantity} {item.y}</p>
                      {item.expiration_date && (
                        <p className="text-xs text-destructive mt-1">
                          Expire le {new Date(item.expiration_date).toLocaleDateString()}
                        </p>
                      )}
                    </div>
                    <div className="flex gap-2">
                      <Button variant="ghost" size="icon" onClick={() => handleShowMacros(item)}>
                        <Info className="w-4 h-4 text-primary" />
                      </Button>
                      <Button variant="ghost" size="icon" onClick={() => handleDeleteItem(item.id)}>
                        <Trash2 className="w-4 h-4 text-destructive" />
                      </Button>
                    </div>
                  </div>
                  {item.allergens && item.allergens.length > 0 && (
                    <div className="flex flex-wrap gap-2">
                      {item.allergens.map((a) => (
                        <Badge key={a} className={`${allergenColors[a] || "bg-gray-400"} text-white text-xs`}>
                          {a}
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

      {/* Dialog Macros */}
      <Dialog open={showMacrosDialog} onOpenChange={setShowMacrosDialog}>
        <DialogContent className="max-w-md">
          <DialogHeader>
            <DialogTitle className="text-2xl text-primary">{selectedItem?.name}</DialogTitle>
            <DialogDescription>Valeurs nutritionnelles pour 100g</DialogDescription>
          </DialogHeader>
          {selectedItem?.macros && (
            <div className="space-y-3">
              {Object.entries(selectedItem.macros).map(([key, value]) => (
                <div key={key} className="flex justify-between items-center p-3 bg-orange-soft rounded-lg">
                  <span className="capitalize">{key.replace("_", " ")}</span>
                  <span className="font-bold text-primary">{value}</span>
                </div>
              ))}
            </div>
          )}
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default Frigo;
