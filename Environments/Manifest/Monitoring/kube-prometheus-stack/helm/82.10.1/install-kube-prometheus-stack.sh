#!/usr/bin/env bash

set -e

# Add the Prometheus Community Helm repository
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# Install the kube-prometheus-stack Helm chart
helm install my-kube-prometheus-stack prometheus-community/kube-prometheus-stack --version 82.10.1 --namespace monitoring --create-namespace
