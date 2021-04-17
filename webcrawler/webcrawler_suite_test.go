package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWebcrawler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Webcrawler Suite")
}
