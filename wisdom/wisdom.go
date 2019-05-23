package wisdom

import "fmt"

// Library of ultimate truth
type Library struct {
	Default Quote
	Quotes  []Quote
	Tags    map[string]string
}

// Quote of wisdom
type Quote struct {
	Author  string
	Message string
	Tags    map[string]interface{}
}

// Wisdom to share
type Wisdom interface {
	Print()
	Tagged(tags []string) bool
}

// Print wisdom
func (q *Quote) Print() {
	fmt.Printf("\n%s\n \u2014 %s\n\n", q.Message, q.Author)
}

// Tagged with
func (q *Quote) Tagged(tags []string) bool {
	for _, t := range tags {
		if _, ok := q.Tags[t]; !ok {
			return false
		}
	}

	return true
}
