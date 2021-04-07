/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha3

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	apiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"

	v1alpha4 "sigs.k8s.io/cluster-api-provider-gcp/api/v1alpha4"
)

// ConvertTo converts this GCPCluster to the Hub version (v1alpha4).
func (src *GCPCluster) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1alpha4.GCPCluster)

	if err := Convert_v1alpha3_GCPCluster_To_v1alpha4_GCPCluster(src, dst, nil); err != nil {
		return err
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha4) to this version.
func (dst *GCPCluster) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1alpha4.GCPCluster)

	if err := Convert_v1alpha4_GCPCluster_To_v1alpha3_GCPCluster(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	return nil
}

// ConvertTo converts this GCPClusterList to the Hub version (v1alpha4).
func (src *GCPClusterList) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1alpha4.GCPClusterList)
	return Convert_v1alpha3_GCPClusterList_To_v1alpha4_GCPClusterList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1alpha4) to this version.
func (dst *GCPClusterList) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1alpha4.GCPClusterList)
	return Convert_v1alpha4_GCPClusterList_To_v1alpha3_GCPClusterList(src, dst, nil)
}

// Convert_v1alpha3_GCPClusterStatus_To_v1alpha4_GCPClusterStatus converts GCPCluster.Status from v1alpha3 to v1alpha4.
func Convert_v1alpha3_GCPClusterStatus_To_v1alpha4_GCPClusterStatus(in *GCPClusterStatus, out *v1alpha4.GCPClusterStatus, s apiconversion.Scope) error { // nolint
	if err := autoConvert_v1alpha3_GCPClusterStatus_To_v1alpha4_GCPClusterStatus(in, out, s); err != nil {
		return err
	}

	return nil
}

// Convert_v1alpha3_GCPClusterSpec_To_v1alpha4_GCPClusterSpec.
func Convert_v1alpha3_GCPClusterSpec_To_v1alpha4_GCPClusterSpec(in *GCPClusterSpec, out *v1alpha4.GCPClusterSpec, s apiconversion.Scope) error { //nolint
	if err := autoConvert_v1alpha3_GCPClusterSpec_To_v1alpha4_GCPClusterSpec(in, out, s); err != nil {
		return err
	}

	return nil
}

// Convert_v1alpha4_GCPClusterSpec_To_v1alpha3_GCPClusterSpec converts from the Hub version (v1alpha4) of the GCPClusterSpec to this version.
func Convert_v1alpha4_GCPClusterSpec_To_v1alpha3_GCPClusterSpec(in *v1alpha4.GCPClusterSpec, out *GCPClusterSpec, s apiconversion.Scope) error { // nolint
	if err := autoConvert_v1alpha4_GCPClusterSpec_To_v1alpha3_GCPClusterSpec(in, out, s); err != nil {
		return err
	}

	return nil
}

// Convert_v1alpha4_GCPClusterStatus_To_v1alpha3_GCPClusterStatus.
func Convert_v1alpha4_GCPClusterStatus_To_v1alpha3_GCPClusterStatus(in *v1alpha4.GCPClusterStatus, out *GCPClusterStatus, s apiconversion.Scope) error { //nolint
	if err := autoConvert_v1alpha4_GCPClusterStatus_To_v1alpha3_GCPClusterStatus(in, out, s); err != nil {
		return err
	}

	return nil
}

// Convert_v1alpha3_NetworkSpec_To_v1alpha4_NetworkSpec.
func Convert_v1alpha3_NetworkSpec_To_v1alpha4_NetworkSpec(in *NetworkSpec, out *v1alpha4.NetworkSpec, s apiconversion.Scope) error { //nolint
	out.Subnets = make(v1alpha4.Subnets, len(in.Subnets))
	for i := range in.Subnets {
		out.Subnets[i] = &v1alpha4.SubnetSpec{}
		if err := Convert_v1alpha3_SubnetSpec_To_v1alpha4_SubnetSpec(in.Subnets[i], out.Subnets[i], s); err != nil {
			return err
		}
	}

	return nil
}

// Convert_v1alpha4_NetworkSpec_To_v1alpha3_NetworkSpec.
func Convert_v1alpha4_NetworkSpec_To_v1alpha3_NetworkSpec(in *v1alpha4.NetworkSpec, out *NetworkSpec, s apiconversion.Scope) error { //nolint
	out.Subnets = make(Subnets, len(in.Subnets))
	for i := range in.Subnets {
		out.Subnets[i] = &SubnetSpec{}
		if err := Convert_v1alpha4_SubnetSpec_To_v1alpha3_SubnetSpec(in.Subnets[i], out.Subnets[i], s); err != nil {
			return err
		}
	}

	return nil
}

// Convert_v1alpha3_APIEndpoint_To_v1alpha4_APIEndpoint is an autogenerated conversion function.
func Convert_v1alpha3_APIEndpoint_To_v1alpha4_APIEndpoint(in *apiv1alpha3.APIEndpoint, out *apiv1alpha4.APIEndpoint, s apiconversion.Scope) error {
	return apiv1alpha3.Convert_v1alpha3_APIEndpoint_To_v1alpha4_APIEndpoint(in, out, s)
}

// Convert_v1alpha4_APIEndpoint_To_v1alpha3_APIEndpoint is an autogenerated conversion function.
func Convert_v1alpha4_APIEndpoint_To_v1alpha3_APIEndpoint(in *apiv1alpha4.APIEndpoint, out *apiv1alpha3.APIEndpoint, s apiconversion.Scope) error {
	return apiv1alpha3.Convert_v1alpha4_APIEndpoint_To_v1alpha3_APIEndpoint(in, out, s)
}
