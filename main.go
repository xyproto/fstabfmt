package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const versionString = "fstabfmt 1.0.0"

func usage() {
	fmt.Println(versionString + `
Usage: fstabfmt [-i FILE]

fstabfmt formats /etc/fstab files.
It can either read from stdin and print to stdout
or modify the given file if the -i flag is used.

-h, --help        Display this help
-v, --version     Display the current version
-s, --spaces NUM  Specify the number of spaces used between fields 
-i FILE           Supply a file that will be modified

`)
}

func format(data []byte, spaces int) []byte {
	var (
		buf       bytes.Buffer
		nl        = []byte{'\n'}
		longest   = make(map[int]int) // The longest length of a field, for each field index
		byteLines = bytes.Split(data, nl)
	)

	// Find the longest field length for each field on each line
	for _, line := range byteLines {
		trimmedLine := bytes.TrimSpace(line)
		if len(trimmedLine) == 0 || bytes.HasPrefix(trimmedLine, []byte{'#'}) {
			continue
		}
		// Find the longest field length for each field
		for i, field := range bytes.Fields(trimmedLine) {
			fieldLength := len(string(field))
			if val, ok := longest[i]; ok {
				if fieldLength > val {
					longest[i] = fieldLength
				}
			} else {
				longest[i] = fieldLength
			}
		}
	}

	// Format the lines nicely
	for _, line := range byteLines {
		trimmedLine := bytes.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue
		}
		if bytes.HasPrefix(trimmedLine, []byte{'#'}) { // Output comments as they are, but trimmed
			buf.Write(trimmedLine)
			buf.Write(nl)
		} else { // Format the fields
			for i, field := range bytes.Fields(trimmedLine) {
				fieldLength := len(string(field))
				padCount := spaces // Space between the fields if all fields have equal length
				if longest[i] > fieldLength {
					padCount += longest[i] - fieldLength
				}
				buf.Write(field)
				if padCount > 0 {
					buf.Write(bytes.Repeat([]byte{' '}, padCount))
				}
			}
			buf.Write(nl)
		}
	}
	return buf.Bytes()
}

func main() {
	var (
		data     []byte
		err      error
		filename = "-"
		modify   bool
		spaces   = 2
	)

	skipNextArg := false
	for i := range os.Args {
		if skipNextArg {
			skipNextArg = false
			continue
		}
		switch os.Args[i] {
		case "-v", "--version":
			fmt.Println(versionString)
			return
		case "-h", "--help":
			usage()
			return
		case "-i":
			filename = os.Args[i+1]
			modify = true
			skipNextArg = true
		case "-s", "--spaces":
			spaces, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: invalid number of spaces")
				os.Exit(1)
			}
			skipNextArg = true
		default:
			filename = os.Args[i]
		}
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
	if !modify || filename == "-" {
		fmt.Print(string(formatted))
	} else {
		err = ioutil.WriteFile(filename, formatted, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
