# kube-prometheus-go
# Requirements
- `Docker` installed
- `terraform` installed
- `kubectl` installed
- `helm` installed
- `kind` installed

# Setup Environment
**- first: move to the terraform environment directory**
```shell
  cd Environments/Terraform
```
**- second: run the script to apply the terraform configuration**
```shell
    sh Scripts/terraform-apply.sh
```
# Setup ArgoCD
**- first: move to the argocd environment directory**
```shell
  cd Environments/Manifest/Argocd
```
**- second: apply the argocd manifests using the script**
```shell
  sh scripts/install-argocd.sh
```
# Deploy Go Application
**- first: move to the go application environment directory**
```shell
  cd Environments/Manifest/Applications/go-application
```
**- second: apply the go application manifests using the script**
```shell
  sh scripts/install-go-app.sh
```
