package main

import (
	"bufio"
	"flag"
	"fmt"
	//	"log"
	"os"
	"regexp"
	"strings"
)

var (
	fileViewFlag bool
	stdViewFlag  bool
	appendFlag   bool
)

func init() {
	flag.BoolVar(&fileViewFlag, "file", false, "file view output flag")
	flag.BoolVar(&stdViewFlag, "std", true, "stdout & stderr view output flag")
	flag.BoolVar(&appendFlag, "b", false, "Append the output to the files rather than overwriting them.")
	//	TODO
	// -i      Ignore the SIGINT signal.
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Printf("no output filepath: %s filepath\n", strings.Join(os.Args, " "))
		os.Exit(1)
	}
	name := args[0]
	defaultModeFlag := os.O_CREATE | os.O_RDWR | os.O_TRUNC
	appendModeFlag := os.O_CREATE | os.O_RDWR | os.O_APPEND
	fileOpenModeFlag := defaultModeFlag
	if appendFlag {
		defaultModeFlag = appendModeFlag
	}
	fp, err := os.OpenFile(name, fileOpenModeFlag, 0644)
	if err != nil {
		fmt.Printf("input file create err:%s\n", err)
		os.Exit(1)
	}
	defer fp.Close()

	ansiContSeqReg := regexp.MustCompilePOSIX(`(\x9B|\x1B\[)[0-?]*[ -\/]*[@-~]`)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		if fileViewFlag {
			fmt.Fprint(fp, text)
		}
		if stdViewFlag {
			fmt.Fprint(os.Stdout, text)
		}

		if fileViewFlag && stdViewFlag {
			continue
		}

		//	remove "\r"
		index := strings.LastIndex(text, "\r")
		index += len("\r")
		text = text[index:]
		//	remove "ansiContSeqReg control sequences"
		text = ansiContSeqReg.ReplaceAllString(text, "")

		if !fileViewFlag {
			fmt.Fprint(fp, text)
		}
		if !stdViewFlag {
			fmt.Fprint(os.Stdout, text)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("stdin read err:%s\n", err)
		os.Exit(1)
	}
}
