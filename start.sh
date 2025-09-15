#!/bin/bash
echo " Arrêt des anciens conteneurs..."
docker-compose down

echo " Rebuild des images..."
docker-compose build

echo " Démarrage des services..."
docker-compose up -d

echo " Logs de l'API :"
docker-compose logs -f task-service
