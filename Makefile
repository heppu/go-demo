.PHONY: test benchmark run docker-build docker-run

VERSION_FILE := VERSION
CURRENT_VERSION := $(shell cat ${VERSION_FILE})
NEW_VERSION := $(shell echo $$((${CURRENT_VERSION}+1)))

version:
	echo ${CURRENT_VERSION}

run:
	go run main.go

test:
	go test -v -race -cover ./...

benchmark:
	go test -benchmem -bench=. ./...

build: test
	env GOOS=linux CGO_ENABLED=0 go build -o app

docker-build:
	docker build --rm -t demo-app:${NEW_VERSION} .
	echo ${NEW_VERSION} > ${VERSION_FILE}

docker-run: docker-build
	docker run -p 8000:8000 demo-app:latest

kube-deploy:
	sed -e 's/VERSION/${CURRENT_VERSION}/g' deployment.yaml | kubectl apply -f -
	
kube-delete:
	sed -e 's/VERSION/${CURRENT_VERSION}/g' deployment.yaml | kubectl delete -f -
	
kube-expose:
	kubectl expose deployment demo-app --type=LoadBalancer

remove-docker-images:
	docker rmi $$(docker images 'demo-app' --format '{{.ID}}')
