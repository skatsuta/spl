package main

import (
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
	nonum                bool
	delim                string
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&cli.nonum, "nonum", false, "hide line numbers")
	flags.StringVar(&cli.delim, "delim", ":", "a delimiter that separates elements of an argument")

	flVersion := flags.Bool("version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if *flVersion {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if e := cli.split(flags.Args()); e != nil {
		fmt.Fprintf(cli.errStream, "Error splitting %s: %s\n", flag.Args(), e)
		return ExitCodeError
	}

	return ExitCodeOK
}

// split splits each element of args with delim and write to out.
func (cli *CLI) split(args []string) error {
	for _, arg := range args {
		elems := strings.Split(arg, cli.delim)
		// width of line numbers
		width := len(elems)/10 + 1

		for i, p := range elems {
			if !cli.nonum {
				if _, e := fmt.Fprintf(cli.outStream, "%"+strconv.Itoa(width)+"d  ", i+1); e != nil {
					return e
				}
			}
			if _, e := fmt.Fprintf(cli.outStream, "%s\n", p); e != nil {
				return e
			}
		}
		// new line
		if _, e := fmt.Fprintln(cli.outStream); e != nil {
			return e
		}
	}

	return nil
}
