module github.com/quen2404/shippy/shippy-cli-consignment

go 1.14

require (
	github.com/micro/go-micro/v2 v2.9.1
	github.com/quen2404/shippy/shippy-service-consignment v0.0.0-20200715124743-37e7d6a92e94
)

// replace github.com/quen2404/shippy/shippy-service-consignment => ../shippy-service-consignment
replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible