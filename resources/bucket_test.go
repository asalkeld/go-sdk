// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"

	mock_v1 "github.com/nitrictech/go-sdk/mocks"
	"github.com/nitrictech/go-sdk/mocks/mockapi"
	nitricv1 "github.com/nitrictech/go-sdk/nitric/v1"
)

var _ = Describe("bucket", func() {
	ctrl := gomock.NewController(GinkgoT())
	Context("New", func() {
		mockConn := mock_v1.NewMockClientConnInterface(ctrl)
		When("valid args", func() {
			mockClient := mock_v1.NewMockResourceServiceClient(ctrl)
			mockStorage := mockapi.NewMockStorage(ctrl)

			m := &manager{
				blockers: map[string]Starter{},
				conn:     mockConn,
				rsc:      mockClient,
				storage:  mockStorage,
			}

			mockClient.EXPECT().Declare(context.Background(),
				&nitricv1.ResourceDeclareRequest{
					Resource: &nitricv1.Resource{
						Type: nitricv1.ResourceType_Bucket,
						Name: "red",
					},
					Config: &nitricv1.ResourceDeclareRequest_Bucket{
						Bucket: &nitricv1.BucketResource{},
					},
				})

			mockClient.EXPECT().Declare(context.Background(),
				&nitricv1.ResourceDeclareRequest{
					Resource: &nitricv1.Resource{
						Type: nitricv1.ResourceType_Policy,
					},
					Config: &nitricv1.ResourceDeclareRequest_Policy{
						Policy: &nitricv1.PolicyResource{
							Principals: []*nitricv1.Resource{{
								Type: nitricv1.ResourceType_Function,
							}},
							Actions: []nitricv1.Action{
								nitricv1.Action_BucketFileGet, nitricv1.Action_BucketFileList, nitricv1.Action_BucketFilePut,
							},
							Resources: []*nitricv1.Resource{{
								Type: nitricv1.ResourceType_Bucket,
								Name: "red",
							}},
						},
					},
				})

			mockBucket := mockapi.NewMockBucket(ctrl)
			mockStorage.EXPECT().Bucket("red").Return(mockBucket)
			b, err := m.NewBucket("red", BucketReading, BucketWriting)

			It("should not return an error", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(b).ShouldNot(BeNil())
			})
		})
	})
})
