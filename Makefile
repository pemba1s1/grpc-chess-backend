generate_grpc_code:
	protoc \
	--go_out=chess \
	--go_opt=paths=source_relative \
	--go-grpc_out=chess \
	--go-grpc_opt=paths=source_relative \
	chess.proto