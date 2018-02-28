# API -service

This is a demo service which is used in the presentation and setup of
[https://github.com/jerryjj/k8s-gke-deployment-tpl](Kubernetes - Happily in production).

## Development

**Get dependencies**

```sh
go get -u google.golang.org/grpc
go get -u github.com/op/go-logging
go get -u github.com/gorilla/mux
```

### Building

```sh
./build.sh
```

**Building and Deploying to the GCP Project**

```sh
export PROJECT_ID=[YOUR_GCP_PROJECT_ID]
./cloud-build.sh
```

### Running

Run the payments-service in another terminal.

```sh
./api-service
```

Do a GET request to http://localhost:8080/payments/status
