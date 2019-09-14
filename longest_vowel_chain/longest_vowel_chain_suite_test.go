package longest_vowel_chain_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLongestVowelChain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LongestVowelChain Suite")
}
