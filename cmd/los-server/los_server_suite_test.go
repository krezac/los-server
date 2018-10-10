package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLosServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LosServer Suite")
}
