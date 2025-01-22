# Stage argument to differentiate builds
ARG APP=python

FROM python:3.9 AS python
WORKDIR /app
COPY python_server.py .
RUN pip install flask
CMD ["python", "python_server.py"]


FROM golang:1.19 AS go
WORKDIR /app
COPY go_server.go .
RUN go build -o go_server go_server.go
CMD ["./go_server"]

FROM $APP
