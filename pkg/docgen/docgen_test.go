package docgen_test

import (
	"github.com/saschagrunert/go-docgen/pkg/docgen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// The actual test suite
var _ = t.Describe("Docgen", func() {
	t.Describe("CliToMarkdown", func() {
		It("should succeed", func() {
			// Given
			// When
			res, err := docgen.CliToMarkdown(t.Sut)

			// Then
			Expect(err).To(BeNil())
			Expect(res).NotTo(BeEmpty())
			Expect(res).To(ContainSubstring("# COMMANDS"))
			Expect(res).To(ContainSubstring("# DESCRIPTION"))
			Expect(res).To(ContainSubstring("# GLOBAL OPTIONS"))
			Expect(res).To(ContainSubstring("## info, i, in"))
			Expect(res).To(ContainSubstring("### sub-config"))
			Expect(res).To(ContainSubstring("Name(8) Description"))
		})
	})

	t.Describe("CliToMan", func() {
		It("should succeed", func() {
			// Given
			// When
			res, err := docgen.CliToMan(t.Sut)

			// Then
			Expect(err).To(BeNil())
			Expect(res).NotTo(BeEmpty())
		})
	})
})
