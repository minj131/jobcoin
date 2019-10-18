package jobcoin_test

import (
	"errors"

	"github.com/minj131/jobcoin/config"
	"github.com/minj131/jobcoin/pkg/api"
	"github.com/minj131/jobcoin/pkg/jobcoin"
	"github.com/minj131/jobcoin/pkg/jobcoin/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jobcoin", func() {
	var (
		err     error
		cfg     *config.Config
		j       *jobcoin.Jobcoin
		mockAPI *mocks.API
	)

	BeforeEach(func() {
		Expect(err).NotTo(HaveOccurred())

		f := float64(0.01)

		cfg = &config.Config{
			API:           &config.API{},
			DestAddresses: []string{"test1", "test2"},
			MixerAddress:  "fake",
			Fee:           &f,
			AdminAddress:  "fake",
		}

		j = jobcoin.New(cfg)
		mockAPI = &mocks.API{}

		j.API = mockAPI

		mockAPI.GetAddressInfoReturns(&api.AddressInfo{Balance: "1"}, nil)
		mockAPI.GetTransactionsReturns([]api.Transactions{}, nil)
		mockAPI.PostTransactionReturns(nil)
	})

	Context("Run", func() {
		It("should return an error if polling for change fails", func() {
			mockAPI.GetAddressInfoReturns(nil, errors.New("fake"))
			err := j.Run()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to poll for change in mixer address"))
		})

		It("should return an error if transfer coins fails", func() {
			mockAPI.PostTransactionReturns(errors.New("fake"))
			err := j.Run()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to transfer coins to admin account"))
		})

		It("should return an error if sending coins fails", func() {
			mockAPI.PostTransactionReturnsOnCall(1, errors.New("fake"))
			err := j.Run()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to send jobcoins to"))
		})

		It("should not return an error", func() {
			err := j.Run()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("PollAddressForChange", func() {
		It("should return non empty string when balance is not zero", func() {
			res, err := j.PollAddressForChange()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
		})
	})

	Context("TransferJobcoin", func() {
		It("should not return an error when transfering coins", func() {
			err := j.TransferJobcoin("", "", "")
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
