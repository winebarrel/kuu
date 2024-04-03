package main

import (
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/winebarrel/kuu"
)

var (
	version string
)

type Options struct {
	Files []string `arg:"" optional:"" help:"Input files."`
	kuu.Options
}

func init() {
	log.SetFlags(0)
}

func parseArgs() *Options {
	var CLI struct {
		Options
		Version kong.VersionFlag
	}

	parser := kong.Must(&CLI, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	return &CLI.Options
}

func main() {
	options := parseArgs()
	err := kuu.Filter(options.Files, &options.Options)

	if err != nil {
		log.Fatal(err)
	}
}
