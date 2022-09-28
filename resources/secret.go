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
	"fmt"

	nitricv1 "github.com/nitrictech/apis/go/nitric/v1"
	"github.com/nitrictech/go-sdk/api/secrets"
)

type SecretPermission string

const (
	SecretReading SecretPermission = "reading"
	SecretWriting SecretPermission = "writing"
)

var SecretEverything []SecretPermission = []SecretPermission{SecretReading, SecretWriting}

func NewSecret(name string, permissions ...SecretPermission) (secrets.SecretRef, error) {
	return run.NewSecret(name, permissions...)
}

func (m *manager) NewSecret(name string, permissions ...SecretPermission) (secrets.SecretRef, error) {
	rsc, err := m.resourceServiceClient()
	if err != nil {
		return nil, err
	}

	colRes := &nitricv1.Resource{
		Type: nitricv1.ResourceType_Secret,
		Name: name,
	}

	dr := &nitricv1.ResourceDeclareRequest{
		Resource: colRes,
		Config: &nitricv1.ResourceDeclareRequest_Secret{
			Secret: &nitricv1.SecretResource{},
		},
	}
	_, err = rsc.Declare(context.Background(), dr)
	if err != nil {
		return nil, err
	}

	actions := []nitricv1.Action{}
	for _, perm := range permissions {
		switch perm {
		case SecretReading:
			actions = append(actions, nitricv1.Action_SecretAccess)
		case SecretWriting:
			actions = append(actions, nitricv1.Action_SecretPut)
		default:
			return nil, fmt.Errorf("secretPermission %s unknown", perm)
		}
	}

	_, err = rsc.Declare(context.Background(), functionResourceDeclareRequest(colRes, actions))
	if err != nil {
		return nil, err
	}

	if m.secrets == nil {
		m.secrets, err = secrets.New()
		if err != nil {
			return nil, err
		}
	}

	return m.secrets.Secret(name), nil
}
