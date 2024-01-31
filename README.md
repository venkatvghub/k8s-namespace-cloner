# k8s-namespace-cloner

A simple Golang application designed to clone a Kubernetes namespace into a target namespace, including specific resources.

## Overview

The `k8s-namespace-cloner` tool facilitates the duplication of resources from one Kubernetes namespace to another. This can be particularly useful in scenarios where you want to replicate a namespace's configuration for testing, backup, or migration purposes.

## Prerequisites

- Go 1.x
- Kubernetes cluster access configured with `kubectl`
- Annotate your namespaces to be cloned with this annotation `cloner.k8s.io/enabled:True`. E.g. Your namespace definition will look like the below:
```
apiVersion: v1
kind: Namespace
metadata:
  annotations:
    cloner.k8s.io/enabled: "True" ## Notice the annotation
  labels:
    kubernetes.io/metadata.name: vv
  name: mynamespace
spec:
  finalizers:
  - kubernetes
```
Note: In case the above annotation isn't available in the source namespace, the cloner will return empty on the `/v1/namespaces` API.

## Features
- Clones a source namespace with the above annotation
- For safety in public cloud environments, only clones services of type ClusterIP, ExternalName and NodePort. Does not clone Loadbalancer service types - as this will create external DNS names if allowed

## Installation

Clone the repository to your local machine:
```
go run main.go
curl http://localhost:8080//api/v1/namespaces ## List Namespaces
curl http://localhost:8080/api/v1/namespaces/:namespace/deployments ## List Deployments for namespace
curl -X POST -H "Content-Type: application/json" -d '{"targetNamespace": "targetNamespace"}' http://localhost:8080/api/v1/namespaces/:sourceNamespace/clone ## Clones the source namespace to a target namespace
```
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


