# DigitalOcean

DigitalOcean (DO) deployment.

## Introduction

This doc describes an example deployment to a DO Kubernetes cluster.

## Create Kubernetes cluster

Additonal DO applications:

* NGINX Ingress Controller
* Kubernetes Monitoring Stack

Managed DO databases:

* PostgreSQL 10

## Database

Create a new database with a new user. Remember VPC network connection parameters, for example:

```text
username = admin
password = ****************
host = private-db-postgresql-demo-do-user-9017261-0.b.db.ondigitalocean.com
port = 25060
database = defaultdb
sslmode = require
```

## Sealed Secrets

### Install Sealed Secrets

Based on <https://github.com/digitalocean/Kubernetes-Starter-Kit-Developers/tree/main/08-kubernetes-sealed-secrets>.

Deploying Sealed Secrets:

```sh
helm repo add sealed-secrets https://bitnami-labs.github.io/sealed-secrets
helm repo update
HELM_CHART_VERSION="2.1.2" helm install sealed-secrets-controller sealed-secrets/sealed-secrets --version "${HELM_CHART_VERSION}" --namespace sealed-secrets --create-namespace -f kubernetes/sealed-secrets/sealed-secrets-values.yaml
```

Installing kubeseal:

```sh
go install github.com/bitnami-labs/sealed-secrets/cmd/kubeseal@v0.17.3
```

### Generate Sealed Secrets

Example for generating sealed secret:

```sh
kubectl create secret generic -n foodstore meal-secret --from-literal=SERVICE_DB_DSN="host=private-db-postgresql-demo-do-user-123456-8.b.db.ondigitalocean.com user=USER password=PASSWORD dbname=DATABASE port=25060 sslmode=require" --from-literal=SERVICE_JWT_KEY="9876" -o json --dry-run=client | kubeseal --controller-namespace sealed-secrets -o yaml
```

Copy the printed out `SERVICE_DB_DSN` and `SERVICE_JWT_KEY` value to `kubernetes/foodstore/digitalocean_values.yaml:.mealService`.

## Ingress

Based on <https://github.com/digitalocean/Kubernetes-Starter-Kit-Developers/blob/main/03-setup-ingress-controller/nginx.md>.

Checking DO network setup:

```sh
doctl compute load-balancer list --format IP,ID,Name,Status
doctl compute load-balancer list --format IP --no-header
doctl compute domain list
whois pgillich.com
```

Register the load balancer IP at a DNS server, for example: `foodstore-pgillich.mooo.com`.

Checking the name resolution:

```sh
whois foodstore-pgillich.mooo.com
nslookup foodstore-pgillich.mooo.com
```

## Deployment

### Service

Review `kubernetes/foodstore/digitalocean_values.yaml`.

To install, run below command:

```sh
helm install --values ./kubernetes/foodstore/digitalocean_values.yaml --create-namespace --namespace foodstore foodstore ./kubernetes/foodstore
```

To upgrade, run below command:

```sh
helm upgrade --values ./kubernetes/foodstore/digitalocean_values.yaml --namespace foodstore foodstore ./kubernetes/foodstore
```
