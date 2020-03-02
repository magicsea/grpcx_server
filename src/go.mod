module github.com/magicsea/grpcx_server/src

go 1.13

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.4
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/magicsea/grpcx_server/src/share v0.0.0-20200229074005-dfbfb1ed90e4
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b
	google.golang.org/grpc v1.27.1
	pb v0.0.0-00010101000000-000000000000
	share v0.0.0-00010101000000-000000000000
)

replace (
	pb => ../pb
	share => ./share
)
