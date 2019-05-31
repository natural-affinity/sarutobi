package main

import (
	"log"
	"os"

	"github.com/natural-affinity/sarutobi/hiruzen"

	docopt "github.com/docopt/docopt-go"
)

// Version identifier
const Version = "0.0.1"

// Shintai library
const Shintai = "data/shintai.yaml"

// Usage message (docopt interface)
const Usage = `
  Sarutobi (Purofessā)
    Inspirational quotes

  Usage:
    sarutobi [<tag>...]
    sarutobi --tags
    sarutobi --help
    sarutobi --version

  Options:
    -t, --tags        display tags
    -h, --help        display help information
    -v, --version     display version information
`

func main() {
	log.SetFlags(log.Lshortfile)

	// parse usage string and collect args
	args, err := docopt.ParseArgs(Usage, os.Args[1:], Version)
	if err != nil {
		log.Fatalf("invalid usage string: %s", err.Error())
	}

	// extract options and args
	subject := func(q hiruzen.Quote) bool {
		return q.Tagged(args["<tag>"].([]string)...)
	}

	shintai, err := Asset(Shintai)
	if err != nil {
		log.Fatalf("invalid asset: %s", err.Error())
	}

	knowledge, err := hiruzen.Recall(shintai)
	if err != nil {
		log.Fatalf("invalid data: %s", err.Error())
	}

	// display quotes or tags
	switch {
	case args["--tags"].(bool):
		knowledge.Tags.Print()
	default:
		sensei := &hiruzen.Sensei{Knowledge: knowledge}
		quotes := sensei.Advise(subject)
		sensei.Inspire(quotes).Print()
	}
}
