FROM python:3.9-slim

WORKDIR /app
COPY server.py .
RUN pip install flask
EXPOSE 8081
CMD ["python", "server.py"]

