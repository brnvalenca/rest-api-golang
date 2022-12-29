create_grpc_service:

	protoc -I ./proto \
	--go_out ./proto/pb --go_opt paths=source_relative \
	--go-grpc_out ./proto/pb --go-grpc_opt paths=source_relative \
	./proto/grpc_services.proto

	protoc -I ./proto \
  --go_out ./proto/pb --go_opt paths=source_relative \
  --go-grpc_out ./proto/pb --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./proto/pb --grpc-gateway_opt paths=source_relative \
  ./proto/grpc_services.proto

	protoc -I ./proto \
		--grpc-gateway_out ./proto/pb \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out logtostderr=true:proto \
		proto/grpc_services.proto

	# Execute the custom struct field tag injection
	protoc-go-inject-tag -input="proto/pb/grpc_services.pb.go"

	[ $(which yq) ] || GO111MODULE=on go install github.com/mikefarah/yq/v4
	yq eval -P proto/grpc_services.swagger.json > swagger/openapi.yaml
	rm ./proto/grpc_services.swagger.json

go-run:
	go run main.go

docker-image:
	docker build . -t grpc-api-server

docker-run:
	docker run -dp 8080:8080 -p 9090:9090 grpc-api-server

build:
	docker-compose build
up:
	docker-compose up
down:
	docker-compose down