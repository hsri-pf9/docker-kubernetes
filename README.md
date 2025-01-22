## **Project Overview**
This project demonstrates the use of **Docker** and **Kubernetes** to deploy services for two different languages: **Python** and **Go**. The project includes:
1. A simple **Hello World Flask app** running inside a Docker container.
2. A **Ping-Pong** service where:
   - **Python** and **Go** both respond to the `/ping` endpoint with a "pong" message.
3. A **combined Docker setup** running both Python and Go services on different ports.
---
### 1. Setting up of Docker
The docker setup was already done in my pc so I proceeded to the next task.
### 2. Hello Word
#### **Overview**
The **Hello World** Flask app demonstrates a simple web application running inside a Docker container. The app listens for HTTP requests on port `5000` and returns a message `"Hello, Docker!"`. This example showcases how to build a basic Python web app and containerize it using Docker, followed by deploying it to a Kubernetes cluster
In this I have implied two steps:-
#### A) Directly pulling the Hello World Container from the Docker Hub
We can directly write the below command in the terminal <br />
```% docker run hello-world```
#### B) Writing all the required files for the Docker image and deploying it on Kubernetes
In the ```hello-world-docker``` folder we have all the folders to build a docker image and deploying it on the kubernetes <br />
1.) The first file which is ```app.py```. It contains the Flask code for the server on port 5000 which will show Hello, Docker! on the screen.<br />
2.) The second file which is the ```Dockerfile```. It is designed to containerize a Python application build with Flask.<br />
* Base Image
  - The Dockerfile uses the ```python:3.9-slim``` image as the base
* Working Directory
  - The working directory inside the container is set to ```/app```.
* Copy Application Files
  - The ```app.py``` is copied into the container
* Install Dependencies
  - The ```pip install flask``` command install Flask
* Expose port
  - The port 5000 is exposed
* Run the application
  -  The application is started with the command ```python app.py```<br />
  
3.) The third file is ```deployment.yaml``` file. This file is used to define and deploy the application to a Kubernetes cluster. <br />
* API Version
  - The file uses ```apps/v1```, the stable API version for deploying workloads in Kubernetes.
* Kind
  - The kind is ```Deployment```, which ensures that the desired number of application replicas are always running.
* Metadata
  - ```name```: Specifies the name of the deployment (```hello-world-deployment```).
* Spec
  - Replicas: Sets the number of application instances to 1.
  - Selector: Matches Pods with the label ```app: hello-world```.
  - Template:
    - Metadata: Defines labels applied to the Pods created by this deployment.
    - Spec:
        - Containers:<br />
            -```name```: The container's name is ```hello-world```.<br />
            -```image```: The Docker image to use, specified as ```hello-world-docker```.<br />
            -```port```: Exposes port 5000 on the container for external access.<br /><br />
            
4.) The fourth file is ```service.yaml``` file. This file is used to expose the application to external traffic in the Kubernetes cluster.
* API Version
  - The file uses ```apps/v1```, the stable API version for deploying workloads in Kubernetes.
* Kind
  - The kind is ```Service```, which provides network access to a set of Pods.
* Metadata
  - ```name```: The name of the service is ```hello-world-service```.
* Spec
  - Type: The service type is LoadBalancer, which exposes the application to external traffic through a cloud provider's load balancer.
  - Selector: The service routes traffic to Pods with the label ```app: hello-world```.
  - Ports
    - protocol: TCP
    - port: The service listens on port 5000.
    - targetPort: Traffic received on port 5000 is forwarded to the application running on the same port in the Pod.<br /><br />
#### Steps to write on terminal
1. ```docker build -t hello-world-docker:latest .```
2. ```docker login```<br /> ```docker tag hello-world-docker:latest <your-dockerhub-username>/hello-world-docker:latest```<br /> ```docker push <your-dockerhub-username>/hello-world-docker:latest```
3. ```kubectl apply -f deployment.yaml```<br /> ```kubectl apply -f service.yaml```
4. Now we should be able to access through (http://localhost:5000) on the browser.
5. This will expose the Flask app on port `5000` locally. You can access the application in your browser or using `curl`: ```curl http://localhost:5000```

You should receive a response:
```Hello, Docker!```

### 3. Ping-Pong Server
#### **Overview**
The **Ping-Pong service** provides two separate services:
1. A **Python Flask service** that responds with `"pong"` to the `/ping` endpoint.
2. A **Go service** that responds with `"pong"` to the `/ping` endpoint.These services are built and containerized separately, using individual Dockerfiles for Python and Go, and then deployed to Kubernetes.
3. A **Python and Go service** that responds with `"pong from Python"` and `"pong from Go"` to the endpoint `/ping` on two different ports.<br />

The Python setup is in the folder named ```ping-pong-project```, the Go setup is in the folder named ```ping-pong-project-go``` and the combined setup is in the folder ```ping-pong```.
The main structure of all the these folders are:-<br />
* `server` folder
  - `server.py` or `main.go` file is there in which the code for the server is there in python and go respectively.
  - `Dockerfile` is there which is designed to containerize the application.
* `k8s` folder
  - `k8s-deployment.yaml` file is there which contains the code for deployment and service of the kubernetes.
* `Makefile` which is to automate repetitive tasks and to simplify the process of building, managing, and deploying software projects.
  - Build the Go server (```make build-go```).
  - Build the Docker image (```make build-docker-go```).
  - Push the Docker image to the registry (```make push-docker-go```).
  - Deploy the app to Kubernetes (```make deploy```).
#### Commands to write on terminal
1. ```make build-python``` or ```make build-go``` based on which folder you are accessing. For the combined Python and Go it will be ```make build-python``` then ```make build-go```.
2. ```make build-docker-python``` or ```make build-docker-go``` based on your choice.
3. ```make push-docker-python``` or ```make push-docker-go``` based on your choice. For combined ```make push-python``` then ```make push-go```.
4. ```make deploy```<br />

Now based on the ports defined on running `curl`: ```curl http://localhost:port``` we will get the output as `pong`.
