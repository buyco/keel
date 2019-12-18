package helper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/url"

	. "github.com/buyco/keel/pkg/helper"
)

var _ = Describe("Request", func() {

	Context("With dynamic route", func() {

		It("should return a map", func() {
			request, err := http.NewRequest("test", "/foo", nil)
			vars := GetRouteVars(request)
			Expect(vars).To(BeAssignableToTypeOf(make(map[string]string)))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("With URL parameters", func() {

		It("should return a Values struct", func() {
			request, err := http.NewRequest("test", "/foo", nil)
			vars := GetQueryVars(request)
			Expect(vars).To(BeAssignableToTypeOf(url.Values{}))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
