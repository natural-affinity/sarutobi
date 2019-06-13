package hiruzen_test

import (
	"testing"

	"github.com/natural-affinity/sarutobi/hiruzen"
)

func TestTagged(t *testing.T) {
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
