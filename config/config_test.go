package config_test

import (
	"path/filepath"

	"github.com/minj131/jobcoin/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		cfg *config.Config
	)

	BeforeEach(func() {
		f := float64(0.01)

		cfg = &config.Config{
			API:           &config.API{},
			DestAddresses: []string{"test1", "test2"},
			MixerAddress:  "mixer",
			Fee:           &f,
			AdminAddress:  "admin",
		}
	})

	Context("Init", func() {
		It("should return an error if api config is not set", func() {
			cfg.API = nil
			err := cfg.Init("")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("API config is not passed"))
		})

		It("should return an error if dest address is not set", func() {
			cfg.DestAddresses = nil
			err := cfg.Init("")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Destination addresses must be set and not empty"))
		})

		It("should return an error if dest address is empty", func() {
			cfg.DestAddresses = []string{}
			err := cfg.Init("")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Destination addresses must be set and not empty"))
		})

		It("should set fee to default if not set", func() {
			cfg.Fee = nil
			err := cfg.Init("")
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error if admin address is not set", func() {
			cfg.AdminAddress = ""
			err := cfg.Init("")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Admin address is not passed"))
		})

		It("should not return any errors when all fields are set correctly", func() {
			err := cfg.Init("")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("ReadConfigFile", func() {
		var path string

		BeforeEach(func() {
			path = "badpath.yaml"
		})

		It("should return an error if a bad path to config is passed", func() {
			err := cfg.ReadConfigFile(path)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to get api configurations"))
		})

		It("should not return an error if a good path to config is passed", func() {
			path = filepath.Join("../", "config.yaml")
			err := cfg.ReadConfigFile(path)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
