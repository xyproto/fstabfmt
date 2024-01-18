package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const versionString = "fstabfmt 1.2.0"

func usage() {
	fmt.Println(versionString + `
Usage: fstabfmt [-i] [-s NUM] [FILE]

fstabfmt formats /etc/fstab files.
It can either read from stdin and print to stdout
or modify the given file if the -i flag is used.

-h, --help         Display this help
-i                 Modify the given file
-s, --spaces NUM   Specify the number of spaces used between fields
-v, --version      Display the current version

`)
}

func main() {
	var (
		data        []byte
		err         error
		filename    = "-"
		modifyFile  bool
		showVersion bool
		spaces      int
	)

	flag.Usage = usage
	flag.IntVar(&spaces, "s", 2, "")
	flag.IntVar(&spaces, "spaces", 2, "")
	flag.BoolVar(&showVersion, "v", false, "")
	flag.BoolVar(&showVersion, "version", false, "")
	flag.BoolVar(&modifyFile, "i", false, "")
	flag.Parse()

	if showVersion {
		fmt.Println(versionString)
		return
	}

	if len(flag.Arg(0)) > 0 {
		filename = flag.Arg(0)
	}

	if filename == "-" {
		data, err = ioutil.ReadAll(os.Stdin)
	} else {
		data, err = ioutil.ReadFile(filename)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	formatted := format(data, spaces)
	if !modifyFile || filename == "-" {
		fmt.Print(string(formatted))
	} else {
		err = ioutil.WriteFile(filename, formatted, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
