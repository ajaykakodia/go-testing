//go:build unit
// +build unit

package yamltohtml_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ajaykakodia/go-testing/yamltohtml"
)

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestMain(m *testing.M) {
	fmt.Println("Hello we started test")
	ret := m.Run()
	fmt.Println("Test have executed")
	os.Exit(ret)
}

func TestYamltoHTML(t *testing.T) {
	testCases := []TestCase{
		{
			desc:     "Test case 1",
			path:     "testdata/test_01.yaml",
			expected: "<html><head><title>My Awsome Pages</title></head><body>Chal hat yeha se</body></html>",
		},
		{
			desc:     "Test case 2",
			path:     "testdata/test_02.yaml",
			expected: "<html><head><title>My Test Pages</title></head><body>Ye Test hai</body></html>",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			result, err := yamltohtml.YamltoHTML(testCase.path)
			if err != nil {
				t.Fail()
			}
			t.Log(result)

			if result != testCase.expected {
				t.Fail()
			}
		})
	}
}
