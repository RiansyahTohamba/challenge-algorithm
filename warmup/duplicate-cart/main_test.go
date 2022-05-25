package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Find maximum number of Items in Cart", func() {
		When("list is contain only one item, e.g. ID=1, most frequently in cart", func() {
			It("should return most frequently item", func() {
				items := []int32{1, 2, 1, 1, 4}

				mostId := FindItem(items)

				Expect(mostId).To(Equal(1))
			})
		})

		When("list is contain two item,e.g. ID=1 & ID=2, which have same occurence in cart", func() {
			It("should return minimum id", func() {
				items := []int32{1, 2, 2, 1, 4}

				mostId := FindItem(items)

				Expect(mostId).To(Equal(1))
			})
		})
	})
})
