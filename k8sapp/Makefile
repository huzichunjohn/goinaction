PROJECT=goinaction/k8sapp
APP=k8sapp

RELEASE?=0.0.2
COMMIT=$(shell git rev-parse --short HEAD)
REPO_INFO=$(shell git config --get remote.origin.url)

CONTAINER_IMAGE=k8sapp
CONTAINER_NAME=k8sapp

GOOS?=linux
GOARCH?=amd64

K8SAPP_LOCAL_HOST?=0.0.0.0
K8SAPP_LOCAL_PORT?=8080
K8SAPP_LOG_LEVEL?=0

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -a -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/version.RELEASE=${RELEASE} -X ${PROJECT}/version.COMMIT=${COMMIT} -X ${PROJECT}/version.REPO=${REPO_INFO}" \
		-o bin/${GOOS}-${GOARCH}/${APP} ${PROJECT}/cmd
	docker build -t ${CONTAINER_IMAGE}:${RELEASE} .

.PHONY: run
run: build
	docker run --name ${CONTAINER_NAME} -p ${K8SAPP_LOCAL_PORT}:${K8SAPP_LOCAL_PORT} \
		-e "K8SAPP_LOCAL_HOST=${K8SAPP_LOCAL_HOST}" \
		-e "K8SAPP_LOCAL_PORT=${K8SAPP_LOCAL_PORT}" \
		-e "K8SAPP_LOG_LEVEL=${K8SAPP_LOG_LEVEL}" \
		-d ${CONTAINER_IMAGE}:${RELEASE}
	sleep 1
	docker logs ${CONTAINER_NAME}

HAS_RUNNED := $(shell docker ps|grep ${CONTAINER_NAME})
HAS_EXITED := $(shell docker ps -a|grep ${CONTAINER_NAME})

.PHONY: logs
logs:
	docker logs ${CONTAINER_NAME}

.PHONY: stop
stop:
ifdef HAS_RUNNED
	docker stop ${CONTAINER_NAME}
endif

.PHONY: start
start: stop
	docker start ${CONTAINER_NAME}

.PHONY: rm
rm:
ifdef HAS_EXITED
	docker rm ${CONTAINER_NAME}
endif

.PHONY: clean
clean: stop rm
	rm -f bin/${GOOS}-${GOARCH}/${APP}
