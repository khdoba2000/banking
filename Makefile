CURRENT_DIR=$(shell pwd)

APP=banking
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
REGISTRY=khdoba

tidy:
	go mod tidy

# go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
shadow: 
	shadow ./...

vet:
	go fmt ./...
	go vet ./...
	# shadow ./...


build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go


build-image:
	docker build --rm -t ${REGISTRY}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${APP}:${TAG} ${REGISTRY}/${APP}:${TAG}


push-image:
	docker push ${REGISTRY}/${APP}:${TAG}
