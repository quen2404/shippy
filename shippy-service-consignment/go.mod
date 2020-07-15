module github.com/quen2404/shippy/shippy-service-consignment

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	github.com/quen2404/shippy/shippy-service-vessel v0.0.0-20200715132542-28c3721c8433
	go.mongodb.org/mongo-driver v1.3.5
	google.golang.org/protobuf v1.23.0
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible

// replace github.com/quen2404/shippy/shippy-service-vessel => ../shippy-service-vessel
