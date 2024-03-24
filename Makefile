PKG := "github.com/GeekQk/vblog"

dep: ## Get the dependencies
	@go mod tidy

run: ## Run Server
	@go run main.go start

gen: ## make protobuf
	@protoc -I=. --go_out=. --go_opt=module="github.com/GeekQk/vblog" --go-grpc_out="."  --go-grpc_opt=module="github.com/GeekQk/vblog"  apps/*/pb/*.proto 
	@protoc-go-inject-tag -input="apps/*/*.pb.go"
	
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
