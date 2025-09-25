.PHONY: gen

gen:
	@protoc \
		-I api \
		-I third_party \
		--go_out=services/common/gen/auth --go_opt=paths=source_relative \
		--go-grpc_out=services/common/gen/auth --go-grpc_opt=paths=source_relative \
		api/auth.proto

	@protoc \
		-I api \
		-I third_party \
		--go_out=services/common/gen/user --go_opt=paths=source_relative \
		--go-grpc_out=services/common/gen/user --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=services/common/gen/user --grpc-gateway_opt=paths=source_relative,logtostderr=true \
		api/user.proto


start-service:
	@go run ./services/auth/*.go &
	@go run ./services/user/*.go &
	@go run ./gateway/*.go &
	@echo "All services started in the background."

start-auth:
	@go run ./services/auth/*.go
start-user:
	@go run ./services/user/*.go
start-gate:
	@go run ./gateway/*.go
