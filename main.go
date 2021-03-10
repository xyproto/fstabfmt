package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const versionString = "fstabfmt 0.1.0"

func usage() {
	fmt.Println(versionString)
	fmt.Println()
	fmt.Println("Usage: fstabfmt [-i FILE]")
	fmt.Println()
	fmt.Println("fstabfmt formats /etc/fstab files.")
	fmt.Println("It can either read from stdin and print to stdout")
	fmt.Println("or modify the given file if the -i flag is used.")
	fmt.Println()
	fmt.Println("-h, --help       Display this help")
	fmt.Println("-v, --version    Display the current version")
	fmt.Println("-i FILE          Supply a file that will be modified")
	fmt.Println()
}

func main() {
	var (
		data     []byte
		err      error
		filename = "-"
	)

	fmt.Println(os.Args)

	if len(os.Args) > 2 {
		if os.Args[1] == "-i" {
			filename = os.Args[2]

		} else {
			usage()
			os.Exit(1)
		}
	} else if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(versionString)
			os.Exit(0)
		case "-h", "--help":
			usage()
			os.Exit(0)
		default:
			filename = os.Args[1]
		}
	} else {
		usage()
		os.Exit(1)
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

	fmt.Println("DATA")
	fmt.Println(string(data))
}
