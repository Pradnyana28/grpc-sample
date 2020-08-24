package services_test

import (
	"context"
	"net"
	"testing"

	"github.com/Pradnyana28/uploads/seed"
	"github.com/Pradnyana28/uploads/serializer"
	"github.com/Pradnyana28/uploads/services"
	"github.com/Pradnyana28/uploads/uploadpb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := seed.NewLaptop()
	expectedId := laptop.Id
	req := &uploadpb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedId, res.Id)

	// check that the laptop is saved to the store
	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check that the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)
}

func startTestLaptopServer(t *testing.T) (*services.LaptopServer, string) {
	laptopServer := services.NewLaptopServer(services.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	uploadpb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0") // random available ports
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) uploadpb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return uploadpb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop *uploadpb.Laptop, other *uploadpb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(other)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
