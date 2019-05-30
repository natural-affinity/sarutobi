package hiruzen

import (
	"github.com/fatih/color"
)

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
	Tagged(tags ...string) bool
}

// Shintai tags
type Shintai interface {
	PrintTags()
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

// PrintTags of library
func (l *Library) PrintTags() {
	for k := range l.Tags {
		color.HiGreen(k)
	}
}
