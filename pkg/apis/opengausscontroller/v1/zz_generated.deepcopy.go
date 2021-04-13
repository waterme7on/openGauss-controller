// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenGauss) DeepCopyInto(out *OpenGauss) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Status = in.Status
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenGauss.
func (in *OpenGauss) DeepCopy() *OpenGauss {
	if in == nil {
		return nil
	}
	out := new(OpenGauss)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenGauss) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenGaussClusterConfiguration) DeepCopyInto(out *OpenGaussClusterConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenGaussClusterConfiguration.
func (in *OpenGaussClusterConfiguration) DeepCopy() *OpenGaussClusterConfiguration {
	if in == nil {
		return nil
	}
	out := new(OpenGaussClusterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenGaussList) DeepCopyInto(out *OpenGaussList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenGauss, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenGaussList.
func (in *OpenGaussList) DeepCopy() *OpenGaussList {
	if in == nil {
		return nil
	}
	out := new(OpenGaussList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenGaussList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenGaussSpec) DeepCopyInto(out *OpenGaussSpec) {
	*out = *in
	if in.OpenGauss != nil {
		in, out := &in.OpenGauss, &out.OpenGauss
		*out = new(OpenGaussClusterConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenGaussSpec.
func (in *OpenGaussSpec) DeepCopy() *OpenGaussSpec {
	if in == nil {
		return nil
	}
	out := new(OpenGaussSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenGaussStatus) DeepCopyInto(out *OpenGaussStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenGaussStatus.
func (in *OpenGaussStatus) DeepCopy() *OpenGaussStatus {
	if in == nil {
		return nil
	}
	out := new(OpenGaussStatus)
	in.DeepCopyInto(out)
	return out
}
