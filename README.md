# gRPC Poc

## Install protoc and googleapis

```bash
brew install protobuf

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

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

### query traceid

```
{service_name=~".+"} | json | trace_id="<trace_id>"
```
