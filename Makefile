API_REPOSITORY_NAME:=hatchgroup/keywordss-api
API_CONTAINER_NAME:=keywordss-api-dev
HOST_APP_BASE:=$(shell pwd)/api
DOCKER_APP_BASE:=/go/src/github.com/hatch-group/keywordss-api/api

docker/run:
	docker run -d --name $(API_CONTAINER_NAME) --env-file _secret/.env -p 8080:8080 -v $(HOST_APP_BASE):$(DOCKER_APP_BASE) $(API_REPOSITORY_NAME):latest
	@echo 'connect port :8080 !!!'

docker/stop:
	docker container stop $(API_CONTAINER_NAME)
	docker container rm $(API_CONTAINER_NAME)

image/build:
	docker build -t $(API_REPOSITORY_NAME) .

image/push:
	docker push $(API_REPOSITORY_NAME)

image/pull:
	docker pull $(API_REPOSITORY_NAME)

image/rm:
	docker image rm -f $(API_REPOSITORY_NAME)
