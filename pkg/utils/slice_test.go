package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/buyco/keel/pkg/utils"
)

var _ = Describe("Slice", func() {

	var (
		testSlice = []string{"foo", "bar"}
	)

	It("should find value in slice", func() {
		exists, index := InArray("foo", testSlice)
		Expect(exists).To(BeTrue())
		Expect(index).To(BeZero())
	})

	It("should not find value in slice", func() {
		exists, index := InArray("woo", testSlice)
		Expect(exists).To(BeFalse())
		Expect(index).To(Equal(-1))
	})

	It("should find string in slice", func() {
		exists := StringInSlice("foo", testSlice)
		Expect(exists).To(BeTrue())
	})

	It("should not find string in slice", func() {
		exists := StringInSlice("woo", testSlice)
		Expect(exists).To(BeFalse())
	})
})
