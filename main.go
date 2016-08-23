package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/remexre/gotl/parser"
)

func main() {
	args := parseArgs()

	in, err := args.getInput()
	if err != nil {
		panic(err)
	}

	src, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}
	err = in.Close()
	if err != nil {
		panic(err)
	}

	doc, err := parser.Parse(args.InputFile, string(src))
	if err != nil {
		panic(err)
	}

	out, err := args.getOutput()
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprint(out, doc.Template())
	if err != nil {
		panic(err)
	}
	err = out.Close()
	if err != nil {
		panic(err)
	}
}

// Args is a type for the arguments needed.
type Args struct {
	InputFile  string
	OutputFile string
}

func parseArgs() *Args {
	a := new(Args)
	flag.StringVar(&a.OutputFile, "out", "", "The file to write output to.")
	flag.Parse()
	if flag.NArg() == 1 {
		a.InputFile = flag.Arg(0)
	} else if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(-1)
	}
	return a
}

func (a *Args) getInput() (io.ReadCloser, error) {
	if a.InputFile == "" {
		return os.Stdin, nil
	}
	file, err := os.Open(flag.Arg(0))
	return file, err
}

func (a *Args) getOutput() (io.WriteCloser, error) {
	if a.OutputFile == "" {
		return os.Stdout, nil
	}
	return os.OpenFile(a.OutputFile, os.O_WRONLY, 0644)
}
