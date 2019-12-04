package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"

	. "github.com/buyco/keel/pkg/utils"
)

var _ = Describe("Test", func() {

	It("should catch stdout", func() {
		stdout := CaptureStdout(func() { log.Printf("foo:bar") })
		Expect(stdout).To(ContainSubstring("foo:bar"))
	})

	It("should not find stdout", func() {
		stdout := CaptureStdout(func() { log.Printf("foo:bar") })
		Expect(stdout).ToNot(ContainSubstring("woo:zoo"))
	})
})
