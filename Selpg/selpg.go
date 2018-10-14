package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type selpgArgs struct {
	startPage  int
	endPage    int
	pageLen    int
	pageType   bool
	inFilename string
	printDest  string
}

var progname string

func setDefault(selpg *selpgArgs) {
	pflag.IntVarP(&selpg.startPage, "start", "s", -1, "Start Page")
	pflag.IntVarP(&selpg.endPage, "end", "e", -1, "End Page")
	pflag.IntVarP(&selpg.pageLen, "length", "l", 72, "The count of rows of each page")
	pflag.BoolVarP(&selpg.pageType, "f", "f", false, "How page breaks")
	pflag.StringVarP(&selpg.printDest, "destination", "d", "", "Destination")
	pflag.Parse()
	progname = os.Args[0]
}

func processArgs(ac int, selpg *selpgArgs) {
	if ac < 3 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		os.Exit(1)
	}

	switch {
	case selpg.startPage <= 0:
		fmt.Fprintf(os.Stderr, "%v: startPage should be bigger than 0\n", progname)
		os.Exit(1)
	case selpg.endPage < selpg.startPage:
		fmt.Fprintf(os.Stderr, "%v: endPage should be bigger than startPage\n", progname)
		os.Exit(1)
	case selpg.pageLen <= 1:
		fmt.Fprintf(os.Stderr, "%v: pageLen should be bigger than 1\n", progname)
		os.Exit(1)
	}
}

func processInput(selpg *selpgArgs) {

}

func main() {
	selpg := selpgArgs{
		startPage:  -1,
		endPage:    -1,
		pageLen:    72,
		pageType:   false,
		inFilename: "0",
		printDest:  "default",
	}
	setDefault(&selpg)
	processArgs(len(os.Args), &selpg)
	processInput(&selpg)
}
