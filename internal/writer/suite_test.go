package writer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/saschagrunert/go-docgen/test/framework"
)

// TestStorage runs the created specs
func TestStorage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunFrameworkSpecs(t, "writer")
}

// nolint: gochecknoglobals
var t *TestFramework

var _ = BeforeSuite(func() {
	t = NewTestFramework(NilFunc, NilFunc)
	t.Setup()
})

var _ = AfterSuite(func() {
	t.Teardown()
})
