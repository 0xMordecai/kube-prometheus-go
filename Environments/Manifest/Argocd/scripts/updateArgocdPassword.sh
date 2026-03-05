#!/usr/bin/env bash

set -e
kubectl patch secret argocd-secret -n argocd \
-p '{"data": {"admin.password": null, "admin.passwordMtime": null}}'

kubectl scale deployment argocd-server --replicas 0 -n argocd

kubectl scale deployment argocd-server --replicas 1 -n argocd
