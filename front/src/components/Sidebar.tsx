import { Link, useLocation } from "react-router-dom";
import { Refrigerator, ChefHat } from "lucide-react";
import { cn } from "@/lib/utils";

const Sidebar = () => {
  const location = useLocation();

  const menuItems = [
    { path: "/", icon: Refrigerator, label: "Mon Frigo" },
    { path: "/recette", icon: ChefHat, label: "Créer une Recette" },
  ];

  return (
    <aside className="w-64 h-screen bg-card border-r border-border sticky top-0 left-0 flex flex-col">
      <div className="p-6 border-b border-border">
        <div className="flex items-center justify-between">
          <h1 className="text-3xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
            Garchamp
          </h1>
          <img
            src="/logo.png"
            alt="Garchamp logo"
            className="h-12 w-12 object-contain drop-shadow -mt-3"
          />
        </div>
        <p className="text-sm text-muted-foreground mt-1">Votre assistant cuisine</p>
      </div>

      <nav className="flex-1 p-4 space-y-2">
        {menuItems.map((item) => {
          const Icon = item.icon;
          const isActive = location.pathname === item.path;

          return (
            <Link
              key={item.path}
              to={item.path}
              className={cn(
                "flex items-center gap-3 px-4 py-3 rounded-xl transition-all duration-300",
                "hover:bg-sidebar-accent group",
                isActive
                  ? "bg-gradient-to-r from-primary to-secondary text-primary-foreground shadow-md"
                  : "text-sidebar-foreground"
              )}
            >
              <Icon className={cn(
                "w-5 h-5 transition-transform duration-300",
                isActive ? "scale-110" : "group-hover:scale-110"
              )} />
              <span className="font-medium">{item.label}</span>
            </Link>
          );
        })}
      </nav>

      <div className="p-4 border-t border-border">
        <div className="bg-gradient-to-br from-orange-soft to-cream rounded-xl p-4 text-center">
        <p className="text-sm font-medium text-foreground">
        🇫🇷 Ma France<br />🍷 Mon pinard<br />🥦  mes recettes
         </p>
        </div>
      </div>
    </aside>
  );
};

export default Sidebar;
