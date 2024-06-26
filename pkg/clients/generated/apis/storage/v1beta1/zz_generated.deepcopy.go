//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketAction) DeepCopyInto(out *BucketAction) {
	*out = *in
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketAction.
func (in *BucketAction) DeepCopy() *BucketAction {
	if in == nil {
		return nil
	}
	out := new(BucketAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketAutoclass) DeepCopyInto(out *BucketAutoclass) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketAutoclass.
func (in *BucketAutoclass) DeepCopy() *BucketAutoclass {
	if in == nil {
		return nil
	}
	out := new(BucketAutoclass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketCondition) DeepCopyInto(out *BucketCondition) {
	*out = *in
	if in.Age != nil {
		in, out := &in.Age, &out.Age
		*out = new(int64)
		**out = **in
	}
	if in.CreatedBefore != nil {
		in, out := &in.CreatedBefore, &out.CreatedBefore
		*out = new(string)
		**out = **in
	}
	if in.CustomTimeBefore != nil {
		in, out := &in.CustomTimeBefore, &out.CustomTimeBefore
		*out = new(string)
		**out = **in
	}
	if in.DaysSinceCustomTime != nil {
		in, out := &in.DaysSinceCustomTime, &out.DaysSinceCustomTime
		*out = new(int64)
		**out = **in
	}
	if in.DaysSinceNoncurrentTime != nil {
		in, out := &in.DaysSinceNoncurrentTime, &out.DaysSinceNoncurrentTime
		*out = new(int64)
		**out = **in
	}
	if in.MatchesPrefix != nil {
		in, out := &in.MatchesPrefix, &out.MatchesPrefix
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchesStorageClass != nil {
		in, out := &in.MatchesStorageClass, &out.MatchesStorageClass
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchesSuffix != nil {
		in, out := &in.MatchesSuffix, &out.MatchesSuffix
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NoncurrentTimeBefore != nil {
		in, out := &in.NoncurrentTimeBefore, &out.NoncurrentTimeBefore
		*out = new(string)
		**out = **in
	}
	if in.NumNewerVersions != nil {
		in, out := &in.NumNewerVersions, &out.NumNewerVersions
		*out = new(int64)
		**out = **in
	}
	if in.WithState != nil {
		in, out := &in.WithState, &out.WithState
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketCondition.
func (in *BucketCondition) DeepCopy() *BucketCondition {
	if in == nil {
		return nil
	}
	out := new(BucketCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketCors) DeepCopyInto(out *BucketCors) {
	*out = *in
	if in.MaxAgeSeconds != nil {
		in, out := &in.MaxAgeSeconds, &out.MaxAgeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResponseHeader != nil {
		in, out := &in.ResponseHeader, &out.ResponseHeader
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketCors.
func (in *BucketCors) DeepCopy() *BucketCors {
	if in == nil {
		return nil
	}
	out := new(BucketCors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketCustomPlacementConfig) DeepCopyInto(out *BucketCustomPlacementConfig) {
	*out = *in
	if in.DataLocations != nil {
		in, out := &in.DataLocations, &out.DataLocations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketCustomPlacementConfig.
func (in *BucketCustomPlacementConfig) DeepCopy() *BucketCustomPlacementConfig {
	if in == nil {
		return nil
	}
	out := new(BucketCustomPlacementConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketEncryption) DeepCopyInto(out *BucketEncryption) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketEncryption.
func (in *BucketEncryption) DeepCopy() *BucketEncryption {
	if in == nil {
		return nil
	}
	out := new(BucketEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleRule) DeepCopyInto(out *BucketLifecycleRule) {
	*out = *in
	in.Action.DeepCopyInto(&out.Action)
	in.Condition.DeepCopyInto(&out.Condition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleRule.
func (in *BucketLifecycleRule) DeepCopy() *BucketLifecycleRule {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLogging) DeepCopyInto(out *BucketLogging) {
	*out = *in
	if in.LogObjectPrefix != nil {
		in, out := &in.LogObjectPrefix, &out.LogObjectPrefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLogging.
func (in *BucketLogging) DeepCopy() *BucketLogging {
	if in == nil {
		return nil
	}
	out := new(BucketLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObservedStateStatus) DeepCopyInto(out *BucketObservedStateStatus) {
	*out = *in
	if in.SoftDeletePolicy != nil {
		in, out := &in.SoftDeletePolicy, &out.SoftDeletePolicy
		*out = new(BucketSoftDeletePolicyStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObservedStateStatus.
func (in *BucketObservedStateStatus) DeepCopy() *BucketObservedStateStatus {
	if in == nil {
		return nil
	}
	out := new(BucketObservedStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketRetentionPolicy) DeepCopyInto(out *BucketRetentionPolicy) {
	*out = *in
	if in.IsLocked != nil {
		in, out := &in.IsLocked, &out.IsLocked
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketRetentionPolicy.
func (in *BucketRetentionPolicy) DeepCopy() *BucketRetentionPolicy {
	if in == nil {
		return nil
	}
	out := new(BucketRetentionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSoftDeletePolicy) DeepCopyInto(out *BucketSoftDeletePolicy) {
	*out = *in
	if in.RetentionDurationSeconds != nil {
		in, out := &in.RetentionDurationSeconds, &out.RetentionDurationSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSoftDeletePolicy.
func (in *BucketSoftDeletePolicy) DeepCopy() *BucketSoftDeletePolicy {
	if in == nil {
		return nil
	}
	out := new(BucketSoftDeletePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSoftDeletePolicyStatus) DeepCopyInto(out *BucketSoftDeletePolicyStatus) {
	*out = *in
	if in.EffectiveTime != nil {
		in, out := &in.EffectiveTime, &out.EffectiveTime
		*out = new(string)
		**out = **in
	}
	if in.RetentionDurationSeconds != nil {
		in, out := &in.RetentionDurationSeconds, &out.RetentionDurationSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSoftDeletePolicyStatus.
func (in *BucketSoftDeletePolicyStatus) DeepCopy() *BucketSoftDeletePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BucketSoftDeletePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketVersioning) DeepCopyInto(out *BucketVersioning) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketVersioning.
func (in *BucketVersioning) DeepCopy() *BucketVersioning {
	if in == nil {
		return nil
	}
	out := new(BucketVersioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketWebsite) DeepCopyInto(out *BucketWebsite) {
	*out = *in
	if in.MainPageSuffix != nil {
		in, out := &in.MainPageSuffix, &out.MainPageSuffix
		*out = new(string)
		**out = **in
	}
	if in.NotFoundPage != nil {
		in, out := &in.NotFoundPage, &out.NotFoundPage
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketWebsite.
func (in *BucketWebsite) DeepCopy() *BucketWebsite {
	if in == nil {
		return nil
	}
	out := new(BucketWebsite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultobjectaccesscontrolProjectTeamStatus) DeepCopyInto(out *DefaultobjectaccesscontrolProjectTeamStatus) {
	*out = *in
	if in.ProjectNumber != nil {
		in, out := &in.ProjectNumber, &out.ProjectNumber
		*out = new(string)
		**out = **in
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultobjectaccesscontrolProjectTeamStatus.
func (in *DefaultobjectaccesscontrolProjectTeamStatus) DeepCopy() *DefaultobjectaccesscontrolProjectTeamStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultobjectaccesscontrolProjectTeamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucket) DeepCopyInto(out *StorageBucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucket.
func (in *StorageBucket) DeepCopy() *StorageBucket {
	if in == nil {
		return nil
	}
	out := new(StorageBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageBucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketAccessControl) DeepCopyInto(out *StorageBucketAccessControl) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketAccessControl.
func (in *StorageBucketAccessControl) DeepCopy() *StorageBucketAccessControl {
	if in == nil {
		return nil
	}
	out := new(StorageBucketAccessControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageBucketAccessControl) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketAccessControlList) DeepCopyInto(out *StorageBucketAccessControlList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageBucketAccessControl, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketAccessControlList.
func (in *StorageBucketAccessControlList) DeepCopy() *StorageBucketAccessControlList {
	if in == nil {
		return nil
	}
	out := new(StorageBucketAccessControlList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageBucketAccessControlList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketAccessControlSpec) DeepCopyInto(out *StorageBucketAccessControlSpec) {
	*out = *in
	out.BucketRef = in.BucketRef
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketAccessControlSpec.
func (in *StorageBucketAccessControlSpec) DeepCopy() *StorageBucketAccessControlSpec {
	if in == nil {
		return nil
	}
	out := new(StorageBucketAccessControlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketAccessControlStatus) DeepCopyInto(out *StorageBucketAccessControlStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketAccessControlStatus.
func (in *StorageBucketAccessControlStatus) DeepCopy() *StorageBucketAccessControlStatus {
	if in == nil {
		return nil
	}
	out := new(StorageBucketAccessControlStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketList) DeepCopyInto(out *StorageBucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageBucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketList.
func (in *StorageBucketList) DeepCopy() *StorageBucketList {
	if in == nil {
		return nil
	}
	out := new(StorageBucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageBucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketSpec) DeepCopyInto(out *StorageBucketSpec) {
	*out = *in
	if in.Autoclass != nil {
		in, out := &in.Autoclass, &out.Autoclass
		*out = new(BucketAutoclass)
		**out = **in
	}
	if in.BucketPolicyOnly != nil {
		in, out := &in.BucketPolicyOnly, &out.BucketPolicyOnly
		*out = new(bool)
		**out = **in
	}
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = make([]BucketCors, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomPlacementConfig != nil {
		in, out := &in.CustomPlacementConfig, &out.CustomPlacementConfig
		*out = new(BucketCustomPlacementConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultEventBasedHold != nil {
		in, out := &in.DefaultEventBasedHold, &out.DefaultEventBasedHold
		*out = new(bool)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(BucketEncryption)
		**out = **in
	}
	if in.LifecycleRule != nil {
		in, out := &in.LifecycleRule, &out.LifecycleRule
		*out = make([]BucketLifecycleRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(BucketLogging)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicAccessPrevention != nil {
		in, out := &in.PublicAccessPrevention, &out.PublicAccessPrevention
		*out = new(string)
		**out = **in
	}
	if in.RequesterPays != nil {
		in, out := &in.RequesterPays, &out.RequesterPays
		*out = new(bool)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = new(BucketRetentionPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.SoftDeletePolicy != nil {
		in, out := &in.SoftDeletePolicy, &out.SoftDeletePolicy
		*out = new(BucketSoftDeletePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.UniformBucketLevelAccess != nil {
		in, out := &in.UniformBucketLevelAccess, &out.UniformBucketLevelAccess
		*out = new(bool)
		**out = **in
	}
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = new(BucketVersioning)
		**out = **in
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = new(BucketWebsite)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketSpec.
func (in *StorageBucketSpec) DeepCopy() *StorageBucketSpec {
	if in == nil {
		return nil
	}
	out := new(StorageBucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBucketStatus) DeepCopyInto(out *StorageBucketStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(BucketObservedStateStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.SelfLink != nil {
		in, out := &in.SelfLink, &out.SelfLink
		*out = new(string)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBucketStatus.
func (in *StorageBucketStatus) DeepCopy() *StorageBucketStatus {
	if in == nil {
		return nil
	}
	out := new(StorageBucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDefaultObjectAccessControl) DeepCopyInto(out *StorageDefaultObjectAccessControl) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDefaultObjectAccessControl.
func (in *StorageDefaultObjectAccessControl) DeepCopy() *StorageDefaultObjectAccessControl {
	if in == nil {
		return nil
	}
	out := new(StorageDefaultObjectAccessControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageDefaultObjectAccessControl) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDefaultObjectAccessControlList) DeepCopyInto(out *StorageDefaultObjectAccessControlList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageDefaultObjectAccessControl, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDefaultObjectAccessControlList.
func (in *StorageDefaultObjectAccessControlList) DeepCopy() *StorageDefaultObjectAccessControlList {
	if in == nil {
		return nil
	}
	out := new(StorageDefaultObjectAccessControlList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageDefaultObjectAccessControlList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDefaultObjectAccessControlSpec) DeepCopyInto(out *StorageDefaultObjectAccessControlSpec) {
	*out = *in
	out.BucketRef = in.BucketRef
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDefaultObjectAccessControlSpec.
func (in *StorageDefaultObjectAccessControlSpec) DeepCopy() *StorageDefaultObjectAccessControlSpec {
	if in == nil {
		return nil
	}
	out := new(StorageDefaultObjectAccessControlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDefaultObjectAccessControlStatus) DeepCopyInto(out *StorageDefaultObjectAccessControlStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.EntityId != nil {
		in, out := &in.EntityId, &out.EntityId
		*out = new(string)
		**out = **in
	}
	if in.Generation != nil {
		in, out := &in.Generation, &out.Generation
		*out = new(int64)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ProjectTeam != nil {
		in, out := &in.ProjectTeam, &out.ProjectTeam
		*out = new(DefaultobjectaccesscontrolProjectTeamStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDefaultObjectAccessControlStatus.
func (in *StorageDefaultObjectAccessControlStatus) DeepCopy() *StorageDefaultObjectAccessControlStatus {
	if in == nil {
		return nil
	}
	out := new(StorageDefaultObjectAccessControlStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageNotification) DeepCopyInto(out *StorageNotification) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageNotification.
func (in *StorageNotification) DeepCopy() *StorageNotification {
	if in == nil {
		return nil
	}
	out := new(StorageNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageNotification) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageNotificationList) DeepCopyInto(out *StorageNotificationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageNotificationList.
func (in *StorageNotificationList) DeepCopy() *StorageNotificationList {
	if in == nil {
		return nil
	}
	out := new(StorageNotificationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageNotificationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageNotificationSpec) DeepCopyInto(out *StorageNotificationSpec) {
	*out = *in
	out.BucketRef = in.BucketRef
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ObjectNamePrefix != nil {
		in, out := &in.ObjectNamePrefix, &out.ObjectNamePrefix
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	out.TopicRef = in.TopicRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageNotificationSpec.
func (in *StorageNotificationSpec) DeepCopy() *StorageNotificationSpec {
	if in == nil {
		return nil
	}
	out := new(StorageNotificationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageNotificationStatus) DeepCopyInto(out *StorageNotificationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.NotificationId != nil {
		in, out := &in.NotificationId, &out.NotificationId
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.SelfLink != nil {
		in, out := &in.SelfLink, &out.SelfLink
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageNotificationStatus.
func (in *StorageNotificationStatus) DeepCopy() *StorageNotificationStatus {
	if in == nil {
		return nil
	}
	out := new(StorageNotificationStatus)
	in.DeepCopyInto(out)
	return out
}
