#!/bin/bash
docker pull abdulrahmantweeq/transaction-service:v4
docker pull abdulrahmantweeq/analytics-service:v4

kubectl apply -f k8s/namespace.yaml

kubectl apply -f k8s/secrets.yaml
kubectl apply -f k8s/configmap.yaml

# Transaction
kubectl apply -f Transaction/k8s/deployment.yaml
kubectl apply -f Transaction/k8s/service.yaml

# Analytics
kubectl apply -f Analytics/k8s/deployment.yaml
kubectl apply -f Analytics/k8s/service.yaml


# kubectl port-forward -n go-sales-stream svc/transaction-service 9090:80
# kubectl port-forward -n go-sales-stream svc/transaction-service 9091:80