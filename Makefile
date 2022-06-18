# this command will push docker image
docker.push:
	docker tag todo-service firdausalif/todo-service
	docker push firdausalif/todo-service

docker.build.push:
	docker build --tag todo-service .
	docker tag todo-service firdausalif/todo-service
	docker push firdausalif/todo-service
