docker/run:
	docker container rm keywordss-api-dev
	docker run -d  --name keywordss-api-dev -p 8080:8080 -v `pwd`/api:/go/src/keywordss/api tea1013/keywordss-api:latest make run
	@echo 'connect port :8080 !!!'

docker/stop:
	docker container stop keywordss-api-dev

image/pull:
	docker pull tea1013/keywordss-api

image/rm:
	docker image rm -f tea1013/keywordss-api
