package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"

	"github.com/remexre/gotl/parser"
	"github.com/remexre/gotl/transforms"
)

func main() {
	args := parseArgs()
	if err := args.process(); err != nil {
		panic(err)
	}
}

func compile(filename, src string) (string, error) {
	doc, err := parser.Parse(filename, src)
	if err != nil {
		return "", err
	}
	doc, err = transforms.Apply(doc)
	if err != nil {
		return "", err
	}
	return doc.Template(), nil
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

func (a *Args) process() error {
	in, err := a.getInput()
	if err != nil {
		return err
	}
	defer in.Close()

	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	html, err := compile(a.InputFile, string(src))
	if err != nil {
		return err
	}

	out, err := a.getOutput()
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.WriteString(out, html)
	return err
}
