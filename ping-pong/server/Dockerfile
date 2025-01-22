# Stage argument to differentiate builds
ARG APP=python

# Base image for Python
FROM python:3.9 AS python
WORKDIR /app
COPY python_server.py .
RUN pip install flask
CMD ["python", "python_server.py"]

# Base image for Go
FROM golang:1.19 AS go
WORKDIR /app
COPY go_server.go .
RUN go build -o go_server go_server.go
CMD ["./go_server"]

# Final stage: Copy from appropriate stage based on APP argument
FROM $APP
