package services_test

import (
	"testing"

	"github.com/Pradnyana28/uploads/seed"
	"github.com/Pradnyana28/uploads/services"
	"github.com/Pradnyana28/uploads/uploadpb"
	"google.golang.org/grpc/codes"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoId := seed.NewLaptop()
	laptopNoId.Id = ""

	laptopFailedId := seed.NewLaptop()
	laptopFailedId.Id = "asd976a8s7d6a9s"

	testCases := []struct {
		name string
		laptop *uploadpb.Laptop
		store services.LaptopStore
		code codes.Code
	}{
		{
			name: "success_with_id",
			laptop: seed.NewLaptop(),
			store: services.NewInMemoryLaptopStore(),
			code: codes.OK,
		},
		{
			name: "success_no_id",
			laptop: laptopNoId,
			store: services.NewInMemoryLaptopStore(),
			code: codes.InvalidArgument,
		},
		{
			name: "failure_invalid_id",
			laptop: laptopInvalidId,
			store: services.NewInMemoryLaptopStore(),
			code: codes.InvalidArgument,
		},
	}
}