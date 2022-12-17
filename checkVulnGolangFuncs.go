package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Chat GTP madness continues

// List of known vulnerable Go functions
var vulnerableFunctions = []string{
	"strings.Repeat",
	"strings.Replace",
	"bytes.Buffer.Write",
	"bytes.Buffer.WriteString",
	"fmt.Fprint",
	"fmt.Fprintf",
	"fmt.Fprintln",
	"fmt.Print",
	"fmt.Printf",
	"fmt.Println",
	"log.Print",
	"log.Printf",
	"log.Println",
	"os.Create",
	"os.NewFile",
	"os.Open",
	"os.OpenFile",
	"strconv.Atoi",
	"strconv.ParseInt",
	"strconv.ParseUint",
	"strconv.Unquote",
	"template.ParseFiles",
	"template.ParseGlob",
	"template.ParseHTML",
	"template.ParseString",
}

func main() {
	// Read the source code file
	file, err := os.Open("source.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Scan the source code line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Check if any of the vulnerable functions are used
		for _, f := range vulnerableFunctions {
			if strings.Contains(line, f) {
				fmt.Printf("Vulnerable function '%s' used in line: %s\n", f, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

