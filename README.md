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

# Deploy kube-prometheus
**- first: head to the kube-prometheus-stack directory**, we will use helm to install kube-prometheus
```shell
  cd Environments/Manifest/Monitoring/kube-prometheus-stack/helm/82.10.1
```
**- second: Add the Prometheus Community Helm repository**
```shell
  helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```
**- third:Install the kube-prometheus-stack Helm chart**
```shell
  helm install my-kube-prometheus-stack prometheus-community/kube-prometheus-stack --version 82.10.1
```
**or you can use the script within the directory**
```shell
  sh install-kube-prometheus-stack
```
