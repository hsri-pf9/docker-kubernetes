.PHONY: build-python build-go push-python push-go deploy

build-python:
	docker build --build-arg APP=python -t ping-pong-python:v2 ./server

build-go:
	docker build --build-arg APP=go -t ping-pong-go:v2 ./server

push-python:
	docker tag ping-pong-python:v2 hsri/ping-pong-python:v2
	docker push hsri/ping-pong-python:v2

push-go:
	docker tag ping-pong-go:v2 hsri/ping-pong-go:v2
	docker push hsri/ping-pong-go:v2

deploy:
	kubectl apply -f k8s/k8s-deployment.yaml
