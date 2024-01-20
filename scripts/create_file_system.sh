#!/bin/bash

# Create main directory
mkdir docker-compose-ci-cd
cd docker-compose-ci-cd

# Create Prometheus directory
mkdir prometheus
touch prometheus/prometheus.yml

# Create Grafana directory
mkdir grafana
mkdir grafana/provisioning
mkdir grafana/provisioning/dashboards
touch grafana/provisioning/dashboards/your-dashboard.json

# Create Kafka directory
mkdir kafka
cat <<EOF > kafka/docker-compose.yml
version: '3.8'

services:
  zookeeper:
    image: wurstmeister/zookeeper:latest
    container_name: zookeeper
    ports:
      - "2181:2181"

  kafka:
    image: wurstmeister/kafka:latest
    container_name: kafka
    ports:
      - "9092:9092"
    expose:
      - "9093"
    environment:
      KAFKA_ADVERTISED_LISTENERS: INSIDE://kafka:9093,OUTSIDE://localhost:9092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: INSIDE:PLAINTEXT,OUTSIDE:PLAINTEXT
      KAFKA_LISTENERS: INSIDE://0.0.0.0:9093,OUTSIDE://0.0.0.0:9092
      KAFKA_INTER_BROKER_LISTENER_NAME: INSIDE
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181

networks:
  default:
    driver: bridge
EOF

# Create your-app directory
mkdir your-app
touch your-app/Dockerfile

# Display success message
echo "File system structure created successfully!"

