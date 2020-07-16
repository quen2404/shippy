module github.com/quen2404/shippy/shippy-cli-user

go 1.14

require (
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/quen2404/shippy/shippy-service-user v0.0.0-20200716090243-9dc32c1012b0
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible