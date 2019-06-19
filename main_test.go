package main_test

import (
	"flag"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/natural-affinity/gotanda"
)

var update = flag.Bool("update", false, "update .golden files")

func TestUsage(t *testing.T) {
	cases := []struct {
		Name string
	}{
		{"help.long"},
		{"help.short"},
		{"version.long"},
		{"version.short"},
		{"invalid.tag.multi"},
		{"invalid.tag.single"},
	}

	for _, tc := range cases {
		r := gotanda.CompareCommand(t, tc, update)
		r.Assert(t, tc)
	}
}

func TestTags(t *testing.T) {
	cases := []struct {
		Name string
	}{
		{"tags.long"},
		{"tags.short"},
	}

	for _, tc := range cases {
		_, command := gotanda.LoadTestFile(t, "testdata", tc.Name+".input")
		golden, expected := gotanda.LoadTestFile(t, "testdata", tc.Name+".golden")
		aout, _ := gotanda.Run(string(command))

		if *update {
			ioutil.WriteFile(golden, aout, 0644)
			expected, _ = ioutil.ReadFile(golden)
		}

		// compare lists of strings (due to random map order)
		estrings := strings.Split(string(expected), "\n")
		sort.Strings(estrings)

		astrings := strings.Split(string(aout), "\n")
		sort.Strings(astrings)

		if len(astrings) != len(estrings) {
			t.Errorf("Test: %s\n Expected: %s\n Actual: %s\n", tc.Name, estrings, astrings)
		}

		for i, v := range estrings {
			if astrings[i] != v {
				t.Errorf("Test: %s\n Expected: %s\n Actual: %s\n", tc.Name, v, astrings[i])
			}
		}
	}
}

func TestQuote(t *testing.T) {
	cases := []struct {
		Name string
	}{
		{"random.quote"},
		{"random.quote.multi.tag"},
		{"random.quote.single.tag"},
	}

	for _, tc := range cases {
		_, command := gotanda.LoadTestFile(t, "testdata", tc.Name+".input")
		aout, _ := gotanda.Run(string(command))

		re := regexp.MustCompile(`\n(.*)\n\s\p{Pd}\s(.*)\n\n`)
		out := !re.Match(aout)

		if out {
			t.Errorf("Test: %s\n Expected: %s\n Actual: %s\n", tc.Name, aout, re.String())
		}
	}
}
