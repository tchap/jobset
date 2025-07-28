package util

import (
	"context"
	"fmt"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/printers"
	"sigs.k8s.io/controller-runtime/pkg/client"
	jobset "sigs.k8s.io/jobset/api/jobset/v1alpha2"
)

type TestRunOption func(*TestRun)

func DumpNamespace(enabled bool) TestRunOption {
	return func(testRun *TestRun) {
		testRun.dumpNamespace = enabled
	}
}

type TestRun struct {
	ctx                   context.Context
	k8sClient             client.Client
	namespaceGenerateName string

	dumpNamespace bool

	Namespace *corev1.Namespace
}

func NewTestRun(ctx context.Context, k8sClient client.Client, namespaceGenerateName string, opts ...TestRunOption) *TestRun {
	testRun := &TestRun{
		ctx:                   ctx,
		k8sClient:             k8sClient,
		namespaceGenerateName: namespaceGenerateName,
	}
	for _, opt := range opts {
		opt(testRun)
	}
	testRun.beforeEach()
	return testRun
}

func (testRun *TestRun) beforeEach() {
	// Create test namespace before each test.
	testRun.Namespace = &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: testRun.namespaceGenerateName,
		},
	}
	gomega.Expect(testRun.k8sClient.Create(testRun.ctx, testRun.Namespace)).To(gomega.Succeed())

	// Wait for namespace to exist before proceeding with test.
	gomega.Eventually(func() error {
		return testRun.k8sClient.Get(testRun.ctx, types.NamespacedName{Namespace: testRun.Namespace.Namespace, Name: testRun.Namespace.Name}, testRun.Namespace)
	}, timeout, interval).Should(gomega.Succeed())
}

func (testRun *TestRun) AfterEach() {
	// Dump the namespace content, optionally.
	// This is only visible on failure since ginkgo.GinkgoWriter is being used.
	if testRun.dumpNamespace {
		var printer printers.YAMLPrinter

		fmt.Fprintf(ginkgo.GinkgoWriter, "\nDumping relevant resources in namespace %s:\n\n", testRun.Namespace.Name)

		// JobSets
		var jobsets jobset.JobSetList
		gomega.Expect(testRun.k8sClient.List(testRun.ctx, &jobsets)).To(gomega.Succeed())
		for _, js := range jobsets.Items {
			gomega.Expect(printer.PrintObj(&js, ginkgo.GinkgoWriter)).To(gomega.Succeed())
		}

		// Jobs
		var jobs batchv1.JobList
		gomega.Expect(testRun.k8sClient.List(testRun.ctx, &jobs)).To(gomega.Succeed())
		for _, job := range jobs.Items {
			// GVK is not set properly for the list items.
			job.SetGroupVersionKind(schema.GroupVersionKind{
				Group:   batchv1.SchemeGroupVersion.Group,
				Version: batchv1.SchemeGroupVersion.Version,
				Kind:    "Job",
			})
			gomega.Expect(printer.PrintObj(&job, ginkgo.GinkgoWriter)).To(gomega.Succeed())
		}
	}

	// Delete test namespace after each test.
	gomega.Expect(testRun.k8sClient.Delete(testRun.ctx, testRun.Namespace)).To(gomega.Succeed())
}
