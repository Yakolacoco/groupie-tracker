---

# 📌 **Groupie Tracker – Application Go / Fyne**

Cette application affiche les artistes du projet **Groupie Tracker**, leurs informations, leurs images, ainsi que **les lieux et dates de concerts** grâce à l’API officielle.

Elle est développée en **Go** avec l’interface graphique **Fyne**.

---

## 🚀 **Fonctionnalités**

- Affichage de tous les artistes (nom, image, membres, création, premier album)
- Recherche dynamique (nom + membres)
- Affichage détaillé d’un artiste
- Chargement et affichage :
  - des **lieux** de concerts
  - des **dates** associées
  - via l’endpoint **Relations** de l’API
- Interface graphique simple et fluide (Fyne)

---

## 📡 **API utilisée**

L’application utilise l’API officielle :

```
https://groupietrackers.herokuapp.com/api
```

Endpoints utilisés :

- `/artists` → liste des artistes
- `relations` (URL fournie par chaque artiste) → lieux + dates

---

## 🧱 **Structure du projet**

```
groupie-tracker/
│
├── main.go
│── go.mod
│── go.sum
└── internal/
    ├── types.go       
    ├── fetch.go     
    ├── search.go       
    ├── geo.go   
```

---

## 🛠️ **Installation**

### 1) Installer les dépendances

Assurez-vous d’avoir Go installé.

Puis installez Fyne :

```
go get fyne.io/fyne/v2
```

---

## ▶️ **Lancer l’application**

Dans le dossier du projet :

```
go run main.go
```

L’interface graphique s’ouvre automatiquement.

---

## 🧩 **Fonctionnement**

### Chargement des artistes

```go
artists, _ := internal.LoadArtists()
```

### Chargement des concerts (lieux + dates)

```go
rel, _ := internal.LoadRelations(ar.Relations)
```

### Affichage dans l’interface

- Image de l’artiste
- Informations principales
- Liste des concerts :

```
Paris :
  - 2018-06-12
  - 2019-07-03

London :
  - 2017-05-21
```

---

## 📷 **Interface**

- Colonne gauche : recherche + liste des artistes
- Colonne droite : image + détails + concerts

---
