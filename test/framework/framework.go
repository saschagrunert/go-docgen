package framework

import (
	"fmt"
	"strings"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
	"github.com/urfave/cli"
)

const (
	testAuthor      = "Author"
	testName        = "Name"
	testDescription = "Description"
	testUsageText   = "This is a usage text."
)

// TestFramework is used to support commonnly used test features
type TestFramework struct {
	setup    func(*TestFramework) error
	teardown func(*TestFramework) error
	Sut      *cli.App
}

// NewTestFramework creates a new test framework instance for a given `setup`
// and `teardown` function
func NewTestFramework(setup, teardown func(*TestFramework) error) *TestFramework {
	sut := cli.NewApp()
	sut.Author = testAuthor
	sut.Name = testName
	sut.Description = testDescription
	sut.UsageText = testUsageText
	sut.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "socket, s",
			Usage: "some usage text",
			Value: "value",
		},
		cli.StringFlag{Name: "flag, fl, f"},
		cli.BoolFlag{
			Name:  "another-flag, b",
			Usage: "another usage text",
		},
	}
	sut.Commands = []cli.Command{{
		Aliases: []string{"c"},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "flag, fl, f"},
			cli.BoolFlag{
				Name:  "another-flag, b",
				Usage: "another usage text",
			},
		},
		Name:  "config",
		Usage: "another usage test",
		Subcommands: []cli.Command{{
			Aliases: []string{"s", "ss"},
			Flags: []cli.Flag{
				cli.StringFlag{Name: "sub-flag, sub-fl, s"},
				cli.BoolFlag{
					Name:  "sub-command-flag, s",
					Usage: "some usage text",
				},
			},
			Name:  "sub-config",
			Usage: "another usage test",
		}},
	}, {
		Aliases: []string{"i", "in"},
		Name:    "info",
		Usage:   "retrieve generic information",
	}, {
		Name: "some-command",
	}}

	return &TestFramework{
		setup,
		teardown,
		sut,
	}
}

// NilFunc is a convenience function which simply does nothing
func NilFunc(f *TestFramework) error {
	return nil
}

// Setup is the global initialization function which runs before each test
// suite
func (t *TestFramework) Setup() {
	// Global initialization for the whole framework goes in here

	// Setup the actual test suite
	gomega.Expect(t.setup(t)).To(gomega.Succeed())
}

// Teardown is the global deinitialization function which runs after each test
// suite
func (t *TestFramework) Teardown() {
	// Global deinitialization for the whole framework goes in here

	// Teardown the actual test suite
	gomega.Expect(t.teardown(t)).To(gomega.Succeed())
}

// Describe is a convenience wrapper around the `ginkgo.Describe` function
func (t *TestFramework) Describe(text string, body func()) bool {
	return ginkgo.Describe("docgen: "+text, body)
}

// RunFrameworkSpecs is a convenience wrapper for running tests
func RunFrameworkSpecs(t *testing.T, suiteName string) {
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, suiteName,
		[]ginkgo.Reporter{reporters.NewJUnitReporter(
			fmt.Sprintf("%v_junit.xml", strings.ToLower(suiteName)))})
}
