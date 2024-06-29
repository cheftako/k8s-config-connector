// Copyright 2024 Google LLC
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

// +mockgcp-support
// apiVersion: kms.cnrm.cloud.google.com/v1beta1
// kind: KMSCryptoKey
// service: google.cloud.kms.v1.KeyManagementService
// resource: CryptoKey

package mockkms

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/kms/v1"
)

func (r *kmsServer) GetCryptoKey(ctx context.Context, req *pb.GetCryptoKeyRequest) (*pb.CryptoKey, error) {
	name, err := r.parseCryptoKeyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CryptoKey{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "CryptoKey %s not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *kmsServer) CreateCryptoKey(ctx context.Context, req *pb.CreateCryptoKeyRequest) (*pb.CryptoKey, error) {
	reqName := fmt.Sprintf("%s/cryptoKeys/%s", req.GetParent(), req.GetCryptoKeyId())
	name, err := r.parseCryptoKeyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetCryptoKey()).(*pb.CryptoKey)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	r.populateDefaultsForCryptoKey(name, obj)

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *kmsServer) populateDefaultsForCryptoKey(name *CryptoKeyName, obj *pb.CryptoKey) {

}

type CryptoKeyName struct {
	Project  *projects.ProjectData
	Location string
	KeyRing  string
	Name     string
}

func (n *CryptoKeyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/keyRings/" + n.KeyRing + "/cryptoKeys/" + n.Name
}

// parseCryptoKeyName parses a string into an CryptoKeyName.
// The expected form is `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
func (r *kmsServer) parseCryptoKeyName(name string) (*CryptoKeyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" && tokens[6] == "cryptoKeys" {
		project, err := r.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &CryptoKeyName{
			Project:  project,
			Location: tokens[3],
			KeyRing:  tokens[5],
			Name:     tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
