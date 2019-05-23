package hiruzen

import (
	"time"

	"github.com/natural-affinity/sarutobi/wisdom"
	"gopkg.in/yaml.v3"
)

// Shintai of hiruzen
const Shintai = "wisdom/shintai.yaml"

// Sensei with knowledge
type Sensei struct {
	Knowledge *wisdom.Library
}

// Professor advises
type Professor interface {
	Advise(relevant func(q wisdom.Quote) bool) wisdom.Quote
}

// Recall universal truths
func Recall(fp string) (*wisdom.Library, error) {
	if fp == "" {
		fp = Shintai
	}

	asset, err := wisdom.Asset(fp)
	if err != nil {
		return nil, err
	}

	lib := &wisdom.Library{}
	if err := yaml.Unmarshal(asset, lib); err != nil {
		return nil, err
	}

	return lib, nil
}

// Advise with relevant wisdom
func (s *Sensei) Advise(relevant func(q wisdom.Quote) bool) wisdom.Quote {
	var r []wisdom.Quote
	for _, q := range s.Knowledge.Quotes {
		if relevant(q) {
			r = append(r, q)
		}
	}

	max := int64(len(r))
	if max == 0 {
		return s.Knowledge.Default
	}

	rand := time.Now().Unix() % max
	return r[rand]
}
