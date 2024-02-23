# k8s-namespace-cloner

A simple Golang application designed to clone a Kubernetes namespace into a target namespace, including specific resources.

## Overview

The `k8s-namespace-cloner` tool facilitates the duplication of resources from one Kubernetes namespace to another. This can be particularly useful in scenarios where you want to replicate a namespace's configuration for testing, backup, or migration purposes.

## Prerequisites

- Go 1.x
- Kubernetes cluster access configured with `kubectl`
- Annotate your namespaces to be cloned with this annotation `cloner.io/enabled:True`. E.g. Your namespace definition will look like the below:
```
apiVersion: v1
kind: Namespace
metadata:
  annotations:
    cloner.io/enabled: "True" ## Notice the annotation
  labels:
    kubernetes.io/metadata.name: vv
  name: mynamespace
spec:
  finalizers:
  - kubernetes
```


## Features
- Clones a source namespace with the above annotation
- For safety in public cloud environments, only clones services of type ClusterIP, ExternalName and NodePort. Does not clone Loadbalancer service types - as this will create external DNS names if allowed

## Installation

Clone the repository to your local machine:
```
go run main.go
```

## Initialization for Cluster (Pre-Requisites)
#### Sample Application Deployment in cluster for testing
```
cd sample-charts
helm template --dry-run sample-app sample-app -f sample-app/values.yaml > deploy.yaml
kubectl apply -f deploy.yaml
```
The above will create a new namespace called 'sample' (specified in the `sample-charts/sample-app/values.yaml`). In case you want to create more such namespaces/applications,
simply change the value of the `namespace` in `values.yaml`, redo `helm template --dry-run sample-app sample-app -f sample-app/values.yaml > deploy.yaml && kubectl apply -f deploy.yaml` and see multiple namespaces

Note: In case the above annotation isn't available in the source namespace, the cloner will return empty on the `/v1/namespaces` API.
### Kind Cluster & Kube Green Deployment
(Note: Instructions with Kind Cluster at https://kube-green.dev/docs/tutorials/kind/. This is for local development)
- Ensure all the namespaces to be cloned have the `cloner.io/enabled:True` (Look at the annotations section above)
- Install Cert Manager if not already (Ref: https://cert-manager.io/docs/installation/kubectl)
```
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.2/cert-manager.yaml
kubectl get pods -n cert-manager ## Ensure the pods are going into a running state
```
- Install kube-green (https://kube-green.dev/docs/install/)
```
kubectl apply -f https://github.com/kube-green/kube-green/releases/latest/download/kube-green.yaml
kubectl get pods -n kube-green ## Ensure the pods are going into a running state
```



## Generating Documentation in Markdown:
`npm install -g widdershins
widdershins --search false --language_tabs 'shell:Shell' 'javascript:JavaScript' --summary docs/swagger.json -o docs/swagger.md
`

## Routes:
Use the swagger documentation at `docs/swagger.json` to load the routes. The following explains a bunch of routes. Route Documentation Available at [docs/swagger.md](docs/swagger.md)

Apply the file `rolebinding.yaml` onto the cluster for giving full operations to this service account for running the code in cluster
```kubectl apply -f rolebinding.yaml```

## Generating Documentation
```
go get -u github.com/swaggo/swag/cmd/swag

```
The Documentation is available in `docs/` folder

## Running
Start in development mode:

`go run main.go`

Start in Production Mode:
`go run main.go -production`


