package filter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func NmapTEXT(input io.Reader) {
	re := regexp.MustCompile(`^(\d+)/tcp\s+open`)

	var host string
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "Nmap scan report") {
			parts := strings.Fields(line)
			if len(parts) >= 5 {
				host = parts[4]
			}
			
		} else if strings.Contains(line, "open") {
			match := re.FindStringSubmatch(line)
			if len(match) > 1 && host != "" {
				port := match[1]
				fmt.Printf("%s:%s\n", host, port)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error scanning text input:", err)
	}
}
