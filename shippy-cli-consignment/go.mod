module github.com/quen2404/shippy/shippy-cli-consignment

replace github.com/quen2404/shippy/shippy-service-consignment => ../shippy-service-consignment

go 1.14

require (
	github.com/quen2404/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.30.0
)
