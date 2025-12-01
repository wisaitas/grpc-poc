.PHONY: rundomain runorchestrator rungateway test

rundomain:
	go run cmd/domain/main.go

runorchestrator:
	go run cmd/orchestrator/main.go

rungateway:
	go run cmd/gateway/main.go

test:
	curl "http://localhost:8080/api/data?id=123"