# Simple Go with kubernetes example

This repo shows a simple walk through how to build and deploy service with http interface to kubernetes. Starting with the hello world example in the first commit and ending up with small REST API with benchmarks and test cases with data race detection.

## Requirements

- [Minikube](https://kubernetes.io/docs/getting-started-guides/minikube/) running
- [Docker client configured to use minikube](https://kubernetes.io/docs/getting-started-guides/minikube/#reusing-the-docker-daemon)
- Go and make installed

## Workflow

### Building

Edit your code and run
```sh
make docker-build
```
This will run docker build command and bump version if build was successful.
If you check the Dockerfile you can se that we are using [multi-stage build](https://docs.docker.com/develop/develop-images/multistage-build/).

In first stage we use `golang:1.10` to test and build the binary and then copy that to actual image which is build from `scratch`.

### Deploying

Run
```sh
make kube-deploy
```

This command will call `kubectl apply` with our `deployment.yaml` file and replace the version in deployment.yaml to match one in `VERSION`

After the first deployment it is required to expose our app by running
```sh
make kube-expose
```

Now the app should be accessible and minikube should give the url for it
```sh
minikube service demo-app --url
```

### Testing and benchmarking

To run all tests locally with race detection and code coverage run
```sh
make benchmark
```

To run all bencmarks run
```sh
make test
```
