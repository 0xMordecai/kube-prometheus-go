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
**- To Destroy terraform resources**
```shell
    sh Scripts/terraform-destroy.sh
```

# Deploy Go Application
**- first: move to the go application environment directory**
```shell
  cd Environments/Manifest/Applications/go-application
```
**- second: create go-application Namespace**
```shell
  kubectl apply -f namespace.yaml
```
**- third: apply Deployment manifests**
```shell
  kubectl apply -f Deployment.yaml
```
**- fourth: apply Service manifests**
```shell
  kubectl apply -f NodePort.yaml
```
