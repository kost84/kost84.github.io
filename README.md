# 🌐 Portfolio & Hub de Projets - Lioris Racheter (Branche V2)

[![Go Version](https://img.shields.io/badge/Go-1.27+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=flat&logo=html5&logoColor=white)](https://developer.mozilla.org/fr/docs/Web/HTML)
[![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=flat&logo=css3&logoColor=white)](https://developer.mozilla.org/fr/docs/Web/CSS)
[![Status](https://img.shields.io/badge/Status-Actif%20(V2%20Dynamique)-success)]()

Bienvenue sur le dépôt de mon **Portfolio personnel & Hub de projets**, conçu et développé de A à Z dans le cadre de ma formation en **Bachelor 1 Informatique à Ynov Campus Aix-en-Provence**.

Ce projet a pour objectif de présenter mon parcours, mes compétences techniques, mes projets académiques et personnels, ainsi que mon CV aux recruteurs et entreprises, tout en servant de support évolutif tout au long de mon cursus.

---

## 🚀 Évolution vers la V2 (Serveur Web Dynamique en Go)

La branche **V2** marque une étape majeure dans le développement du projet : le passage d'un site statique à une **application web dynamique propulsée par un serveur backend en Go (Golang)**.

### Nouveautés de la V2 :
- **Serveur HTTP autonome en Go** (`main.go`) écoutant sur le port local `8080`.
- **Nouveau Menu principal interactif** inspiré du style Behance (`/` ou `/menu`), permettant de naviguer facilement entre les différents univers et projets.
- **Rendu dynamique avec `html/template`** : les templates HTML sont parsés et alimentés dynamiquement par le moteur de templates de Go.
- **Gestion des données via JSON (`menu.json` & `project.json`)** : les projets et éléments du menu ne sont plus codés en dur dans le HTML, mais chargés et désérialisés à la volée depuis des fichiers JSON (`encoding/json`), permettant des mises à jour simples et rapides.
- **Routage et gestion des assets statiques** : distribution automatisée des feuilles de style (`/css/`) et des fichiers multimédias/documents (`/imageetdoc/`).

---

## 🤖 Mention concernant l'Assistance par IA

> [!NOTE]
> **Utilisation de l'Intelligence Artificielle à des fins d'assistance CSS :**  
> Dans le cadre de ce projet, une **Intelligence Artificielle a été utilisée à des fins d'assistance pour le travail sur le CSS** (notamment pour l'optimisation de la disposition, l'harmonisation visuelle des couleurs, la clarté du menu et la mise en page responsive des pages telles que le menu, l'accueil, les projets et la page contact).  
> La structure générale, l'architecture backend en Go, la logique de routage, le templating dynamique et l'intégration des données restent entièrement conçus et implémentés par l'auteur.

---

## 📁 Structure du Projet

```text
kost84.github.io/
├── css/                             # Feuilles de style du site (assistées par IA)
│   ├── acceuil.css                  # Style de la page profil & compétences
│   ├── contact.css                  # Style de la page contact
│   ├── index.css                    # Style de la page d'accueil du portfolio
│   ├── menu.css                     # Style du menu principal (style Behance)
│   └── project.css                  # Style des cartes de projets
├── imageetdoc/                      # Médias, images et documents
│   ├── arriereplan/                 # Fonds d'écran et bannières
│   ├── doc/                         # Documents à télécharger (CV au format PDF)
│   ├── icone/                       # Icônes graphiques
│   └── imageproject/                # Captures d'écran d'illustration des projets
├── page/                            # Fichiers HTML / Templates Go
│   ├── acceuil.html                 # Présentation (Qui suis-je, compétences, stages)
│   ├── contact.html                 # Liens réseaux, CV et coordonnées
│   ├── index.html                   # Page d'accueil / vitrine du portfolio
│   ├── menu.html                    # Page du menu principal (style Behance)
│   ├── page3.html                   # Page complémentaire
│   └── project.html                 # Template dynamique d'affichage des projets
├── go.mod                           # Fichier de module Go
├── main.go                          # Point d'entrée du serveur web en Golang
├── menu.json                        # Données structurées des éléments du menu principal
├── project.json                     # Données structurées des projets affichés
└── README.md                        # Documentation du projet
```

---

## 🛠️ Technologies Utilisées

| Domaine | Technologies |
| :--- | :--- |
| **Backend** | [Go (Golang)](https://go.dev/) (`net/http`, `html/template`, `encoding/json`, `os`) |
| **Frontend** | HTML5, CSS3 *(avec assistance IA pour le style)* |
| **Format de données** | JSON (`menu.json`, `project.json`) |
| **Outils & Versioning** | Git, GitHub, VS Code |

---

## 💻 Installation et Lancement Local

### Prérequis
- Avoir installé [Go](https://go.dev/dl/) sur sa machine (version 1.20 ou supérieure recommandée).
- Avoir [Git](https://git-scm.com/) installé.

### Étapes d'exécution

1. **Cloner le dépôt :**
   ```bash
   git clone https://github.com/kost84/kost84.github.io.git
   cd kost84.github.io
   ```

2. **Basculer sur la branche V2 :**
   ```bash
   git checkout V2
   ```

3. **Lancer le serveur Go :**
   ```bash
   go run main.go
   ```

4. **Accéder au site :**  
   Ouvrez votre navigateur web à l'adresse suivante :  
   👉 [http://localhost:8080](http://localhost:8080)

---

## 📄 Routes & Pages Disponibles

- **`/` ou `/menu`** : Menu principal interactif (style Behance) répertoriant les différentes sections.
- **`/portfolio` ou `/index`** : Page de bienvenue et d'introduction au portfolio.
- **`/acceuil`** : Présentation personnelle, parcours académique, expériences professionnelles (stages) et compétences techniques (Hard & Soft Skills).
- **`/project`** : Galerie dynamique des projets réalisés (chargés depuis `project.json`).
- **`/contact`** : Coordonnées, lien vers le profil GitHub et téléchargement direct du CV en PDF.

---

## 👤 Auteur

- **Lioris Racheter** ([@kost84](https://github.com/kost84))
- Étudiant en Bachelor 1 Informatique — Ynov Campus Aix-en-Provence
