package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/R0X4R/nbeam/pkg/filter"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	peekBytes, err := reader.Peek(1)
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Error reading input stream:", err)
		os.Exit(1)
	}

	if len(peekBytes) == 0 {
		return
	}

	var fullStream io.Reader = reader

	// Check if the input looks like XML (starts with '<')
	if peekBytes[0] == '<' {
		filter.NmapXML(fullStream)
	} else {
		filter.NmapTEXT(fullStream)
	}
}
