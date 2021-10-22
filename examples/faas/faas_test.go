package faas_examples

import (
	"io"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	v1 "github.com/nitrictech/apis/go/nitric/v1"
	mock_v1 "github.com/nitrictech/go-sdk/mocks"
	"google.golang.org/grpc"
)

func newMockStream(ctrl *gomock.Controller, msgs []*v1.ServerMessage) func(stream v1.FaasService_TriggerStreamServer) error {
	return func(stream v1.FaasService_TriggerStreamServer) error {
		for _, m := range msgs {
			stream.Send(m)
		}

		return io.EOF
	}
}

func TestFaasSnippets(t *testing.T) {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// Create a mock storage service server...
	ctrl := gomock.NewController(t)
	ms := mock_v1.NewMockFaasServiceServer(ctrl)

	// Return the mock trigger stream when called
	gomock.InOrder(
		// Return the stream for the evts snippet test
		ms.EXPECT().TriggerStream(gomock.Any()).Do(newMockStream(ctrl, []*v1.ServerMessage{
			{
				Id: "1234",
				Content: &v1.ServerMessage_TriggerRequest{
					TriggerRequest: &v1.TriggerRequest{
						Data: []byte("{\"payload\": {\"test\": \"test\"}}"),
						Context: &v1.TriggerRequest_Topic{
							Topic: &v1.TopicTriggerContext{
								Topic: "mock-topic",
							},
						},
					},
				},
			},
		})).Return(nil).Times(1),
	)

	// Start the gRPC server with the mock instance and await for it
	// to be called
	lis, _ := net.Listen("tcp", ":50051")

	v1.RegisterFaasServiceServer(grpcServer, ms)
	go grpcServer.Serve(lis)
	// call the snippets to test
	evts()

	// Cleanup
	grpcServer.Stop()
	lis.Close()
	ctrl.Finish()
}
