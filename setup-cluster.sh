#!/bin/bash

set -e

echo "Creating kind cluster..."
kind create cluster

echo "Setting kubectl context to kind-kind..."
kubectl config use-context kind-kind

echo "Installing KServe using quickstart script..."
curl -s "https://raw.githubusercontent.com/kserve/kserve/release-0.15/hack/quick_install.sh" | bash

echo "Verifying installation..."
echo "Checking KServe pods:"
kubectl get pods -n kserve
echo "Checking Knative Serving pods:"
kubectl get pods -n knative-serving

echo "Cluster setup complete! You can now run the MCP server."
echo "To run the server: go run cmd/server/main.go" 