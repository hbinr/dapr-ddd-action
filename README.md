## Run
To run this demo in Dapr, run:
```bash
dapr run \
    --app-id http-service-demo \
    --app-port 8080 \
    --app-protocol http \
    --dapr-http-port 3500 \
    --components-path ./config \
    go run main.go
```

or run `make run`

- `dapr-study` study note