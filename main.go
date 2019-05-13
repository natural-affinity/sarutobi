package main

import (
	"fmt"
	"log"
	"os"

	docopt "github.com/docopt/docopt-go"
)

// Version identifier
const Version = "0.0.1"

// Usage message (docopt interface)
const Usage = `
  Sarutobi (Purofessā)
    Inspirational quotes

  Usage:
    sarutobi
    sarutobi --tag t
    sarutobi --author a
    sarutobi --help
    sarutobi --version

  Options:
    -h, --help        display help information
    -v, --version     display version information
    -t, --tag t       limit relevant quotes by tag
    -a, --author a    limit relevant quotes by author
`

func main() {
	log.SetFlags(log.Lshortfile)

	// parse usage string and collect args
	args, err := docopt.ParseArgs(Usage, os.Args[1:], Version)
	if err != nil {
		log.Fatalf("invalid usage string: %s", err.Error())
	}

	// extract options and args
	tag, _ := args.String("--tag")
	author, _ := args.String("--author")

	fmt.Printf("%s_%s", tag, author)
}
