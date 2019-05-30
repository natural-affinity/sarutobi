package hiruzen

import (
	"github.com/fatih/color"
)

// Tags for filtering
type Tags map[string]interface{}

// Library of ultimate truth
type Library struct {
	Default Quote
	Quotes  []Quote
	Tags    Tags
}

// Quote of wisdom
type Quote struct {
	Author  string
	Message string
	Tags    Tags
}

// Wisdom to share
type Wisdom interface {
	Printer
	Tagged(tags ...string) bool
}

// Printer of things
type Printer interface {
	Print()
}

// Tagged with
func (q *Quote) Tagged(tags ...string) bool {
	for _, t := range tags {
		if _, ok := q.Tags[t]; !ok {
			return false
		}
	}

	return true
}

// Print wisdom
func (q *Quote) Print() {
	color.HiYellow("\n%s\n", q.Message)
	color.HiRed(" \u2014 %s\n\n", q.Author)
}

// Print tags
func (t Tags) Print() {
	for k := range t {
		color.HiGreen(k)
	}
}
