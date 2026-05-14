#!/usr/bin/env bash

# This script installs the metrics server on the configured k8s cluster(kubectl configured cluster)
# ref: https://github.com/kubernetes-sigs/metrics-server

kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
