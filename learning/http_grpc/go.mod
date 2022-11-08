module http_grpc

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/jergoo/go-grpc-example v0.0.0-20200718045408-25c3f5e9b045 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
	http_grpc/api v0.0.0
)

replace http_grpc/api => ./
