#!/bin/bash
docker pull abdulrahmantweeq/transaction-service:v5
docker pull abdulrahmantweeq/analytics-service:v5

kubectl apply -f server/namespace.yaml

kubectl apply -f server/secrets.yaml
kubectl apply -f server/configmap.yaml

# Transaction
kubectl apply -f transaction/server/resourses/deployment.yaml
kubectl apply -f transaction/server/resourses/service.yaml

# Analytics
kubectl apply -f analytics/server/resourses/deployment.yaml
kubectl apply -f analytics/server/resourses/service.yaml


# kubectl port-forward -n go-sales-stream svc/transaction-service 9090:80
# kubectl port-forward -n go-sales-stream svc/transaction-service 9091:80