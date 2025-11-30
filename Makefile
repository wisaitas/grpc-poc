.PHONY: rundomain runorchestrator rungateway

rundomain:
	go run cmd/domain/main.go

runorchestrator:
	go run cmd/orchestrator/main.go

rungateway:
	go run cmd/gateway/main.go