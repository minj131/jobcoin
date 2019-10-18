package util_test

import (
	"github.com/minj131/jobcoin/pkg/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {
	Context("GenerateRandomHash", func() {
		It("should return all unique random values", func() {
			counter := make(map[string]struct{})
			for i := 0; i < 1000; i++ {
				rng := util.GenerateRandomHash(6)
				if _, ok := counter[rng]; !ok { // only append if unique
					counter[rng] = struct{}{}
				}
			}
			Expect(len(counter)).To(Equal(1000))
		})
	})

	Context("DivideEvenly", func() {
		It("should divide a number evenly", func() {
			check := util.DivideEvenly(12.0, 3)
			Expect(check).To(Equal(4.0))
		})
	})

	Context("FeeCalculator", func() {
		It("should correctly deterine the fee and return the fee and new balance", func() {
			fee := 0.5
			balance := 10.0
			newBalance, _ := util.FeeCalculator(balance, fee)
			Expect(newBalance).To(Equal(5.0))
		})
	})
})
