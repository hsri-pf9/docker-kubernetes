.PHONY: build-python build-go push-python push-go deploy

build-python:
	docker build --build-arg APP=python -t ping-pong-python:v1 ./server

build-go:
	docker build --build-arg APP=go -t ping-pong-go:v1 ./server

push-python:
	docker tag ping-pong-python:v1 hsri/ping-pong-python:v1
	docker push hsri/ping-pong-python:v1

push-go:
	docker tag ping-pong-go:v1 hsri/ping-pong-go:v1
	docker push hsri/ping-pong-go:v1

deploy:
	kubectl apply -f k8s/k8s-deployment.yaml
