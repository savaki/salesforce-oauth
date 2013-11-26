package oauth

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOauth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Salesforce-Oauth Suite")
}
