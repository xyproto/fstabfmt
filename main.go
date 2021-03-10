package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const versionString = "fstabfmt 1.0.0"

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

func format(data []byte, spaces int) []byte {
	var (
		buf       bytes.Buffer
		nl        = []byte{'\n'}
		longest   = make(map[int]int, 0) // The longest length of a field, for each field index
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
	)
	if len(os.Args) > 2 {
		if os.Args[1] == "-i" {
			filename = os.Args[2]
			modify = true
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
	formatted := format(data, 2) // Separate fields with 2 spaces
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
