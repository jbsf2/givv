package givv

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGivv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Givv Suite")
}
