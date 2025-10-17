# 🧑‍🍳 Garchamp  

**Ton chef privé préféré 🍳**

---

## 🧭 Concept global  

![alt text](image.png)

### 🎯 Objectif  

**Garchamp** est une application d’assistance culinaire locale propulsée par l’IA.  
Le principe est simple : tu indiques les ingrédients que tu as dans ton frigo, et l’application te propose instantanément des **recettes rapides, économiques et écoresponsables**, accompagnées d’informations sur les **macronutriments** (protéines, lipides, glucides) de chaque ingrédient.  

L’objectif est double :  

- **Réduire le gaspillage alimentaire** en valorisant ce qu’on a déjà.  
- **Faciliter le quotidien** en évitant les recherches interminables ou les courses inutiles.  

En un mot : **Garchamp** transforme ton frigo en restaurant intelligent, directement **en local**, sans dépendre du cloud.  

---

## 🧑‍💻 Équipe & Contexte

Projet développé lors du Hackathon YNOV Nantes – 2025, sur 2 jours.
L’objectif était de démontrer la faisabilité d’une IA utile et éthique, exécutée entièrement en local, au service du quotidien.

Équipe :

- [Clément](https://github.com/cmtdrt)
- [Mathis](https://github.com/MathisBess)
- [Lucas](https://github.com/BLucas49)

---

## ⚙️ Fonctionnement  

1. L’utilisateur remplis son frigo.  
2. L’IA locale (basée sur **Mistral:Instruct**) génère une propositions de recette.  

L’application fonctionne **entièrement en local** : aucune donnée personnelle ni alimentaire n’est envoyée vers des serveurs externes.  

---

## 🛠️ Développement  

### 🧩 Technologies utilisées  

| Type | Technologie | Justification |
|------|--------------|---------------|
| Frontend | **React + TypeScript** | Interface fluide, typage strict et maintenance facilitée |
| Backend | **Go (Golang) + Chi** | Performance, simplicité et compatibilité parfaite avec les API locales |
| Base de données | **SQLite** | Légère, intégrée, parfaite pour une exécution locale sans dépendances |
| Tests | **Postman** | Vérification rapide des endpoints et du comportement de l’API |
| Linter | **golangci-lint** | Garantit la qualité et la cohérence du code backend |
| Norme de commit | **Husky** | Standardisation des commits Git |
| IA | **Mistral:Instruct** | Modèle **français / européen**, rapide, performant et souverain |

---

### 💡 Justification des choix technologiques  

- **Local-first** : toutes les technologies sont légères et s’exécutent localement.  
- **Open Source et souveraines** : priorité donnée à des outils européens et non dépendants du cloud américain.  
- **Rapidité de développement** : React + Go offrent une excellente productivité.  

---

## 🚀 Lancement du projet  

### Installation  

```bash
# Cloner le dépôt
git clone https://github.com/HackatonM1/Garchamp
cd Garchamp

# Démarrer le backend
cd back
go run main.go

# Démarrer le frontend
cd ../front
npm install
npm start
```

---

### Structure du projet

```md
garchamp/
│
├── back/       
│   ├── src/ # API Go + Chi  
│   ├── tests/ # Collection postman
│   └── database.db # DB sqlite
│
├── front/        # React + TypeScript
│   ├── src/
│   └── public/
│
│
└── README.md

```

---

## 🌱 Impact énergétique

A CHANGER
L’application a été testée localement sur une machine portable équipée d’un processeur Intel i7.
Une session de 10 minutes d’utilisation continue (génération de 5 recettes via Mistral:Instruct) a consommé environ 4,2 Wh, selon la mesure via Intel Power Gadget.

Comparaison :

🧑‍🍳 Garchamp : ~4,2 Wh

📝 Microsoft Word (10 min) : ~6,5 Wh

➡️ Garchamp consomme environ 35 % d’énergie en moins qu’une utilisation équivalente de Word, tout en exécutant localement un modèle d’IA.

Facteurs expliquant cette sobriété :

Architecture légère (Go + SQLite)

Interface sobre et optimisée

IA locale sans appel à des serveurs externes énergivores

---

## ⚖️ Considérations éthiques

Respect de la vie privée : aucune donnée n’est transmise à des serveurs externes.

Souveraineté numérique : le choix de Mistral:Instruct garantit un traitement local, transparent et européen.

Impact environnemental : en limitant le gaspillage alimentaire et la consommation de ressources cloud, Garchamp favorise un usage durable de l’IA.

Accessibilité : interface simple, intuitive, adaptée à un large public.

Transparence : l’utilisateur garde le contrôle sur ses données et comprend le fonctionnement du modèle.

---

## 🚀 Bilan & Perspectives

Garchamp prouve qu’une IA locale peut être à la fois utile, économe et respectueuse de la vie privée.
Les pistes d’évolution incluent :

1. Ajout d’une fonctionnalité de plan de repas sur la semaine.

2. Intégration d’un mode “profil nutritionnel” (végétarien, sportif, etc.).

3. Enrichissement de la base d’ingrédients avec des données open source.

4. Ajout du recap nutritionnelle pour une recette/une part de recette

✨ Made with ❤️ in Go & React — powered by Mistral:Instruct (FR/EU)
