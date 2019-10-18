package jobcoin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJobcoin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jobcoin Suite")
}
