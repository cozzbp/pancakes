package pancakes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPancakes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pancakes Suite")
}
