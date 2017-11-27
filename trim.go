package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	left := flag.Bool("l", true, "Trim left of the string")
	right := flag.Bool("r", false, "Trim right of the string")

	flag.Parse()

	pattern := flag.Arg(0)

	s := flag.Arg(1)

	if piped() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			s := scanner.Text()
			s = trim(s, pattern, *left, *right)
			fmt.Fprintf(os.Stdout, s+"\n")
		}
	} else if s != "" {
		fmt.Fprintf(os.Stdout, trim(s, pattern, *left, *right)+"\n")
	} else {
		os.Exit(1)
	}
}

func trim(s, pattern string, left, right bool) string {
	if left {
		s = strings.TrimPrefix(s, pattern)
	}
	if right {
		s = strings.TrimSuffix(s, pattern)
	}
	return s
}

func piped() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	}
	return true
}
