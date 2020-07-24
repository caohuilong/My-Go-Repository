module grpc_mysql

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.23.0
	grpc_mysql/api/proto v0.0.0
)

replace grpc_mysql/api/proto => ./
