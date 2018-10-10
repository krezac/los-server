package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/krezac/los-server/cmd/los-server"
)

var _ = Describe("LosServer", func() {

	Describe("Trying to test function", func() {
		Context("Which should return true", func() {
			It("should return true", func() {
				Expect(main.FunctionToTest()).To(Equal(true))
			})
		})
	})

})
