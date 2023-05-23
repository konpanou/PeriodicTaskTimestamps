# Periodic Task Timestamps

This is an implementation of a JSON/HTTP service that returns all matching timestamps of a periodic task between two points in time.

## Building and running the project

The solution has been developed using Golang version 1.20. 

Before building the application we need to download any missing dependencies. We can do that by issuing the following command in the project root directory:

```
go mod download
```

In order to run the application we can issue the following command from the project root directory:

```
go build -o ./periodictimestamps
```
The application expects the host address and port to be specified at the application startup. In order to run the service listening on all interfaces on port 8000 we can issue the following command:

```azure
./periodictimestamps 0.0.0.0 8000
```

## Running the project using Docker

The project includes a Dockerfile and a docker-compose.yml file for building and running the application using docker.

In order to build and run the application using docker we can issue the following command:

```
docker-compose up
```

The first time the command is run it will build the docker image for the service and then run the container using port 8000. Note that the port needs to be free in order for the application to work.

## Tests

The project includes unit tests for the function that calculates the periodic timestamps which can be found in the api/timestamps_test.go file.

## Example requests

Following are some example api requests:

- Request periodic timestamps for period day ("1d")

```
curl "http://localhost:8000/ptlist?period=1d&tz=Europe/Athens&t1=20210214T204603Z&t2=20210217T123456Z"
```

- Request periodic timestamps for period month ("1mo")

```
curl "http://localhost:8000/ptlist?period=1mo&tz=Europe/Athens&t1=20210214T204603Z&t2=20211115T123456Z"
```
- Request periodic timestamps for period hour ("1h")

```
curl "http://localhost:8000/ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z"
```

- Request periodic timestamps for period year ("1y")

```
curl "http://localhost:8000/ptlist?period=1y&tz=Europe/Athens&t1=20210214T204603Z&t2=20231115T123456Z"
```
