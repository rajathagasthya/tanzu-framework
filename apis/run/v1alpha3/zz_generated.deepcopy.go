//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBootstrap) DeepCopyInto(out *ClusterBootstrap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ClusterBootstrapTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBootstrap.
func (in *ClusterBootstrap) DeepCopy() *ClusterBootstrap {
	if in == nil {
		return nil
	}
	out := new(ClusterBootstrap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBootstrap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBootstrapList) DeepCopyInto(out *ClusterBootstrapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBootstrap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBootstrapList.
func (in *ClusterBootstrapList) DeepCopy() *ClusterBootstrapList {
	if in == nil {
		return nil
	}
	out := new(ClusterBootstrapList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBootstrapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBootstrapPackage) DeepCopyInto(out *ClusterBootstrapPackage) {
	*out = *in
	if in.ValuesFrom != nil {
		in, out := &in.ValuesFrom, &out.ValuesFrom
		*out = new(ValuesFrom)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBootstrapPackage.
func (in *ClusterBootstrapPackage) DeepCopy() *ClusterBootstrapPackage {
	if in == nil {
		return nil
	}
	out := new(ClusterBootstrapPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBootstrapStatus) DeepCopyInto(out *ClusterBootstrapStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBootstrapStatus.
func (in *ClusterBootstrapStatus) DeepCopy() *ClusterBootstrapStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterBootstrapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBootstrapTemplate) DeepCopyInto(out *ClusterBootstrapTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ClusterBootstrapTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBootstrapTemplate.
func (in *ClusterBootstrapTemplate) DeepCopy() *ClusterBootstrapTemplate {
	if in == nil {
		return nil
	}
	out := new(ClusterBootstrapTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBootstrapTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBootstrapTemplateList) DeepCopyInto(out *ClusterBootstrapTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBootstrapTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBootstrapTemplateList.
func (in *ClusterBootstrapTemplateList) DeepCopy() *ClusterBootstrapTemplateList {
	if in == nil {
		return nil
	}
	out := new(ClusterBootstrapTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBootstrapTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBootstrapTemplateSpec) DeepCopyInto(out *ClusterBootstrapTemplateSpec) {
	*out = *in
	if in.CNI != nil {
		in, out := &in.CNI, &out.CNI
		*out = new(ClusterBootstrapPackage)
		(*in).DeepCopyInto(*out)
	}
	if in.CSI != nil {
		in, out := &in.CSI, &out.CSI
		*out = new(ClusterBootstrapPackage)
		(*in).DeepCopyInto(*out)
	}
	if in.CPI != nil {
		in, out := &in.CPI, &out.CPI
		*out = new(ClusterBootstrapPackage)
		(*in).DeepCopyInto(*out)
	}
	if in.Kapp != nil {
		in, out := &in.Kapp, &out.Kapp
		*out = new(ClusterBootstrapPackage)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalPackages != nil {
		in, out := &in.AdditionalPackages, &out.AdditionalPackages
		*out = make([]*ClusterBootstrapPackage, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ClusterBootstrapPackage)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBootstrapTemplateSpec.
func (in *ClusterBootstrapTemplateSpec) DeepCopy() *ClusterBootstrapTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterBootstrapTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerImageInfo) DeepCopyInto(out *ContainerImageInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerImageInfo.
func (in *ContainerImageInfo) DeepCopy() *ContainerImageInfo {
	if in == nil {
		return nil
	}
	out := new(ContainerImageInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KappConfig) DeepCopyInto(out *KappConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KappConfig.
func (in *KappConfig) DeepCopy() *KappConfig {
	if in == nil {
		return nil
	}
	out := new(KappConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KappController) DeepCopyInto(out *KappController) {
	*out = *in
	in.Deployment.DeepCopyInto(&out.Deployment)
	out.Config = in.Config
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KappController.
func (in *KappController) DeepCopy() *KappController {
	if in == nil {
		return nil
	}
	out := new(KappController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KappControllerConfig) DeepCopyInto(out *KappControllerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KappControllerConfig.
func (in *KappControllerConfig) DeepCopy() *KappControllerConfig {
	if in == nil {
		return nil
	}
	out := new(KappControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KappControllerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KappControllerConfigList) DeepCopyInto(out *KappControllerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KappControllerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KappControllerConfigList.
func (in *KappControllerConfigList) DeepCopy() *KappControllerConfigList {
	if in == nil {
		return nil
	}
	out := new(KappControllerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KappControllerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KappControllerConfigSpec) DeepCopyInto(out *KappControllerConfigSpec) {
	*out = *in
	in.KappController.DeepCopyInto(&out.KappController)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KappControllerConfigSpec.
func (in *KappControllerConfigSpec) DeepCopy() *KappControllerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KappControllerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KappControllerConfigStatus) DeepCopyInto(out *KappControllerConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KappControllerConfigStatus.
func (in *KappControllerConfigStatus) DeepCopy() *KappControllerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KappControllerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KappDeployment) DeepCopyInto(out *KappDeployment) {
	*out = *in
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KappDeployment.
func (in *KappDeployment) DeepCopy() *KappDeployment {
	if in == nil {
		return nil
	}
	out := new(KappDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesSpec) DeepCopyInto(out *KubernetesSpec) {
	*out = *in
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = new(ContainerImageInfo)
		**out = **in
	}
	if in.Pause != nil {
		in, out := &in.Pause, &out.Pause
		*out = new(ContainerImageInfo)
		**out = **in
	}
	if in.CoreDNS != nil {
		in, out := &in.CoreDNS, &out.CoreDNS
		*out = new(ContainerImageInfo)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesSpec.
func (in *KubernetesSpec) DeepCopy() *KubernetesSpec {
	if in == nil {
		return nil
	}
	out := new(KubernetesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSImage) DeepCopyInto(out *OSImage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSImage.
func (in *OSImage) DeepCopy() *OSImage {
	if in == nil {
		return nil
	}
	out := new(OSImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OSImage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSImageList) DeepCopyInto(out *OSImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OSImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSImageList.
func (in *OSImageList) DeepCopy() *OSImageList {
	if in == nil {
		return nil
	}
	out := new(OSImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OSImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSImageSpec) DeepCopyInto(out *OSImageSpec) {
	*out = *in
	out.OS = in.OS
	in.Image.DeepCopyInto(&out.Image)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSImageSpec.
func (in *OSImageSpec) DeepCopy() *OSImageSpec {
	if in == nil {
		return nil
	}
	out := new(OSImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSImageStatus) DeepCopyInto(out *OSImageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1beta1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSImageStatus.
func (in *OSImageStatus) DeepCopy() *OSImageStatus {
	if in == nil {
		return nil
	}
	out := new(OSImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSInfo) DeepCopyInto(out *OSInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSInfo.
func (in *OSInfo) DeepCopy() *OSInfo {
	if in == nil {
		return nil
	}
	out := new(OSInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TanzuKubernetesRelease) DeepCopyInto(out *TanzuKubernetesRelease) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TanzuKubernetesRelease.
func (in *TanzuKubernetesRelease) DeepCopy() *TanzuKubernetesRelease {
	if in == nil {
		return nil
	}
	out := new(TanzuKubernetesRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TanzuKubernetesRelease) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TanzuKubernetesReleaseList) DeepCopyInto(out *TanzuKubernetesReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TanzuKubernetesRelease, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TanzuKubernetesReleaseList.
func (in *TanzuKubernetesReleaseList) DeepCopy() *TanzuKubernetesReleaseList {
	if in == nil {
		return nil
	}
	out := new(TanzuKubernetesReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TanzuKubernetesReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TanzuKubernetesReleaseSpec) DeepCopyInto(out *TanzuKubernetesReleaseSpec) {
	*out = *in
	in.Kubernetes.DeepCopyInto(&out.Kubernetes)
	if in.OSImages != nil {
		in, out := &in.OSImages, &out.OSImages
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.BootstrapPackages != nil {
		in, out := &in.BootstrapPackages, &out.BootstrapPackages
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TanzuKubernetesReleaseSpec.
func (in *TanzuKubernetesReleaseSpec) DeepCopy() *TanzuKubernetesReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(TanzuKubernetesReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TanzuKubernetesReleaseStatus) DeepCopyInto(out *TanzuKubernetesReleaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1beta1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TanzuKubernetesReleaseStatus.
func (in *TanzuKubernetesReleaseStatus) DeepCopy() *TanzuKubernetesReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(TanzuKubernetesReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValuesFrom) DeepCopyInto(out *ValuesFrom) {
	*out = *in
	if in.ProviderRef != nil {
		in, out := &in.ProviderRef, &out.ProviderRef
		*out = new(v1.TypedLocalObjectReference)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValuesFrom.
func (in *ValuesFrom) DeepCopy() *ValuesFrom {
	if in == nil {
		return nil
	}
	out := new(ValuesFrom)
	in.DeepCopyInto(out)
	return out
}
