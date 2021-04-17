package imdb_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestImdb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Imdb Suite")
}
