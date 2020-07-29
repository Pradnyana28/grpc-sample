package services_test

import (
	"testing"

	"github.com/Pradnyana28/uploads/seed"
	"github.com/Pradnyana28/uploads/services"
	"github.com/Pradnyana28/uploads/uploadpb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoId := seed.NewLaptop()
	laptopNoId.Id = ""

	laptopInvalidId := seed.NewLaptop()
	laptopInvalidId.Id = "invalid-id"

	laptopDuplicateId := seed.NewLaptop()
	storeDuplicateId := services.NewInMemoryLaptopStore()
	err := storeDuplicateId.Save(laptopDuplicateId)
	require.Nil(t, err)

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
		{
			name: "failure_duplicate_id",
			laptop: laptopDuplicateId,
			store: storeDuplicateId,
			code: codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func (t *testing.T) {
			t.Parallel()

			req := &uploadpb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := services.NewLaptopServer(tc.store)
			
		})
	}
}