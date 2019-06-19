package hiruzen_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"testing"

	"github.com/natural-affinity/gotanda"
	"github.com/natural-affinity/sarutobi/hiruzen"
)

var update = flag.Bool("update", false, "update .golden files")

func TestTaggedQuote(t *testing.T) {
	cases := []struct {
		Name       string
		Tags       hiruzen.Tags
		SearchTags []string
		Expected   bool
	}{
		{"nil", nil, nil, true},
		{"nil.tags", nil, []string{}, true},
		{"nil.search", hiruzen.Tags{"Tag1": nil}, nil, true},
		{"empty", hiruzen.Tags{}, []string{}, true},
		{"empty.search", hiruzen.Tags{"Tag1": nil}, []string{}, true},
		{"no.match.single", hiruzen.Tags{"Tag1": nil}, []string{"Tag2"}, false},
		{"no.match.multi.search", hiruzen.Tags{"Tag1": nil}, []string{"Tag1", "Tag2"}, false},
		{"match.single.search", hiruzen.Tags{"Tag1": nil, "Tag2": nil}, []string{"Tag2"}, true},
		{"match.multi.search", hiruzen.Tags{"Tag1": nil, "Tag2": nil}, []string{"Tag1", "Tag2"}, true},
	}

	for _, tc := range cases {
		q := &hiruzen.Quote{Tags: tc.Tags}
		actual := q.Tagged(tc.SearchTags...)

		if actual != tc.Expected {
			t.Errorf("Test: %s\n Expected: %t\n Actual: %t\n", tc.Name, tc.Expected, actual)
		}
	}
}

func TestPrintQuote(t *testing.T) {
	cases := []struct {
		Name  string
		Quote *hiruzen.Quote
	}{
		{"print.empty.quote", &hiruzen.Quote{}},
		{"print.author.quote", &hiruzen.Quote{Author: "Toyo"}},
		{"print.quote", &hiruzen.Quote{Message: "Hello World", Author: "Brian K."}},
	}

	for _, tc := range cases {
		golden, expected := gotanda.LoadTestFile(t, "../testdata", tc.Name+".golden")
		abyte, _ := gotanda.Capture(func() {
			tc.Quote.Print()
		})

		if *update {
			ioutil.WriteFile(golden, abyte, 0644)
			expected, _ = ioutil.ReadFile(golden)
		}

		if !bytes.Equal(expected, abyte) {
			t.Errorf("Test: %s\n Expected: %s\n Actual: %s\n", tc.Name, expected, abyte)
		}
	}
}

func TestPrintTags(t *testing.T) {
	cases := []struct {
		Name string
		Tags hiruzen.Tags
	}{
		{"print.empty.tags", hiruzen.Tags{}},
		{"print.single.key.tags", hiruzen.Tags{"Tag1": nil}},
		{"print.multi.keyvalue.tags", hiruzen.Tags{"Tag1": "A", "spaced tag": "B"}},
		{"print.multi.mixed.tags", hiruzen.Tags{"Tag1": "A", "Tag2": nil}},
	}

	for _, tc := range cases {
		golden, expected := gotanda.LoadTestFile(t, "../testdata", tc.Name+".golden")
		abyte, _ := gotanda.Capture(func() {
			tc.Tags.Print()
		})

		if *update {
			ioutil.WriteFile(golden, abyte, 0644)
			expected, _ = ioutil.ReadFile(golden)
		}

		if !bytes.Equal(expected, abyte) {
			t.Errorf("Test: %s\n Expected: %s\n Actual: %s\n", tc.Name, expected, abyte)
		}
	}
}
