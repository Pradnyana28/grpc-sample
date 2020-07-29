package services

import (
	"context"
	"errors"
	"log"

	"github.com/Pradnyana28/uploads/uploadpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer ...
type LaptopServer struct {
	Store LaptopStore
}

// NewLaptopServer ...
func NewLaptopServer(s LaptopStore) *LaptopServer {
	return &LaptopServer{s}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *uploadpb.CreateLaptopRequest,
) (*uploadpb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// check if it's a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	// save laptop to database / in-memory storage
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to the store: %w", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)
	
	res := &uploadpb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}