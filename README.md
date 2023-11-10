# Simple Microservice Design (no-database included)

This project is a simple microservice design which mocks an API call for fetching a static price

## Usage

- use this code in your terminal to run the server instance:

```bash
make server
```

- you can run your client on top of the server by running this:

```bash
make client
```

# Containerization

- Simple run this code to create your own container and run the code:

```bash
# use sudo if needed
# you can use `docker run your_container_name -d` to run the container in detached mode`

docker build -t your_container_name .
docker run your_container_name
```