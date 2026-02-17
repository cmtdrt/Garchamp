# 🧑‍🍳 Garchamp  

**Your favorite private chef 🍳**

---

## 🧭 Overall Concept  

### 🎯 Objective  

**Garchamp** is a local AI-powered cooking assistant application.  
The principle is simple: you enter the ingredients you have in your fridge, and the app instantly suggests **quick, budget-friendly and eco-friendly recipes**, along with **macronutrient** information (proteins, fats, carbohydrates) for each ingredient.  

The goals are threefold:  

- **Reduce food waste** by making the most of what you already have.  
- **Simplify everyday life** by avoiding endless searches or unnecessary grocery trips.  
- **Raise awareness of AI consumption and its environmental impact** — you can monitor your PC’s resource usage in real time (CPU load, power draw, memory) to see the cost of running the local model.  

In short: **Garchamp** turns your fridge into a smart restaurant, entirely **locally**, with no cloud dependency.  

---

## 🧑‍💻 Team & Context

Project developed during the YNOV Nantes Hackathon – 2025, over 2 days.
The aim was to demonstrate the feasibility of a useful and ethical AI, running entirely on-device, serving everyday needs.

Team:

- [Clément](https://github.com/cmtdrt)
- [Mathis](https://github.com/MathisBess)
- [Lucas](https://github.com/BLucas49)

---

## ⚙️ How It Works  

1. The user fills in their fridge.  
2. The local AI (based on **Mistral:Instruct**) generates recipe suggestions.  

The application runs **entirely locally**: no personal or food data is sent to external servers.  

---

## 🛠️ Development  

### 🧩 Technologies Used  

| Type | Technology | Justification |
|------|--------------|---------------|
| Frontend | **React + TypeScript** | Smooth UI, strict typing and easier maintenance |
| Backend | **Go (Golang) + Chi** | Performance, simplicity and perfect compatibility with local APIs |
| Database | **SQLite** | Lightweight, embedded, ideal for local execution with no external dependencies |
| Testing | **Postman** | Quick verification of endpoints and API behavior |
| Linter | **golangci-lint** | Ensures backend code quality and consistency |
| Commit standard | **Husky** | Standardizes Git commits |
| AI | **Mistral:Instruct** | **French / European** model, fast, performant and sovereign |

---

### 💡 Rationale for Technology Choices  

- **Local-first**: all technologies are lightweight and run locally.  
- **Open source and sovereign**: priority given to European tools with no dependency on US cloud.  
- **Rapid development**: React + Go offer excellent productivity.  

---

## 🚀 Getting Started  

### Installation  

```bash
# Clone the repository
git clone https://github.com/HackatonM1/Garchamp
cd Garchamp

# Start the backend
cd back
go run main.go

# Start the frontend
cd ../front
npm install
npm start

# Remember to configure the .env files
```

---

### Project Structure

```md
garchamp/
│
├── back/       
│   ├── src/ # Go + Chi API  
│   ├── tests/ # Postman collection
│   └── database.db # SQLite DB
│
├── front/        # React + TypeScript
│   ├── src/
│   └── public/
│
│
└── README.md

```

---

## 🌱 Energy Impact

This application combines a lightweight React/TypeScript frontend with a Go backend and a self-hosted AI model (Mistral:Instruct). Using a local language model implies significant CPU/GPU usage.

| Activity                              | Average CPU/GPU load                      | Average power consumption | Comparison                                          |
| ------------------------------------- | ------------------------------------------- | ------------------------------- | ---------------------------------------------------- |
| Garchamp (1 h) | CPU ~50–70 % | ~60–90 Wh                       | Equivalent to ~8–12 h of work in Microsoft Word   |
| React + Go (1 h, no AI)             | CPU ~10–15 %                                | ~10–15 Wh                       | Comparable to 1–2 h of Word                           |
| Microsoft Word (1 h)                  | CPU ~5 %                                    | ~5–7 Wh                         | -                             |
| HD video on YouTube (1 h)            | CPU ~30 %                                   | ~20–25 Wh                       | - |

Notes:

Values are averages on a typical laptop (Intel i7, 16 GB RAM).

---

## ⚖️ Ethical Considerations

**Privacy**: no data is sent to external servers.

**Digital sovereignty**: choosing Mistral:Instruct ensures local, transparent and European processing.

**Environmental impact**: by limiting food waste and cloud resource consumption, Garchamp promotes sustainable use of AI.

**Accessibility**: simple, intuitive interface, suited to a wide audience.

**Transparency**: users keep control of their data and understand how the model works.

---

## 🚀 Summary & Roadmap

Garchamp proves that local AI can be useful, efficient and privacy-respecting.
Planned evolutions include:

1. Adding a weekly meal plan feature.

2. Integrating a “nutritional profile” mode (vegetarian, athlete, etc.).

3. Enriching the ingredient database with open-source data.

4. Adding nutritional recap per recipe or per serving.

✨ Made with ❤️ in Go & React — powered by Mistral:Instruct (FR/EU)
