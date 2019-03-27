API_REPOSITORY_NAME:=hatchgroup/keywordss-api
API_CONTAINER_NAME:=keywordss-api-dev
DB_REPOSITORY_NAME:=hatchgroup/keywordss-db
DB_CONTAINER_NAME:=keywordss-db-dev

HOST_APP_BASE:=$(shell pwd)/api
DOCKER_APP_BASE:=/go/src/github.com/hatch-group/keywordss-api/api
DB_VOLUME_PATH:=$(shell pwd)/_secret/keywordss-data

docker/run:
	$(MAKE) docker/run/server
	$(MAKE) docker/run/db

docker/run/server:
	docker run -d --name $(API_CONTAINER_NAME) --env-file _secret/.env -p 8080:8080 -v $(HOST_APP_BASE):$(DOCKER_APP_BASE) $(API_REPOSITORY_NAME):latest
	@echo 'connect port :8080 !!!'

docker/run/db:
	docker run -d -p 3306:3306 --name $(DB_CONTAINER_NAME) -v $(DB_VOLUME_PATH):/var/lib/mysql --env-file _secret/.env $(DB_REPOSITORY_NAME):latest
	@echo 'Connect DB port :3306 !!!'

docker/stop:
	$(MAKE) docker/stop/api
	$(MAKE) docker/stop/db

docker/stop/api:
	docker container stop $(API_CONTAINER_NAME)
	docker container rm $(API_CONTAINER_NAME)

docker/stop/db:
	docker container stop $(DB_CONTAINER_NAME)
	docker container rm $(DB_CONTAINER_NAME)

db/init:
	mysql -u root -p -h 127.0.0.1 --port 3306 -e 'CREATE DATABASE keywordss;'
	mysql -u root -p -h 127.0.0.1 --port 3306 keywordss < sql-files/init.sql

db/delete:
	mysql -u root -p -h 127.0.0.1 --port 3306 -e 'DROP DATABASE keywordss;'

image/build:
	docker build -t $(API_REPOSITORY_NAME) .

image/push:
	docker push $(API_REPOSITORY_NAME)

image/pull:
	docker pull $(API_REPOSITORY_NAME)

image/rm:
	docker image rm -f $(API_REPOSITORY_NAME)
