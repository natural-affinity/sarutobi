package hiruzen

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
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
	Message string
	Author  string
	Tags    Tags
}

// Printer of things
type Printer interface {
	Print()
}

// Wisdom to share
type Wisdom interface {
	Printer
	Tagged(tags ...string) bool
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
	w := colorable.NewColorableStdout()
	fmt.Fprintf(w, "\n%s\n", color.HiYellowString(q.Message))
	fmt.Fprintf(w, " %s %s\n\n", color.HiRedString("\u2014"), color.HiRedString(q.Author))
}

// Print tags
func (t Tags) Print() {
	w := colorable.NewColorableStdout()
	for k := range t {
		fmt.Fprintf(w, "%s\n", color.HiGreenString(k))
	}
}
