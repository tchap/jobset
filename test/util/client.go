package util

import (
	"context"

	"github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"
	jobset "sigs.k8s.io/jobset/api/jobset/v1alpha2"
)

func CreateJobSet(ctx context.Context, k8sClient client.Client, js *jobset.JobSet) {
	gomega.Eventually(func() error {
		return k8sClient.Create(ctx, js)
	}, timeout, interval).Should(gomega.Succeed())
}

func UpdateJobSet(ctx context.Context, k8sClient client.Client, js *jobset.JobSet) {
	gomega.Eventually(func() error {
		return k8sClient.Update(ctx, js)
	}, timeout, interval).Should(gomega.Succeed())
}
