package main

import (
	"flag"
	"fmt"
)

const name = "eolctl"

var (
	printHelp    *bool
	printVersion *bool
)

func main() {
	initFlags()
	flag.Parse()

	if *printHelp {
		printUsage()
		return
	}
	if *printVersion {
		fmt.Fprintln(flag.CommandLine.Output(), name, Version)
		return
	}
}

func initFlags() {
	flag.CommandLine.Usage = printUsage
	printHelp = flag.Bool("h", false, "Print this usage note")
	printVersion = flag.Bool("v", false, "Print version info")
}

func printUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s\n\nOptions:\n", synopsis())
	flag.CommandLine.PrintDefaults()
}

func synopsis() string {
	var allName string
	flag.VisitAll(func(flag *flag.Flag) {
		allName += flag.Name
	})

	return fmt.Sprintf("%s [-%s]", name, allName)
}
