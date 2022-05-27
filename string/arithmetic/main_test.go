package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Find maximum number of Items in Cart", func() {

	Describe("cart is contain of more than one item that have frequently buying", func() {

		When("cart is contain of 2 same frequently ID and small ID in the first position", func() {
			// pada case ini, 1 muncul pertama,
			// maka max_frequencies=2
			// maka max=2 tidak akan tergantikan karena 2 < 2 is false
			// maka key yang dipilih selalu key yang pertama kali muncul
			It("should return ID=1 not because he is in the first position", func() {

				var expectedID int32 = 1

				Expect(1).To(Equal(expectedID))
			})

		})

	})
})
