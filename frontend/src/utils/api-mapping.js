const apiMapping = {
    getRequestObject: (feature, params) => {
        let requestObject = {};
        const backendHost = 'http://localhost:8080';
  
        switch (feature) {
            case 'FETCH_ALL_NAMESPACE_DEPLOYMENTS': {
                const { namespaceName } = params;
                requestObject = {
                    method: 'GET',
                    url: `${backendHost}/api/v1/namespaces/${namespaceName}/deployments/display`,
                };
                break;
            }

            case 'FETCH_ALL_NAMESPACE_CONFIG_MAPS': {
                const { namespaceName } = params;
                requestObject = {
                    method: 'GET',
                    url: `${backendHost}/api/v1/namespaces/${namespaceName}/configmaps/display`,
                };
                break;
            }

            case 'FETCH_ALL_NAMESPACE_SECRETS': {
                const { namespaceName } = params;
                requestObject = {
                    method: 'GET',
                    url: `${backendHost}/api/v1/namespaces/${namespaceName}/secrets/display`,
                };
                break;
            }

            case 'UPDATE_DEPLOYMENT_IMAGE': {
                const { image, deployment } = params;
                const container = Object.keys(deployment.containers[0])[0];
                requestObject = {
                    method: 'POST',
                    url: `${backendHost}/api/v1/deployments/${deployment.name}`,
                    data: {
                        container,
                        image,
                        namespace: deployment.namespace,
                    }
                };
                break;
            }

            case 'INCREASE_DEPLOYMENT_REPLICAS': {
                const { deployment } = params;
                requestObject = {
                    method: 'POST',
                    url: `${backendHost}/api/v1/deployments/${deployment.name}/scaleup`,
                };
                break;
            }

            case 'UPDATE_CONFIG_MAP': {
                const { updateObject, namespaceName, configMap } = params;
                requestObject = {
                    method: 'POST',
                    url: `${backendHost}/api/v1/configmaps/${configMap}`,
                    data: {
                        data: updateObject,
                        namespace: namespaceName,
                    }
                };
                break;
            }

            case 'UPDATE_SECRET': {
                const { updateObject, namespaceName, secret } = params;
                requestObject = {
                    method: 'POST',
                    url: `${backendHost}/api/v1/secrets/${secret}`,
                    data: {
                        data: updateObject,
                        namespace: namespaceName,
                    }
                };
                break;
            }

            case 'CLONE_NAMESPACE': {
                const { targetNamespace, namespaceName } = params;
                requestObject = {
                    method: 'POST',
                    url: `${backendHost}/api/v1/namespaces/${namespaceName}/cloneNamespace`,
                    data: {
                        targetNamespace,
                    }
                };
                break;
            }

            case 'FETCH_ALL_NAMESPACES': {
                requestObject = {
                    method: 'GET',
                    url: `${backendHost}/api/v1/namespaces`,
                };
                break;
            }

            default: {
                console.log(`No API mapping for the provided feature ${feature}`);
            }
        }

        return requestObject;
    }
}

export default apiMapping;