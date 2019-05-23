package hiruzen

import (
	"time"

	"github.com/natural-affinity/sarutobi/wisdom"
	"gopkg.in/yaml.v3"
)

// Shintai of hiruzen
const Shintai = "wisdom/shintai.yaml"

// DefaultQuote d
var DefaultQuote = &wisdom.Quote{Message: "Test", Author: "None", Tags: nil}

// Sensei with knowledge
type Sensei struct {
	knowledge *wisdom.Library
}

// Subject of relevance
type Subject func(q wisdom.Quote) bool

// Professor provides wisdom
type Professor interface {
	Recall(shintai string) error
	Advise(relevant Subject) ([]wisdom.Quote, error)
	Summarize(wisdom []wisdom.Quote) *wisdom.Quote
}

// Recall shintai
func (s *Sensei) Recall(shintai string) error {
	knowledge, err := wisdom.Asset(shintai)
	if err != nil {
		return err
	}

	lib := &wisdom.Library{}
	if err := yaml.Unmarshal(knowledge, lib); err != nil {
		return err
	}

	s.knowledge = lib
	return nil
}

// Advise with relevant wisdom
func (s *Sensei) Advise(relevant Subject) ([]wisdom.Quote, error) {
	if s.knowledge == nil {
		if err := s.Recall(Shintai); err != nil {
			return nil, err
		}
	}

	var related []wisdom.Quote
	for _, wisdom := range s.knowledge.Quotes {
		if relevant(wisdom) {
			related = append(related, wisdom)
		}
	}

	return related, nil
}

// Summarize wisdom in a single quote
func (s *Sensei) Summarize(wisdom []wisdom.Quote) *wisdom.Quote {
	max := int64(len(wisdom))
	if max == 0 {
		return DefaultQuote
	}

	rnd := time.Now().Unix() % max
	return &wisdom[rnd]
}
