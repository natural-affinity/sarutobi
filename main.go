package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	docopt "github.com/docopt/docopt-go"
)

// Version identifier
const Version = "0.0.1"

// Usage message (docopt interface)
const Usage = `
  Sarutobi
    Inspirational quotes

  Usage:
    sarutobi [--author a | --tag t] <file>
    sarutobi --help
    sarutobi --version

  Options:
    -h, --help        display help information
    -v, --version     display version information
    -a, --author a    limit relevant quotes by author
    -t, --tag t       limit relevant quotes by tag
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
	research, err := ioutil.ReadFile(args["<file>"].(string))
	if err != nil {
		log.Fatalf("invalid file: %s", err.Error())
	}

	fmt.Printf("%s_%s_%s", tag, author, research)
}
