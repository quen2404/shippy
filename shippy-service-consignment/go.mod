module github.com/quen2404/shippy/shippy-service-consignment

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible