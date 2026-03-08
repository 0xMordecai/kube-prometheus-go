#!/usr/bin/env bash

set -e

kubectl delete --ignore-not-found=true -f manifests/ -f manifests/setup
