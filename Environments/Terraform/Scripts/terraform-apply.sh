#!/usr/bin/env bash
# a bash script to automate terraform deployment

set -e
# Helpers
log() {
  echo "[INFO] $1"
}
# Prepare your working directory for other commands
log "Initializing Terraform"
terraform init
# Check whether the configuration is valid
log "Validating configuration"
terraform validate
# Show changes required by the current configuration
log "Creating plan"
terraform plan
# Create or update infrastructure
log "Applying plan"
terraform apply -auto-approve -var="branch=${GITHUB_REF##*/}"

log "Cluster creation completed successfully"
