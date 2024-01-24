# k8s-namespace-cloner

A simple Golang application designed to clone a Kubernetes namespace into a target namespace, including specific resources.

## Overview

The `k8s-namespace-cloner` tool facilitates the duplication of resources from one Kubernetes namespace to another. This can be particularly useful in scenarios where you want to replicate a namespace's configuration for testing, backup, or migration purposes.

## Prerequisites

- Go 1.x
- Kubernetes cluster access configured with `kubectl`

## Installation

Clone the repository to your local machine:
```
go run main.go
curl http://localhost:8080/api/namespaces ## List Namespaces
curl http://localhost:8080/api/discover/myns/deployments ## List Deployments for namespace <myns>
curl -X POST -H "Content-Type: application/json" -d '{"targetNamespace": "target"}' http://localhost:8080/api/clone/myns ## Clones the source namespace <myns> to a target namespace <targer>
```

