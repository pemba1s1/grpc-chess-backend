generate_grpc_code:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	chess/chess.proto

build_docker:
	docker build -t grpc-go-chess-server .

run_docker:
	docker run -p 8080:8080 grpc-go-chess-server