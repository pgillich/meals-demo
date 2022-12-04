# Meals demo

Demo for a meals service. It's written in Go.

<!-- markdownlint-disable MD013 -->

## Deployment

The deployment is worked out for a local Kubernetes described at <https://github.com/pgillich/kind-on-dev> with default values. Add `foodstore.kind-01.company.com` to the `/etc/hosts`, similar to `monitoring.kind-01.company.com`.

Another deployment example is described in [DIGITALOCEAN.md](DIGITALOCEAN.md).

### Postgres

Run below command:

```sh
kubectl apply -k kubernetes/postgres
```

After a while, the postgres service can be accessed on a Load Balancer port, see:

```sh
kubectl get svc -n postgres -o wide
NAME       TYPE           CLUSTER-IP      EXTERNAL-IP    PORT(S)          AGE   SELECTOR
postgres   LoadBalancer   10.96.224.184   172.18.1.128   5432:31113/TCP   31m   app=postgres
```

The connection can be checked by `psql`, and the database should be created, for example:

```sh
psql -h 172.18.1.128 -U admin --password -p 5432 postgresdb
Password:
psql (12.9 (Ubuntu 12.9-0ubuntu0.20.04.1), server 10.1)
Type "help" for help.

postgresdb=# create database foodstore;
postgresdb=# exit
```

The app compiled by `make build` can also be run out of the Kubernetes cluster, for examle:

```sh
PORT=8080 SERVICE_DB_DIALECT=postgres SERVICE_DB_DSN="host=172.18.1.128 user=admin password=test123 dbname=foodstore port=5432 sslmode=disable" SERVICE_DB_SAMPLE=true SERVICE_DB_DEBUG=true SERVICE_JWT_KEY="1234" ./build/bin/meals-demo
```

### Image

The built image should be accessible from the cluster. The simplest way to make it accessible for Kind is loading the image, for example:

```sh
kind load docker-image --name demo pgillich/meals-demo:v0.0.6
```

Another alternative is pulling the image from Docker Hub, see more info here: <https://hub.docker.com/r/pgillich/meals-demo/tags>

### Service

Review `kubernetes/foodstore/kind_values.yaml` and run below command:

```sh
helm install --values ./kubernetes/foodstore/kind_values.yaml --create-namespace --namespace foodstore foodstore ./kubernetes/foodstore
```

Above deployment is confugured for Traefik of <https://github.com/pgillich/kind-on-dev>, so the service can be accessed with `http://foodstore.kind-01.company.com`, for example:

```sh
curl http://foodstore.kind-01.company.com/v1/tags
```

## Usage

The API is specified in `api/foodstore.yaml`, in OpenAPI 2.0. It can be seen in browser at <http://localhost:8088/docs>, after running below command:

```sh
make openapi-view
```

### Examples, getting info

The `/v1/livez` endpoint is used by Kubernetes.

The `/v1/version` returns build information, for example:

```sh
curl 127.0.0.1:8080/v1/version
{"appName":"meals-demo","buildTime":"2022-02-06T14:19:15+0100","goMod":"module github.com/pgillich/meals-demo\n\ngo 1.17\n\nrequire (\n\tgithub.com/go-openapi/errors v0.20.1\n\tgithub.com/go-openapi/loads v0.21.1\n\tgithub.com/go-openapi/runtime v0.23.0\n\tgithub.com/go-openapi/spec v0.20.4\n\tgithub.com/go-openapi/strfmt v0.21.0\n\tgithub.com/go-openapi/swag v0.19.15\n\tgithub.com/go-openapi/validate v0.20.3\n\tgithub.com/jessevdk/go-flags v1.5.0\n\tgithub.com/stretchr/testify v1.7.0\n\tgolang.org/x/net v0.0.0-20210421230115-4e50805a0758\n)\n\nrequire (\n\temperror.dev/errors v0.8.0 // indirect\n\tgithub.com/PuerkitoBio/purell v1.1.1 // indirect\n\tgithub.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect\n\tgithub.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect\n\tgithub.com/davecgh/go-spew v1.1.1 // indirect\n\tgithub.com/docker/go-units v0.4.0 // indirect\n\tgithub.com/go-openapi/analysis v0.21.2 // indirect\n\tgithub.com/go-openapi/jsonpointer v0.19.5 // indirect\n\tgithub.com/go-openapi/jsonreference v0.19.6 // indirect\n\tgithub.com/go-stack/stack v1.8.0 // indirect\n\tgithub.com/jinzhu/gorm v1.9.16 // indirect\n\tgithub.com/jinzhu/inflection v1.0.0 // indirect\n\tgithub.com/josharian/intern v1.0.0 // indirect\n\tgithub.com/lib/pq v1.1.1 // indirect\n\tgithub.com/mailru/easyjson v0.7.6 // indirect\n\tgithub.com/mattn/go-sqlite3 v1.14.0 // indirect\n\tgithub.com/mitchellh/mapstructure v1.4.1 // indirect\n\tgithub.com/oklog/ulid v1.3.1 // indirect\n\tgithub.com/pkg/errors v0.9.1 // indirect\n\tgithub.com/pmezard/go-difflib v1.0.0 // indirect\n\tgo.mongodb.org/mongo-driver v1.7.3 // indirect\n\tgo.uber.org/atomic v1.7.0 // indirect\n\tgo.uber.org/multierr v1.6.0 // indirect\n\tgolang.org/x/sys v0.0.0-20210420072515-93ed5bcd2bfe // indirect\n\tgolang.org/x/text v0.3.7 // indirect\n\tgopkg.in/yaml.v2 v2.4.0 // indirect\n\tgopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect\n)\n","version":"v0.0.2-1-g8cb47e5"}
```

### Examples, listing all records

Geting all records:

```sh
curl 127.0.0.1:8080/v1/tags
[{"id":3,"name":"gluten free"},{"id":1,"name":"spicy"},{"id":2,"name":"vegan"}]

curl 127.0.0.1:8080/v1/ingredients
[{"description":"Tomato","id":6,"name":"tomato"},{"description":"Onion","id":7,"name":"onion"},{"description":"Tomato sauce","id":1,"name":"tomato sauce"},{"description":"Bacon","id":3,"name":"bacon"},{"description":"Salami","id":4,"name":"salami"},{"description":"Sour cream sauce","id":2,"name":"sour cream sauce"},{"description":"Mozzarella","id":5,"name":"mozzarella"}]

curl 127.0.0.1:8080/v1/meal/findByTag
[{"description":"Spicy pizza","id":1,"ingredients":[{"description":"Tomato sauce","id":1,"name":"tomato sauce"},{"description":"Bacon","id":3,"name":"bacon"},{"description":"Salami","id":4,"name":"salami"}],"kcal":123,"name":"Spicy","pictureUrl":"http://a.com","price":3.25,"tags":[{"id":1,"name":"spicy"}]},{"description":"Vegan pizza","id":2,"ingredients":[{"description":"Sour cream sauce","id":2,"name":"sour cream sauce"},{"description":"Mozzarella","id":5,"name":"mozzarella"}],"kcal":234,"name":"Vegan","pictureUrl":"http://a.com","price":4.1,"tags":[{"id":2,"name":"vegan"}]}]
```

### Examples, filtering meals

Getting meals filtered by tag ID:

```sh
curl 127.0.0.1:8080/v1/meal/findByTag?tag=1
[{"description":"Spicy pizza","id":1,"ingredients":[{"description":"Tomato sauce","id":1,"name":"tomato sauce"},{"description":"Bacon","id":3,"name":"bacon"},{"description":"Salami","id":4,"name":"salami"}],"kcal":123,"name":"Spicy","pictureUrl":"http://a.com","price":3.25,"tags":[{"id":1,"name":"spicy"}]}]
```

Getting only one meal by ID:

```sh
curl 127.0.0.1:8080/v1/meal/1
{"description":"Spicy pizza","id":1,"ingredients":[{"description":"Tomato sauce","id":1,"name":"tomato sauce"},{"description":"Bacon","id":3,"name":"bacon"},{"description":"Salami","id":4,"name":"salami"}],"kcal":123,"name":"Spicy","pictureUrl":"http://a.com","price":3.25,"tags":[{"id":1,"name":"spicy"}]}
```

### Examples, changing data

Example for getting JWT token:

```sh
curl -X POST -H 'Content-Type: application/json' 127.0.0.1:8080/v1/login -d '{"email":"yoda@star.wars", "password":"master"}'
```

Creating a new meal (the ID at the end of path is needed, but skipped):

```sh
curl -X POST -H 'Authorization: Bearer eyJh...' -H 'Content-Type: application/json' 127.0.0.1:8080/v1/meal/0 -d '{"description":"Tomato pizza","ingredients":[{"description":"Tomato sauce","id":1,"name":"tomato sauce"}],"kcal":200,"name":"Spicy","pictureUrl":"http://c.com","price":3.25,"tags":[{"id":3,"name":"gluten free"}]}'
```

Updating a meal (the ID at the end of path is needed, but skipped):

```sh
curl -X PUT -H 'Authorization: Bearer eyJh...' -H 'Content-Type: application/json' 127.0.0.1:8080/v1/meal/0 -d '{"description":"Tomato pizza","id":5,"ingredients":[{"description":"Tomato sauce","id":1,"name":"tomato sauce"}],"kcal":200,"name":"Spicy","pictureUrl":"http://c.com","price":3.55,"tags":[{"id":3,"name":"gluten free"}]}'
```

Deleting a meal:

```sh
curl -X DELETE -H 'Authorization: Bearer eyJh...' 127.0.0.1:8080/v1/meal/5
```

## Development

### Prerequisites

* Ubuntu
* Docker
* `git`
* `make`
* Golang compiler, linters, etc. are not needed (they are run in container)

### Building

Builds exacutable binary to `build/bin/`.

```sh
make build
```

### Image making

Below command makes Docker image:

```sh
make image
```

### Checks

Includes belows:

* `make test`:`go test`
* `make lint`: <https://github.com/golangci/golangci-lint>
* `make shellcheck`: <https://github.com/koalaman/shellcheck>
* `make mdlint`: <https://github.com/DavidAnson/markdownlint-cli2>

```sh
make check
```

### OpenAPI sleketons

Go source code generating can be executed by below command:

```sh
make openapi-server
```

target dirs:

* `internal/models`
* `internal/restapi`

## Improvements

Below chapters describe possible improvemts.

### Model

* More description (/id=0 at POST/PUT)
* gorm.Model ?
* More DB field tags (uniq, index, uniqueIndex, type, size, check)
* More OpenAPI field tags (min, max)
* OpenAPI optional fields (zero value?)
* Own Go field tags (complex rules)
* user authentication and authorization
* OpenAPI 3 (?)
* Integrating with a HTTP framework (Gin, Echo, …)
  * <https://github.com/mikkeloscar/gin-swagger>

### Source code

* Log libraries (for centralized log collector/parser)
  * <https://github.com/pgillich/errfmt/tree/emperror>
  * <https://github.com/logur/logur>
* Error handling library (emperror ?)
* Better HTML Status codes (Create: 201)
* More HTML Status error codes
* RFC5424, RFC7807
* separated endpoints, packages, dao for tags, ingredients and user + CRUD
* Better many-to-many handling (in transaction, use only ID or name fields)

### Development environment

* In-container build/deploy (Skaffold)

### Integration with other components

* Different repo: OpenAPI spec, generated client (Go, JS, etc)

### Repo

* main→master

### Test

* Postgres for `go test`, instead of SQLite (random db & user for each test suite/case)
* Automatic Unit/Function/Component tests (OpenAPI client)
* Mocks for unit tests ?
* Negative tests
* Table driven tests
  * <https://github.com/pgillich/date_calculator/blob/main/pkg/calendar/calendar_internal_test.go>
  * <https://github.com/pgillich/chat-bot/blob/master/pkg/frontend/frontend_test.go#L164>
* Automatic Integration, System, Stability, Stress tests
  * <https://github.com/pgillich/chat-bot/blob/master/pkg/frontend/frontend_test.go#L25>
* Automatic security checks
  * <https://github.com/pgillich/sample-blog/security/dependabot>
* Postman examples

Build

* Common Makefile + scripts in a new repo
* Go builder image, including all needed tools (build, linter)
* Deployment image, including all needed tool (helm, kubectl)
* Check: `go.mod` changed during build, `go mod tidy`
* To `/version` endpoint: `go mod graph`

### Deployments

* sensitive info in Sealed Secrets
* Makefile + docker for all actions (helm, kustomize)
* Pod resources
* GitOps (Flux, ArgoCD)
* Supporting more clusters easily
* Helm repository
* Autofill default records from file (ConfigMap, Secret or SealedSecret) or from Configuration Manager
* Get all config values from Configuration Manager

### CI

* Integrating to a CI (GitHub CI, Travis CI)
  * <https://github.com/pgillich/date_calculator/blob/main/.travis.yml>
  * <https://github.com/pgillich/grafana-tree-panel/tree/main/.github/workflows>
* Checks, image build & push
* Automatic master→developer git merge
* Start E2E tests

### Observability

* `/pprof` endpoint
* `/metrics` Prometheus endpoint
* Prometheus Operator PodMonitor, Grafana chart
* Loki (Promtail)
* OpenTracing (Jaeger)
* Alert rules and targets

### HA, Scaling, Zero Downtime, Rollback

* GORM AutoMigration, replicaCount: 1
* NBC DB schema change
* autoscaling.enabled = true
* backup/restore
