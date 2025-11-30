# gRPC Poc

## Generate gRPC code

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/service.proto
```

## test with curl

```bash
curl "http://localhost:8080/api/data?id=123"
```
