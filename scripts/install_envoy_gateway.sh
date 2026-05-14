#!/usr/bin/env csh

# helm status envoy-gateway -n envoy-gateway

#CHART_VERSION=v0.0.0-latest
CHART_VERSION=v1.8.0
NAME=envoy-gateway
helm install $NAME oci://docker.io/envoyproxy/gateway-helm \
  --version $CHART_VERSION \
  -n envoy-gateway \
  --create-namespace \
  -f ./helm/envoy-gateway/values.yaml


# helm uninstall eg