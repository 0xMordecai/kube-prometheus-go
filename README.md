# kube-prometheus-go
# Requirements
- `Docker` installed
- `curl` installed
# To run the container
- first: pull the image from `hub.docker/kings5layer/prometheus-go-app`
```shell
    docker pull kings5layer/prometheus-go-app:46daaab
```
- second: run the container
```shell
    docker run -d --name go-app -p 2112:2112 kings5layer/prometheus-go-app:46daaab
```
# Access the metrics
- To access the metrics run:
```shell
  curl http://localhost:2112/metrics 
```