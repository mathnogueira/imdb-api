package movie_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMovie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movie Suite")
}
