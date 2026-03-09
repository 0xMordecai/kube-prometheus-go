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
**- move to the go application environment directory**
```shell
  cd Environments/Manifest/Applications/go-application
```
**- create go-application Namespace**
```shell
  kubectl apply -f namespace.yaml
```
**- apply Deployment manifests**
```shell
  kubectl apply -f Deployment.yaml
```
**- apply Service manifests**
```shell
  kubectl apply -f Service.yaml
```

# Deploy kube-prometheus
**we will use `helm` to deploy kube-prometheus-stack**

**- Add the Prometheus Community Helm repository**
```shell
  helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
  helm repo update
```
**- Deploy the chart into a new namespace "monitoring"**
```shell
  helm install my-kube-prometheus-stack prometheus-community/kube-prometheus-stack --version 82.10.1 --namespace monitoring --create-namespace
```
**- or you can use the script within the directory**
**head to the kube-prometheus-stack directory**
```shell
  cd Environments/Manifest/Monitoring/kube-prometheus-stack/helm/82.10.1
```
```shell
  sh install-kube-prometheus-stack
```
