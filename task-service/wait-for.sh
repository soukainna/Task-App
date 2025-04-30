#!/bin/sh
set -e

echo "En attente que PostgreSQL (db:5432) réponde..."

until nc -z db 5432; do
  echo "PostgreSQL pas encore prêt... attente..."
  sleep 1
done

echo "PostgreSQL est prêt 🎉"
exec ./task-service
