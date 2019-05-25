package hiruzen

import (
	"time"

	"github.com/natural-affinity/sarutobi/wisdom"
	"gopkg.in/yaml.v3"
)

// Sensei with knowledge
type Sensei struct {
	Knowledge *Library
}

// Professor advises
type Professor interface {
	Advise(relevant func(q Quote) bool) []Quote
	Inspire(quotes []Quote) *Quote
}

// Recall universal truths
func Recall(fp string) (*Library, error) {
	asset, err := wisdom.Asset(fp)
	if err != nil {
		return nil, err
	}

	lib := &Library{}
	if err := yaml.Unmarshal(asset, lib); err != nil {
		return nil, err
	}

	return lib, nil
}

// Advise on topics
func (s *Sensei) Advise(relevant func(q Quote) bool) []Quote {
	var r []Quote
	for _, q := range s.Knowledge.Quotes {
		if relevant(q) {
			r = append(r, q)
		}
	}

	return r
}

// Inspire with wisdom
func (s *Sensei) Inspire(quotes []Quote) *Quote {
	max := int64(len(quotes))
	if max == 0 {
		return &s.Knowledge.Default
	}

	r := time.Now().Unix() % max
	return &quotes[r]
}
