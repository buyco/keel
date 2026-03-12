package helper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/buyco/keel/pkg/helper"
)

var _ = Describe("Map", func() {

	var (
		testMap = map[string]any{"foo": 0}
	)

	It("should copy given map", func() {
		copyMap := CopyMap(testMap)
		Expect(copyMap["foo"]).To(BeZero())
	})
})
