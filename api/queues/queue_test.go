package queues

import (
	"fmt"

	"github.com/golang/mock/gomock"
	v1 "github.com/nitrictech/go-sdk/interfaces/nitric/v1"
	mock_v1 "github.com/nitrictech/go-sdk/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ = Describe("Queue", func() {
	ctrl := gomock.NewController(GinkgoT())

	Context("Send", func() {
		When("the gRPC server returns an error", func() {
			mockQ := mock_v1.NewMockQueueClient(ctrl)

			mockQ.EXPECT().SendBatch(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("mock error"))

			q := &queueImpl{
				name: "test-queue",
				c:    mockQ,
			}

			_, err := q.Send([]*Task{
				&Task{
					ID:          "1234",
					PayloadType: "test-payload",
					Payload: map[string]interface{}{
						"test": "test",
					},
				},
			})

			It("should pass through the error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("mock error"))
			})
		})

		When("the task send succeeds", func() {
			mockQ := mock_v1.NewMockQueueClient(ctrl)
			mockStruct, _ := structpb.NewStruct(map[string]interface{}{
				"test": "test",
			})

			mockQ.EXPECT().SendBatch(gomock.Any(), gomock.Any()).Return(&v1.QueueSendBatchResponse{
				FailedTasks: []*v1.FailedTask{
					&v1.FailedTask{
						Message: "Failed to send task",
						Task: &v1.NitricTask{
							Id:          "1234",
							PayloadType: "test-payload",
							Payload:     mockStruct,
						},
					},
				},
			}, nil)

			q := &queueImpl{
				name: "test-queue",
				c:    mockQ,
			}

			fts, _ := q.Send([]*Task{
				&Task{
					ID:          "1234",
					PayloadType: "test-payload",
					Payload: map[string]interface{}{
						"test": "test",
					},
				},
			})

			It("should recieve the failed tasks from the QueueSendBatchResponse", func() {
				Expect(fts).To(HaveLen(1))
				Expect(fts[0].Reason).To(Equal("Failed to send task"))
				Expect(fts[0].Task.ID).To(Equal("1234"))
				Expect(fts[0].Task.PayloadType).To(Equal("test-payload"))
				Expect(fts[0].Task.Payload).To(Equal(map[string]interface{}{
					"test": "test",
				}))
			})
		})

		Context("Receive", func() {
			When("Retrieving tasks with depth less than 1", func() {
				q := &queueImpl{
					name: "test-queue",
				}

				_, err := q.Receive(0)

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("Depth cannot be less than 1"))
				})
			})

			When("The grpc successfully returns", func() {
				mockStruct, _ := structpb.NewStruct(map[string]interface{}{
					"test": "test",
				})
				mockQ := mock_v1.NewMockQueueClient(ctrl)

				mockQ.EXPECT().Receive(gomock.Any(), gomock.Any()).Return(&v1.QueueReceiveResponse{
					Tasks: []*v1.NitricTask{
						&v1.NitricTask{
							Id:          "1234",
							Payload:     mockStruct,
							PayloadType: "mock-payload",
							LeaseId:     "1234",
						},
					},
				}, nil)

				q := &queueImpl{
					name: "test-queue",
					c:    mockQ,
				}

				t, _ := q.Receive(1)

				It("should recieve a single task", func() {
					Expect(t).To(HaveLen(1))
				})

				rt, ok := t[0].(*receivedTaskImpl)

				It("the task should be of type recieveTaskImpl", func() {
					Expect(ok).To(BeTrue())
				})

				It("Should contain the returned task", func() {
					tsk := rt.Task()

					Expect(tsk.ID).To(Equal("1234"))
					Expect(tsk.PayloadType).To(Equal("mock-payload"))
					Expect(tsk.Payload).To(Equal(map[string]interface{}{
						"test": "test",
					}))
				})
			})
		})
	})
})