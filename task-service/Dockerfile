FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# 👇 Installer netcat pour tester la connexion au port 5432
RUN apt-get update && apt-get install -y netcat-openbsd

# Compiler le binaire
RUN go build -o task-service

# Rendre le script exécutable
RUN chmod +x wait-for.sh

EXPOSE 8080

# Lancer l'app avec le script d'attente
CMD ["./wait-for.sh"]
