docker/run:
	docker run -d  --name keywordss-api-dev -p 8080:8080 -v `pwd`/api:/go/src/keywordss/api hatchgroup/keywordss-api:latest make run
	@echo 'connect port :8080 !!!'

docker/stop:
	docker container stop keywordss-api-dev
	docker container rm keywordss-api-dev

image/push:
	docker push hatchgroup/keywordss-api

image/pull:
	docker pull hatchgroup/keywordss-api

image/rm:
	docker image rm -f hatchgroup/keywordss-api
