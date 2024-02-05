---
title: Kubernetes Namespace Cloner API v3.0.0
language_tabs:
  - shell: Shell
  - javascript: JavaScript
language_clients:
  - shell: ""
  - javascript: ""
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="kubernetes-namespace-cloner-api">Kubernetes Namespace Cloner API v3.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Kubernetes Namespace Cloner API URI:<br>&nbsp;&nbsp;https://{nw-server-hostname}:8080/api/v1<br><br>

Base URLs:

* <a href="http://localhost:8080/api/v1">http://localhost:8080/api/v1</a>

* <a href="https://localhost:8080/api/v1">https://localhost:8080/api/v1</a>

<h1 id="kubernetes-namespace-cloner-api-default">Default</h1>

## Clone a namespace

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:8080/api/v1/cloneNamespace \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "sourceNamespace": "string",
  "targetNamespace": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/v1/cloneNamespace',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

`POST /cloneNamespace`

Clone a namespace and its objects to a new namespace

> Body parameter

```json
{
  "sourceNamespace": "string",
  "targetNamespace": "string"
}
```

<h3 id="clone-a-namespace-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[controllers.NSClonerRequestBody](#schemacontrollers.nsclonerrequestbody)|true|Namespace clone request body|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="clone-a-namespace-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|

<aside class="success">
This operation does not require authentication
</aside>

## Get all namespaces

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:8080/api/v1/namespaces \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/v1/namespaces',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

`GET /namespaces`

Get all namespaces in the cluster

> Example responses

> 200 Response

```json
[
  "string"
]
```

<h3 id="get-all-namespaces-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|

<h3 id="get-all-namespaces-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## Display config maps for a specific namespace

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:8080/api/v1/namespaces/:namespace/configmaps/display \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/v1/namespaces/:namespace/configmaps/display',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

`GET /namespaces/:namespace/configmaps/display`

Display all config maps in the specified namespace

<h3 id="display-config-maps-for-a-specific-namespace-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|Namespace name|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="display-config-maps-for-a-specific-namespace-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|

<aside class="success">
This operation does not require authentication
</aside>

## Display deployments for a specific namespace

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:8080/api/v1/namespaces/:namespace/deployments/display \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/v1/namespaces/:namespace/deployments/display',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

`GET /namespaces/:namespace/deployments/display`

Display all deployments in the specified namespace

<h3 id="display-deployments-for-a-specific-namespace-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|Namespace name|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="display-deployments-for-a-specific-namespace-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|

<aside class="success">
This operation does not require authentication
</aside>

## Display secrets for a specific namespace

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:8080/api/v1/namespaces/:namespace/secrets/display \
  -H 'Accept: application/json'

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/v1/namespaces/:namespace/secrets/display',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

`GET /namespaces/:namespace/secrets/display`

Display all secrets in the specified namespace

<h3 id="display-secrets-for-a-specific-namespace-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|Namespace name|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="display-secrets-for-a-specific-namespace-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|

<aside class="success">
This operation does not require authentication
</aside>

## Update a config map

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:8080/api/v1/updateConfigMap \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "data": {
    "property1": "string",
    "property2": "string"
  },
  "name": "string",
  "namespace": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/v1/updateConfigMap',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

`POST /updateConfigMap`

Update a config map in a specific namespace

> Body parameter

```json
{
  "data": {
    "property1": "string",
    "property2": "string"
  },
  "name": "string",
  "namespace": "string"
}
```

<h3 id="update-a-config-map-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[controllers.ConfigMapPatchRequestBody](#schemacontrollers.configmappatchrequestbody)|true|ConfigMap Update Request Body|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="update-a-config-map-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|

<aside class="success">
This operation does not require authentication
</aside>

## Update deployment image

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:8080/api/v1/updateDeploymentImage \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "container": "string",
  "deployment": "string",
  "image": "string",
  "namespace": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/v1/updateDeploymentImage',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

`POST /updateDeploymentImage`

Update the image of a deployment in a specific namespace

> Body parameter

```json
{
  "container": "string",
  "deployment": "string",
  "image": "string",
  "namespace": "string"
}
```

<h3 id="update-deployment-image-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[controllers.DeploymentPatchRequestBody](#schemacontrollers.deploymentpatchrequestbody)|true|Deployment Image Set Request Body|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="update-deployment-image-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|

<aside class="success">
This operation does not require authentication
</aside>

## Update a secret

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:8080/api/v1/updateSecret \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```javascript
const inputBody = '{
  "data": {},
  "name": "string",
  "namespace": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/v1/updateSecret',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

`POST /updateSecret`

Update a secret in a specific namespace

> Body parameter

```json
{
  "data": {},
  "name": "string",
  "namespace": "string"
}
```

<h3 id="update-a-secret-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[controllers.SecretPatchRequestBody](#schemacontrollers.secretpatchrequestbody)|true|Secret Update Request Body|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="update-a-secret-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_controllers.ConfigMapPatchRequestBody">controllers.ConfigMapPatchRequestBody</h2>
<!-- backwards compatibility -->
<a id="schemacontrollers.configmappatchrequestbody"></a>
<a id="schema_controllers.ConfigMapPatchRequestBody"></a>
<a id="tocScontrollers.configmappatchrequestbody"></a>
<a id="tocscontrollers.configmappatchrequestbody"></a>

```json
{
  "data": {
    "property1": "string",
    "property2": "string"
  },
  "name": "string",
  "namespace": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|data|object|false|none|none|
|» **additionalProperties**|string|false|none|none|
|name|string|false|none|none|
|namespace|string|false|none|none|

<h2 id="tocS_controllers.DeploymentPatchRequestBody">controllers.DeploymentPatchRequestBody</h2>
<!-- backwards compatibility -->
<a id="schemacontrollers.deploymentpatchrequestbody"></a>
<a id="schema_controllers.DeploymentPatchRequestBody"></a>
<a id="tocScontrollers.deploymentpatchrequestbody"></a>
<a id="tocscontrollers.deploymentpatchrequestbody"></a>

```json
{
  "container": "string",
  "deployment": "string",
  "image": "string",
  "namespace": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|container|string|false|none|none|
|deployment|string|false|none|none|
|image|string|false|none|none|
|namespace|string|false|none|none|

<h2 id="tocS_controllers.NSClonerRequestBody">controllers.NSClonerRequestBody</h2>
<!-- backwards compatibility -->
<a id="schemacontrollers.nsclonerrequestbody"></a>
<a id="schema_controllers.NSClonerRequestBody"></a>
<a id="tocScontrollers.nsclonerrequestbody"></a>
<a id="tocscontrollers.nsclonerrequestbody"></a>

```json
{
  "sourceNamespace": "string",
  "targetNamespace": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|sourceNamespace|string|false|none|none|
|targetNamespace|string|false|none|none|

<h2 id="tocS_controllers.SecretPatchRequestBody">controllers.SecretPatchRequestBody</h2>
<!-- backwards compatibility -->
<a id="schemacontrollers.secretpatchrequestbody"></a>
<a id="schema_controllers.SecretPatchRequestBody"></a>
<a id="tocScontrollers.secretpatchrequestbody"></a>
<a id="tocscontrollers.secretpatchrequestbody"></a>

```json
{
  "data": {},
  "name": "string",
  "namespace": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|data|object|false|none|none|
|name|string|false|none|none|
|namespace|string|false|none|none|

