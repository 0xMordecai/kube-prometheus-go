#!/usr/bin/env bash

set -e


kubectl port-forward svc/argocd-server -n argocd 8080:443
