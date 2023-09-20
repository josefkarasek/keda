//go:build !ignore_autogenerated

/*
Copyright 2023 The KEDA Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/autoscaling/v2"
	"k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedConfig) DeepCopyInto(out *AdvancedConfig) {
	*out = *in
	if in.HorizontalPodAutoscalerConfig != nil {
		in, out := &in.HorizontalPodAutoscalerConfig, &out.HorizontalPodAutoscalerConfig
		*out = new(HorizontalPodAutoscalerConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedConfig.
func (in *AdvancedConfig) DeepCopy() *AdvancedConfig {
	if in == nil {
		return nil
	}
	out := new(AdvancedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthEnvironment) DeepCopyInto(out *AuthEnvironment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthEnvironment.
func (in *AuthEnvironment) DeepCopy() *AuthEnvironment {
	if in == nil {
		return nil
	}
	out := new(AuthEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthPodIdentity) DeepCopyInto(out *AuthPodIdentity) {
	*out = *in
	if in.IdentityID != nil {
		in, out := &in.IdentityID, &out.IdentityID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthPodIdentity.
func (in *AuthPodIdentity) DeepCopy() *AuthPodIdentity {
	if in == nil {
		return nil
	}
	out := new(AuthPodIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthSecretTargetRef) DeepCopyInto(out *AuthSecretTargetRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthSecretTargetRef.
func (in *AuthSecretTargetRef) DeepCopy() *AuthSecretTargetRef {
	if in == nil {
		return nil
	}
	out := new(AuthSecretTargetRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationRef) DeepCopyInto(out *AuthenticationRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationRef.
func (in *AuthenticationRef) DeepCopy() *AuthenticationRef {
	if in == nil {
		return nil
	}
	out := new(AuthenticationRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVault) DeepCopyInto(out *AzureKeyVault) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]AzureKeyVaultSecret, len(*in))
		copy(*out, *in)
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(AzureKeyVaultCredentials)
		(*in).DeepCopyInto(*out)
	}
	if in.PodIdentity != nil {
		in, out := &in.PodIdentity, &out.PodIdentity
		*out = new(AuthPodIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.Cloud != nil {
		in, out := &in.Cloud, &out.Cloud
		*out = new(AzureKeyVaultCloudInfo)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVault.
func (in *AzureKeyVault) DeepCopy() *AzureKeyVault {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultClientSecret) DeepCopyInto(out *AzureKeyVaultClientSecret) {
	*out = *in
	out.ValueFrom = in.ValueFrom
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultClientSecret.
func (in *AzureKeyVaultClientSecret) DeepCopy() *AzureKeyVaultClientSecret {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultClientSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultCloudInfo) DeepCopyInto(out *AzureKeyVaultCloudInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultCloudInfo.
func (in *AzureKeyVaultCloudInfo) DeepCopy() *AzureKeyVaultCloudInfo {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultCloudInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultCredentials) DeepCopyInto(out *AzureKeyVaultCredentials) {
	*out = *in
	if in.ClientSecret != nil {
		in, out := &in.ClientSecret, &out.ClientSecret
		*out = new(AzureKeyVaultClientSecret)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultCredentials.
func (in *AzureKeyVaultCredentials) DeepCopy() *AzureKeyVaultCredentials {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKeyVaultSecret) DeepCopyInto(out *AzureKeyVaultSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKeyVaultSecret.
func (in *AzureKeyVaultSecret) DeepCopy() *AzureKeyVaultSecret {
	if in == nil {
		return nil
	}
	out := new(AzureKeyVaultSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudEvent) DeepCopyInto(out *CloudEvent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudEvent.
func (in *CloudEvent) DeepCopy() *CloudEvent {
	if in == nil {
		return nil
	}
	out := new(CloudEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudEvent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudEventHandler) DeepCopyInto(out *CloudEventHandler) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudEventHandler.
func (in *CloudEventHandler) DeepCopy() *CloudEventHandler {
	if in == nil {
		return nil
	}
	out := new(CloudEventHandler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudEventList) DeepCopyInto(out *CloudEventList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudEvent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudEventList.
func (in *CloudEventList) DeepCopy() *CloudEventList {
	if in == nil {
		return nil
	}
	out := new(CloudEventList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudEventList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudEventSpec) DeepCopyInto(out *CloudEventSpec) {
	*out = *in
	if in.EventHandlers != nil {
		in, out := &in.EventHandlers, &out.EventHandlers
		*out = make([]CloudEventHandler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudEventSpec.
func (in *CloudEventSpec) DeepCopy() *CloudEventSpec {
	if in == nil {
		return nil
	}
	out := new(CloudEventSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTriggerAuthentication) DeepCopyInto(out *ClusterTriggerAuthentication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTriggerAuthentication.
func (in *ClusterTriggerAuthentication) DeepCopy() *ClusterTriggerAuthentication {
	if in == nil {
		return nil
	}
	out := new(ClusterTriggerAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTriggerAuthentication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTriggerAuthenticationList) DeepCopyInto(out *ClusterTriggerAuthenticationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterTriggerAuthentication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTriggerAuthenticationList.
func (in *ClusterTriggerAuthenticationList) DeepCopy() *ClusterTriggerAuthenticationList {
	if in == nil {
		return nil
	}
	out := new(ClusterTriggerAuthenticationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTriggerAuthenticationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credential.
func (in *Credential) DeepCopy() *Credential {
	if in == nil {
		return nil
	}
	out := new(Credential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fallback) DeepCopyInto(out *Fallback) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fallback.
func (in *Fallback) DeepCopy() *Fallback {
	if in == nil {
		return nil
	}
	out := new(Fallback)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersionKindResource) DeepCopyInto(out *GroupVersionKindResource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionKindResource.
func (in *GroupVersionKindResource) DeepCopy() *GroupVersionKindResource {
	if in == nil {
		return nil
	}
	out := new(GroupVersionKindResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HashiCorpVault) DeepCopyInto(out *HashiCorpVault) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]VaultSecret, len(*in))
		copy(*out, *in)
	}
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(Credential)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HashiCorpVault.
func (in *HashiCorpVault) DeepCopy() *HashiCorpVault {
	if in == nil {
		return nil
	}
	out := new(HashiCorpVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthStatus) DeepCopyInto(out *HealthStatus) {
	*out = *in
	if in.NumberOfFailures != nil {
		in, out := &in.NumberOfFailures, &out.NumberOfFailures
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthStatus.
func (in *HealthStatus) DeepCopy() *HealthStatus {
	if in == nil {
		return nil
	}
	out := new(HealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerConfig) DeepCopyInto(out *HorizontalPodAutoscalerConfig) {
	*out = *in
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(v2.HorizontalPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerConfig.
func (in *HorizontalPodAutoscalerConfig) DeepCopy() *HorizontalPodAutoscalerConfig {
	if in == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rollout) DeepCopyInto(out *Rollout) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rollout.
func (in *Rollout) DeepCopy() *Rollout {
	if in == nil {
		return nil
	}
	out := new(Rollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleTarget) DeepCopyInto(out *ScaleTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleTarget.
func (in *ScaleTarget) DeepCopy() *ScaleTarget {
	if in == nil {
		return nil
	}
	out := new(ScaleTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleTriggers) DeepCopyInto(out *ScaleTriggers) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AuthenticationRef != nil {
		in, out := &in.AuthenticationRef, &out.AuthenticationRef
		*out = new(AuthenticationRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleTriggers.
func (in *ScaleTriggers) DeepCopy() *ScaleTriggers {
	if in == nil {
		return nil
	}
	out := new(ScaleTriggers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledJob) DeepCopyInto(out *ScaledJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledJob.
func (in *ScaledJob) DeepCopy() *ScaledJob {
	if in == nil {
		return nil
	}
	out := new(ScaledJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledJobList) DeepCopyInto(out *ScaledJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScaledJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledJobList.
func (in *ScaledJobList) DeepCopy() *ScaledJobList {
	if in == nil {
		return nil
	}
	out := new(ScaledJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledJobSpec) DeepCopyInto(out *ScaledJobSpec) {
	*out = *in
	if in.JobTargetRef != nil {
		in, out := &in.JobTargetRef, &out.JobTargetRef
		*out = new(v1.JobSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedJobsHistoryLimit != nil {
		in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	out.Rollout = in.Rollout
	if in.MinReplicaCount != nil {
		in, out := &in.MinReplicaCount, &out.MinReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	in.ScalingStrategy.DeepCopyInto(&out.ScalingStrategy)
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]ScaleTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledJobSpec.
func (in *ScaledJobSpec) DeepCopy() *ScaledJobSpec {
	if in == nil {
		return nil
	}
	out := new(ScaledJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledJobStatus) DeepCopyInto(out *ScaledJobStatus) {
	*out = *in
	if in.LastActiveTime != nil {
		in, out := &in.LastActiveTime, &out.LastActiveTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledJobStatus.
func (in *ScaledJobStatus) DeepCopy() *ScaledJobStatus {
	if in == nil {
		return nil
	}
	out := new(ScaledJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObject) DeepCopyInto(out *ScaledObject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObject.
func (in *ScaledObject) DeepCopy() *ScaledObject {
	if in == nil {
		return nil
	}
	out := new(ScaledObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledObject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectList) DeepCopyInto(out *ScaledObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScaledObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectList.
func (in *ScaledObjectList) DeepCopy() *ScaledObjectList {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectSpec) DeepCopyInto(out *ScaledObjectSpec) {
	*out = *in
	if in.ScaleTargetRef != nil {
		in, out := &in.ScaleTargetRef, &out.ScaleTargetRef
		*out = new(ScaleTarget)
		**out = **in
	}
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.IdleReplicaCount != nil {
		in, out := &in.IdleReplicaCount, &out.IdleReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicaCount != nil {
		in, out := &in.MinReplicaCount, &out.MinReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.Advanced != nil {
		in, out := &in.Advanced, &out.Advanced
		*out = new(AdvancedConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]ScaleTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Fallback != nil {
		in, out := &in.Fallback, &out.Fallback
		*out = new(Fallback)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectSpec.
func (in *ScaledObjectSpec) DeepCopy() *ScaledObjectSpec {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectStatus) DeepCopyInto(out *ScaledObjectStatus) {
	*out = *in
	if in.ScaleTargetGVKR != nil {
		in, out := &in.ScaleTargetGVKR, &out.ScaleTargetGVKR
		*out = new(GroupVersionKindResource)
		**out = **in
	}
	if in.OriginalReplicaCount != nil {
		in, out := &in.OriginalReplicaCount, &out.OriginalReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.LastActiveTime != nil {
		in, out := &in.LastActiveTime, &out.LastActiveTime
		*out = (*in).DeepCopy()
	}
	if in.ExternalMetricNames != nil {
		in, out := &in.ExternalMetricNames, &out.ExternalMetricNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceMetricNames != nil {
		in, out := &in.ResourceMetricNames, &out.ResourceMetricNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		copy(*out, *in)
	}
	if in.Health != nil {
		in, out := &in.Health, &out.Health
		*out = make(map[string]HealthStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.PausedReplicaCount != nil {
		in, out := &in.PausedReplicaCount, &out.PausedReplicaCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectStatus.
func (in *ScaledObjectStatus) DeepCopy() *ScaledObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingStrategy) DeepCopyInto(out *ScalingStrategy) {
	*out = *in
	if in.CustomScalingQueueLengthDeduction != nil {
		in, out := &in.CustomScalingQueueLengthDeduction, &out.CustomScalingQueueLengthDeduction
		*out = new(int32)
		**out = **in
	}
	if in.PendingPodConditions != nil {
		in, out := &in.PendingPodConditions, &out.PendingPodConditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingStrategy.
func (in *ScalingStrategy) DeepCopy() *ScalingStrategy {
	if in == nil {
		return nil
	}
	out := new(ScalingStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyRef) DeepCopyInto(out *SecretKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyRef.
func (in *SecretKeyRef) DeepCopy() *SecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(SecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerAuthentication) DeepCopyInto(out *TriggerAuthentication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerAuthentication.
func (in *TriggerAuthentication) DeepCopy() *TriggerAuthentication {
	if in == nil {
		return nil
	}
	out := new(TriggerAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerAuthentication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerAuthenticationList) DeepCopyInto(out *TriggerAuthenticationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TriggerAuthentication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerAuthenticationList.
func (in *TriggerAuthenticationList) DeepCopy() *TriggerAuthenticationList {
	if in == nil {
		return nil
	}
	out := new(TriggerAuthenticationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerAuthenticationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerAuthenticationSpec) DeepCopyInto(out *TriggerAuthenticationSpec) {
	*out = *in
	if in.PodIdentity != nil {
		in, out := &in.PodIdentity, &out.PodIdentity
		*out = new(AuthPodIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretTargetRef != nil {
		in, out := &in.SecretTargetRef, &out.SecretTargetRef
		*out = make([]AuthSecretTargetRef, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]AuthEnvironment, len(*in))
		copy(*out, *in)
	}
	if in.HashiCorpVault != nil {
		in, out := &in.HashiCorpVault, &out.HashiCorpVault
		*out = new(HashiCorpVault)
		(*in).DeepCopyInto(*out)
	}
	if in.AzureKeyVault != nil {
		in, out := &in.AzureKeyVault, &out.AzureKeyVault
		*out = new(AzureKeyVault)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerAuthenticationSpec.
func (in *TriggerAuthenticationSpec) DeepCopy() *TriggerAuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerAuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerAuthenticationStatus) DeepCopyInto(out *TriggerAuthenticationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerAuthenticationStatus.
func (in *TriggerAuthenticationStatus) DeepCopy() *TriggerAuthenticationStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerAuthenticationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueFromSecret) DeepCopyInto(out *ValueFromSecret) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueFromSecret.
func (in *ValueFromSecret) DeepCopy() *ValueFromSecret {
	if in == nil {
		return nil
	}
	out := new(ValueFromSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultSecret) DeepCopyInto(out *VaultSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultSecret.
func (in *VaultSecret) DeepCopy() *VaultSecret {
	if in == nil {
		return nil
	}
	out := new(VaultSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WithTriggers) DeepCopyInto(out *WithTriggers) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WithTriggers.
func (in *WithTriggers) DeepCopy() *WithTriggers {
	if in == nil {
		return nil
	}
	out := new(WithTriggers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WithTriggers) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WithTriggersList) DeepCopyInto(out *WithTriggersList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WithTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WithTriggersList.
func (in *WithTriggersList) DeepCopy() *WithTriggersList {
	if in == nil {
		return nil
	}
	out := new(WithTriggersList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WithTriggersList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WithTriggersSpec) DeepCopyInto(out *WithTriggersSpec) {
	*out = *in
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]ScaleTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WithTriggersSpec.
func (in *WithTriggersSpec) DeepCopy() *WithTriggersSpec {
	if in == nil {
		return nil
	}
	out := new(WithTriggersSpec)
	in.DeepCopyInto(out)
	return out
}
