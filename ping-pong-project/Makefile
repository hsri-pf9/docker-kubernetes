.PHONY: build-python run-python build-docker-python push-docker-python deploy

build-python:
	python3 python-server/server.py

build-docker-python:
	docker build -t ping-pong-python ./python-server

push-docker-python:
	docker tag ping-pong-python hsri/ping-pong-python
	docker push hsri/ping-pong-python

deploy:
	kubectl apply -f k8s/k8s-deployment.yaml

