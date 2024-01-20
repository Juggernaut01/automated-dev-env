#!/bin/bash

# Deploy Zookeeper and Kafka
echo "Deploying Zookeeper and Kafka..."
docker-compose -f docker-compose.kafka.yml up -d

# Wait for Kafka to start
echo "Waiting for Kafka to start..."
sleep 10

# Deploy Jenkins
echo "Deploying Jenkins..."
docker-compose -f docker-compose.jenkins.yml up -d

# Deploy GitLab
echo "Deploying GitLab..."
docker-compose -f docker-compose.gitlab.yml up -d

# Deploy PostgreSQL
echo "Deploying PostgreSQL..."
docker-compose -f docker-compose.postgres.yml up -d

# Deploy Prometheus and Grafana
echo "Deploying Prometheus and Grafana..."
docker-compose -f docker-compose.monitoring.yml up -d

# Deploy your application
echo "Deploying your application..."
docker-compose -f docker-compose.app.yml up -d

echo "Services deployed successfully!"

