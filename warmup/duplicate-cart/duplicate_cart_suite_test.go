package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDuplicateCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DuplicateCart Suite")
}
