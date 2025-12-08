.PHONY: rundomain runorchestrator rungateway test proto

rundomain:
	go run cmd/domain/main.go

runorchestrator:
	go run cmd/orchestrator/main.go

rungateway:
	go run cmd/gateway/main.go

test:
	curl "http://localhost:8080/api/data?id=123"

domainproto:
	mkdir -p internal/domain/pb/gen
	rm -f internal/domain/pb/gen/*.go
	rm -f internal/domain/pb/*.pb.go
	
	protoc --proto_path=. \
		--go_out=. --go_opt=module=github.com/wisaitas/grpc-poc \
		--go-grpc_out=. --go-grpc_opt=module=github.com/wisaitas/grpc-poc \
		internal/domain/pb/user.proto \
		internal/domain/pb/service.proto

orchestratorproto:
	mkdir -p internal/orchestrator/pb/gen
	rm -f internal/orchestrator/pb/gen/*.go
	rm -f internal/orchestrator/pb/*.pb.go
	
	protoc --proto_path=. \
		--go_out=. --go_opt=module=github.com/wisaitas/grpc-poc \
		--go-grpc_out=. --go-grpc_opt=module=github.com/wisaitas/grpc-poc \
		internal/orchestrator/pb/auth.proto \
		internal/orchestrator/pb/service.proto

proto:
	@find . -name "*.proto" -exec dirname {} \; | sort -u | while read dir; do \
		mkdir -p $$dir/gen; \
		rm -f $$dir/gen/*.go; \
		rm -f $$dir/*.pb.go; \
	done
	
	protoc --proto_path=. \
		--go_out=. --go_opt=module=github.com/wisaitas/grpc-poc \
		--go-grpc_out=. --go-grpc_opt=module=github.com/wisaitas/grpc-poc \
		$$(find . -name "*.proto")