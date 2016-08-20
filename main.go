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
	var outputFile string
	flag.StringVar(&outputFile, "out", "", "The file to write output to.")
	flag.Parse()

	in, inFile, err := getInput()
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

	doc, err := parser.Parse(inFile, string(src))
	if err != nil {
		panic(err)
	}

	out, err := getOutput(outputFile)
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

func getInput() (io.ReadCloser, string, error) {
	if flag.NArg() == 0 {
		return os.Stdin, "", nil
	} else if flag.NArg() == 1 {
		file, err := os.Open(flag.Arg(0))
		return file, flag.Arg(0), err
	}

	flag.Usage()
	os.Exit(0)
	return nil, "", flag.ErrHelp
}

func getOutput(outputFile string) (io.WriteCloser, error) {
	if outputFile == "" {
		return os.Stdout, nil
	}
	return os.OpenFile(outputFile, os.O_WRONLY, 0644)
}
