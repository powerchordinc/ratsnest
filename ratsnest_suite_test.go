package ratsnest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRatsnest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ratsnest Suite")
}
