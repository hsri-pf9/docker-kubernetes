.PHONY: build-go run-go build-docker-go push-docker-go deploy

build-go:
	go run go-server/main.go

build-docker-go:
	docker build -t ping-pong-go ./go-server

push-docker-go:
	docker tag ping-pong-go hsri/ping-pong-go
	docker push hsri/ping-pong-go

deploy:
	kubectl apply -f k8s/k8s-deployment.yaml
