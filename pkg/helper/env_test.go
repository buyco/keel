package helper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/buyco/keel/pkg/helper"
)

var _ = Describe("Env", func() {

	var err error

	Context("With allowed environment", func() {

		It("should load env file", func() {
			stdout := CaptureStdout(func() { err = LoadEnvFile("development") })
			Expect(stdout).To(BeEmpty())
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
