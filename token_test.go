package oauth

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Token", func() {
	Context("With json marshall/unmarshall", func() {
		It("should do something", func() {
			token := Token{
				Id: "1234",
			}
			bytes, err := json.Marshal(token)
			Expect(err).To(BeNil())
			Expect(string(bytes)).To(Equal("{\"id\":\"1234\"}"))
		})
	})
})
