package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

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
var visit map[string]bool = make(map[string]bool)

func setDefault(selpg *selpgArgs) {
	pflag.IntVarP(&selpg.startPage, "start", "s", -1, "Start Page")
	pflag.IntVarP(&selpg.endPage, "end", "e", -1, "End Page")
	pflag.IntVarP(&selpg.pageLen, "length", "l", 72, "The count of rows of each page")
	pflag.BoolVarP(&selpg.pageType, "f", "f", false, "How page breaks")
	pflag.StringVarP(&selpg.printDest, "destination", "d", "default", "Destination")
	pflag.Parse()
	progname = os.Args[0]
	pflag.Visit(func(f *pflag.Flag) {
		visit[f.Name] = true
	})
}

func processArgs(ac int, selpg *selpgArgs) {
	if pflag.NArg() > 0 {
		selpg.inFilename = pflag.Arg(0)
		_, err := os.Open(selpg.inFilename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	}
	switch {
	case ac < 3:
		fmt.Fprintf(os.Stderr, "%v: not enough arguments\n", progname)
		pflag.Usage()
		os.Exit(1)
	case !visit["start"]:
		fmt.Fprintf(os.Stderr, "You have to give the start page by -s\n")
		pflag.Usage()
		os.Exit(1)
	case !visit["end"]:
		fmt.Fprintf(os.Stderr, "You have to give the end page by -e\n")
		pflag.Usage()
		os.Exit(1)
	case selpg.startPage <= 0:
		fmt.Fprintf(os.Stderr, "%v: startPage must be bigger than 0\n", progname)
		os.Exit(1)
	case selpg.endPage < selpg.startPage:
		fmt.Fprintf(os.Stderr, "%v: endPage must be bigger than startPage\n", progname)
		os.Exit(1)
	case selpg.pageLen <= 1:
		fmt.Fprintf(os.Stderr, "%v: pageLen must be bigger than 1\n", progname)
		os.Exit(1)
	case visit["f"] && visit["length"]:
		fmt.Fprintf(os.Stderr, "You can't use two type at the same time\n")
		pflag.Usage()
		os.Exit(1)
	}
}

func processInput(selpg *selpgArgs) {
	var inputReader *bufio.Reader
	var outputWriter *bufio.Writer
	var err error
	var file *os.File
	var stdin io.WriteCloser
	var cmd *exec.Cmd

	if selpg.inFilename == "" {
		inputReader = bufio.NewReader(os.Stdin)
	} else {
		file, err = os.Open(selpg.inFilename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		inputReader = bufio.NewReader(file)
	}
	if selpg.printDest == "default" {
		outputWriter = bufio.NewWriter(os.Stdout)
	} else {
		cmd = exec.Command("lp", "-d"+selpg.printDest)
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	var line []byte
	page := 1
	if selpg.pageType {
		for {
			line, err = inputReader.ReadBytes('\f')
			if err != nil {
				break
			}
			if selpg.startPage <= page && page <= selpg.endPage {
				if selpg.printDest == "default" {
					outputWriter.Write(line)
					outputWriter.Flush()
				} else {
					_, err := io.WriteString(stdin, string(line))
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
				}
			}
			page++
		}
	} else {
		lineCount := 0
		for {
			line, err = inputReader.ReadBytes('\n')
			if err != nil {
				break
			}
			lineCount++
			if lineCount > selpg.pageLen {
				lineCount = 1
				page++
			}
			if selpg.startPage <= page && page <= selpg.endPage {
				if selpg.printDest == "default" {
					outputWriter.Write(line)
					outputWriter.Flush()
				} else {
					_, err := io.WriteString(stdin, string(line))
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
				}
			}
		}
	}
	if selpg.printDest != "default" {
		stdin.Close()
		stderr, _ := cmd.CombinedOutput()
		fmt.Fprintln(os.Stderr, string(stderr))
	}
}

func main() {
	selpg := selpgArgs{
		startPage:  -1,
		endPage:    -1,
		pageLen:    72,
		pageType:   false,
		inFilename: "",
		printDest:  "default",
	}
	setDefault(&selpg)
	processArgs(len(os.Args), &selpg)
	processInput(&selpg)
}
