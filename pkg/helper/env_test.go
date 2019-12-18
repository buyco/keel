package helper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/buyco/keel/pkg/helper"
)

var _ = Describe("Env", func() {

	var err error

	Context("With allowed environment", func() {

		It("should fail to load non-existent file", func() {
			stdout := CaptureStdout(func() { err = LoadEnvFile("foo", "development") })
			Expect(stdout).To(BeEmpty())
			Expect(err).To(HaveOccurred())
		})

		It("should load env file", func() {
			stdout := CaptureStdout(func() { err = LoadEnvFile("../../internal/tests/.env-test", "development") })
			Expect(stdout).To(BeEmpty())
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("With not allowed environment", func() {

		It("should not load non-existent file", func() {
			err = LoadEnvFile("foo", "bar")
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
