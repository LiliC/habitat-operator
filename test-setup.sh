#!/bin/sh
# Setup for Travis-CI.

# kubectl is a requirement for using minikube
curl -Lo kubectl https://storage.googleapis.com/kubernetes-release/release/v1.7.0/bin/linux/amd64/kubectl && chmod +x kubectl && sudo mv kubectl /usr/local/bin/

# Get minikube.
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/

echo "starting minikube..."
sudo minikube start --vm-driver=none --kubernetes-version=v1.7.0

# Fix the kubectl context, as its stale.
minikube update-context

# Debug info.
echo "minikube status"
minikube status

echo "minikube logs"
minikube logs

echo "minikube ip"
minikube ip

# Hack for waiting for our minikube to be ready
# FIXME: replace with kubectl cluster-info to be ready
echo "sleeping..."
sleep 5m

# Testing if Kubernetes works.
echo "create random secret"
kubectl create secret generic user-toml-secret

echo "getting secret"
kubectl get secret

