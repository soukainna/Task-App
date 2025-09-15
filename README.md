# Task-App en Go avec MySQL, Docker

Ce projet est une application de gestion de tâches ("task app") codée en **Go**, avec une base de données **MySQL**, le tout conteneurisé avec **Docker** et **Docker Compose**.

---

## 🔎 Objectif principal

L'application expose une API REST en Go pour manipuler des tâches (création, lecture, modification, suppression), persistées dans une base MySQL.

---

## 📊 Fonctionnalités principales

* API RESTful : GET, POST, PUT, PATCH, DELETE sur `/tasks`
* Connexion à une base **MySQL** via **GORM** (ORM Go)
* Endpoint `/health` pour vérifier que le backend tourne
* Frontend simple en HTML/JavaScript (AJAX)
* Intégration continue via GitHub Actions (CI)
* ❄️ Tests automatisés Go avec `go test`

---

## 🧰 Architecture technique

```
- task-service/
  - main.go               => Point d’entrée de l’application
  - models/               => Structure d’une tâche (Task)
  - controllers/          => Logique des routes HTTP
  - database/             => Connexion MySQL (GORM)
  - routes/               => Routing (mux + handlers)

- frontend/
  - index.html            => Interface utilisateur
  - app.js                => Appels AJAX Go REST

- Dockerfile              => Image de l’API Go
- docker-compose.yml      => Lancement multi-service
- .github/workflows/      => CI/CD GitHub Actions
```

---

## 🚀 Lancement rapide (Docker Compose)

### Script de lancement Docker : `start.sh`

```bash
#!/bin/bash
# Nettoyer les anciens conteneurs
docker-compose down

# Rebuild les images
docker-compose build

# Lancer en arrière-plan
docker-compose up -d

# Logs de l'application
docker-compose logs -f task-service
```

> Rends le script exécutable avec : `chmod +x start.sh`

### 1. Lancer les services Docker

```bash
./start.sh
```

---

## 🚪 Accès aux services

| Service     | URL                                            | Description                         |
| ----------- | ---------------------------------------------- | ----------------------------------- |
| Backend API | [http://localhost:8080](http://localhost:8080) | API REST (GET/POST/PATCH... /tasks) |
| Frontend    | [http://localhost:8081](http://localhost:8081) | UI HTML/JS                          |

---

## ⚖️ Exemple d’utilisation de l’API (via curl)

### Créer une tâche

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Faire les courses", "completed":false}'
```

### Modifier une tâche (PATCH)

```bash
curl -X PATCH http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"completed":true}'
```

### Supprimer une tâche

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

### Liste des tâches

```bash
curl http://localhost:8080/tasks
```

---

## ✅ Tests unitaires

### Lancer les tests (via Docker uniquement)

```bash
docker-compose run --rm tester
```

> ⚠️ Les tests utilisent la base de données MySQL définie dans `docker-compose.yml`. Ils doivent être exécutés dans l'environnement Docker pour fonctionner correctement.

---

## 🚜 CI/CD GitHub Actions

* Le pipeline `go.yml` fait :

  * Setup Go 1.22
  * Lancement d’un service MySQL
  * Build de l’API
  * Tests

