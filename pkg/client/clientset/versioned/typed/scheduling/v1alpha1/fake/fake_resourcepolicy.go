/*
Copyright 2025 The Kubernetes Authors.

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

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/KunWuLuan/resourcepolicyapi/pkg/apis/scheduling/v1alpha1"
	schedulingv1alpha1 "github.com/KunWuLuan/resourcepolicyapi/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeResourcePolicies implements ResourcePolicyInterface
type fakeResourcePolicies struct {
	*gentype.FakeClientWithList[*v1alpha1.ResourcePolicy, *v1alpha1.ResourcePolicyList]
	Fake *FakeSchedulingV1alpha1
}

func newFakeResourcePolicies(fake *FakeSchedulingV1alpha1, namespace string) schedulingv1alpha1.ResourcePolicyInterface {
	return &fakeResourcePolicies{
		gentype.NewFakeClientWithList[*v1alpha1.ResourcePolicy, *v1alpha1.ResourcePolicyList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("resourcepolicies"),
			v1alpha1.SchemeGroupVersion.WithKind("ResourcePolicy"),
			func() *v1alpha1.ResourcePolicy { return &v1alpha1.ResourcePolicy{} },
			func() *v1alpha1.ResourcePolicyList { return &v1alpha1.ResourcePolicyList{} },
			func(dst, src *v1alpha1.ResourcePolicyList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.ResourcePolicyList) []*v1alpha1.ResourcePolicy {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.ResourcePolicyList, items []*v1alpha1.ResourcePolicy) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
