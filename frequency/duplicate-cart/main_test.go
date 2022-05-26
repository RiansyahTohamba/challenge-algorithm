package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Find maximum number of Items in Cart", func() {
	Describe("list is contain only one item most frequently in cart", func() {
		When("only one ID=1, most frequently in cart", func() {
			It("should return 1", func() {
				items := []int32{2, 2, 1, 1, 1, 4}

				mostId := FindItem(items)
				var expectedID int32 = 1

				Expect(mostId).To(Equal(expectedID))
			})
		})
	})

	Describe("cart is contain of more than one item that have frequently buying", func() {

		When("cart is contain of 2 same frequently ID and small ID in the first position", func() {
			// pada case ini, 1 muncul pertama,
			// maka max_frequencies=2
			// maka max=2 tidak akan tergantikan karena 2 < 2 is false
			// maka key yang dipilih selalu key yang pertama kali muncul
			It("should return ID=1 not because he is in the first position", func() {
				items := []int32{1, 2, 2, 1, 4, 4}

				mostId := FindItem(items)
				var expectedID int32 = 1

				Expect(mostId).To(Equal(expectedID))
			})

		})
		When("cart is contain of 2 same frequently ID and big ID in the first position", func() {
			It("should return ID=1 even it is in the last position", func() {
				items := []int32{2, 2, 4, 4, 1, 1}

				mostId := FindItem(items)
				var expectedID int32 = 1
				Expect(mostId).To(Equal(expectedID))
			})
		})

	})
})
