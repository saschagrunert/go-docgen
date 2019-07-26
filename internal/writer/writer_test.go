package writer_test

import (
	"bytes"

	"github.com/saschagrunert/go-docgen/internal/writer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/urfave/cli"
)

// The actual test suite
var _ = t.Describe("Writer", func() {
	t.Describe("New", func() {
		It("should succeed", func() {
			// Given
			app := cli.NewApp()

			// When
			res := writer.New(app)

			// Then
			Expect(res).NotTo(BeNil())
		})
	})

	t.Describe("Write", func() {
		It("should succeed", func() {
			// Given
			w := writer.New(t.Sut)
			var b bytes.Buffer

			// When
			err := w.Write(&b)

			// Then
			Expect(err).To(BeNil())
			Expect(b.String()).To(ContainSubstring("# DESCRIPTION"))
			Expect(b.String()).To(ContainSubstring("# NAME"))
			Expect(b.String()).To(ContainSubstring("# SYNOPSIS"))
		})
	})
})
