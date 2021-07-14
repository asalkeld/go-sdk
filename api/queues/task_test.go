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

var _ = Describe("Task", func() {
	ctrl := gomock.NewController(GinkgoT())
	Context("taskToWire", func() {
		When("translating a task to a proto task", func() {
			t, _ := taskToWire(&Task{
				ID:          "1234",
				PayloadType: "test-payload",
				Payload: map[string]interface{}{
					"test": "test",
				},
			})

			It("should have the same ID", func() {
				Expect(t.GetId()).To(Equal("1234"))
			})

			It("should have the same payload type", func() {
				Expect(t.GetPayloadType()).To(Equal("test-payload"))
			})

			It("should translate the provided payload to a proto Struct", func() {
				Expect(t.GetPayload().AsMap()).To(Equal(
					map[string]interface{}{
						"test": "test",
					},
				))
			})
		})
	})

	Context("wireToTask", func() {
		When("translating a proto task to a task", func() {
			pbs, _ := structpb.NewStruct(map[string]interface{}{
				"test": "test",
			})
			t := wireToTask(&v1.NitricTask{
				Id:          "1234",
				PayloadType: "test-payload",
				Payload:     pbs,
			})

			It("should have the same ID", func() {
				Expect(t.ID).To(Equal("1234"))
			})

			It("should have the same payload type", func() {
				Expect(t.PayloadType).To(Equal("test-payload"))
			})

			It("should translate the provided struct into a map", func() {
				Expect(t.Payload).To(Equal(map[string]interface{}{
					"test": "test",
				}))
			})
		})
	})

	Context("receivedTaskImpl", func() {
		Context("Task", func() {
			When("Retrieving the inner task", func() {
				t := &Task{
					ID:          "1234",
					PayloadType: "test-payload",
					Payload: map[string]interface{}{
						"test": "test",
					},
				}
				rt := &receivedTaskImpl{task: t}

				It("should return a reference to the inner task", func() {
					Expect(rt.Task()).To(Equal(rt.task))
				})
			})
		})

		Context("Queue", func() {
			When("Retrieving the queue name", func() {
				rt := &receivedTaskImpl{queue: "test-queue"}

				It("should return the private queue field", func() {
					Expect(rt.Queue()).To(Equal(rt.queue))
				})
			})
		})

		Context("Complete", func() {
			When("the complete succeeds", func() {
				mockQc := mock_v1.NewMockQueueClient(ctrl)

				mockQc.EXPECT().Complete(gomock.Any(), gomock.Any()).Return(&v1.QueueCompleteResponse{}, nil)

				t := &receivedTaskImpl{
					qc:      mockQc,
					queue:   "test-queue",
					leaseId: "1234",
					task: &Task{
						ID:          "1234",
						PayloadType: "test-payload",
						Payload: map[string]interface{}{
							"test": "test",
						},
					},
				}

				err := t.Complete()

				It("should not return an error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			When("the gRPC server returns an error", func() {
				mockQc := mock_v1.NewMockQueueClient(ctrl)

				mockQc.EXPECT().Complete(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("mock error"))

				t := &receivedTaskImpl{
					qc:      mockQc,
					queue:   "test-queue",
					leaseId: "1234",
					task: &Task{
						ID:          "1234",
						PayloadType: "test-payload",
						Payload: map[string]interface{}{
							"test": "test",
						},
					},
				}

				err := t.Complete()

				It("should pass through the gRPC error", func() {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("mock error"))
				})
			})
		})
	})
})