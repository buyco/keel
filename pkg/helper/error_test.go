package helper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/buyco/keel/pkg/helper"
)

var _ = Describe("Error", func() {

	It("should create error with variables args", func() {
		err := ErrorPrintf("foo %s", "bar")
		Expect(err.Error()).To(MatchRegexp("foo bar"))
	})

	It("should create error with string", func() {
		err := ErrorPrint("foo")
		Expect(err.Error()).To(MatchRegexp("foo"))
	})
})
