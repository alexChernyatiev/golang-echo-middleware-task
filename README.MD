### Build a middleware using echo framework

1. Create a handler which sends how many days left until 1 Jan 2025 and response with HTTP 200 OK status code.

2. Build a middleware, which checks HTTP header `User-Role` presents and contains `admin` and prints `red button user detected` to the console (using default log package or any 3rd party) if so.

## Run

```shell
go run github.com/alexChernyatiev/golang-echo-middleware-task
```

## Test

```shell
curl --location --request GET '127.0.0.1:8080/status' \
--header 'User-Role: admin'
```