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
```

Generating Documentation in Markdown:
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


