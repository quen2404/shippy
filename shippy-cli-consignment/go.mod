module github.com/quen2404/shippy/shippy-cli-consignment

replace github.com/quen2404/shippy/shippy-service-consignment => ../shippy-service-consignment

go 1.14

require (
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/quen2404/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.30.0 // indirect
	google.golang.org/grpc/examples v0.0.0-20200714235929-a6c3c6968e90 // indirect
)
