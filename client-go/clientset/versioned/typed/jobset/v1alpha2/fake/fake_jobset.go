/*
Copyright 2023 The Kubernetes Authors.
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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	gentype "k8s.io/client-go/gentype"
	v1alpha2 "sigs.k8s.io/jobset/api/jobset/v1alpha2"
	jobsetv1alpha2 "sigs.k8s.io/jobset/client-go/applyconfiguration/jobset/v1alpha2"
	typedjobsetv1alpha2 "sigs.k8s.io/jobset/client-go/clientset/versioned/typed/jobset/v1alpha2"
)

// fakeJobSets implements JobSetInterface
type fakeJobSets struct {
	*gentype.FakeClientWithListAndApply[*v1alpha2.JobSet, *v1alpha2.JobSetList, *jobsetv1alpha2.JobSetApplyConfiguration]
	Fake *FakeJobsetV1alpha2
}

func newFakeJobSets(fake *FakeJobsetV1alpha2, namespace string) typedjobsetv1alpha2.JobSetInterface {
	return &fakeJobSets{
		gentype.NewFakeClientWithListAndApply[*v1alpha2.JobSet, *v1alpha2.JobSetList, *jobsetv1alpha2.JobSetApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha2.SchemeGroupVersion.WithResource("jobsets"),
			v1alpha2.SchemeGroupVersion.WithKind("JobSet"),
			func() *v1alpha2.JobSet { return &v1alpha2.JobSet{} },
			func() *v1alpha2.JobSetList { return &v1alpha2.JobSetList{} },
			func(dst, src *v1alpha2.JobSetList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha2.JobSetList) []*v1alpha2.JobSet { return gentype.ToPointerSlice(list.Items) },
			func(list *v1alpha2.JobSetList, items []*v1alpha2.JobSet) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
