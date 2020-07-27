module github.com/magicsea/grpcx_server/src/share

go 1.13

require (
	github.com/fsnotify/fsnotify v1.4.7
	github.com/spf13/viper v1.6.2
	google.golang.org/grpc v1.27.1
	sd v0.0.0-00010101000000-000000000000
)

replace (
	pb => ../../pb
	sd => ../sd
)
