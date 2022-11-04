package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
)

func Tokenize(raw []byte) []Instruction {
	ShowDialog("simple", "Output of lexer", fmt.Sprintf("%v", Split(string(raw))))

	return []Instruction{}
}

func Split(raw string) [][]string {
	scanner := bufio.NewScanner(strings.NewReader(raw))
	buf := make([][]string, 0, 100)
	for scanner.Scan() {
		currln := []byte(scanner.Text())

		startOfLn := true
		inQuote := false
		isSpaceRepeating := false
		currlnlist := make([]string, 0, 10)
		currlnbuf := ""

	cln:
		for i, element := range currln {
			switch element {
			case ' ', '\t':
				if startOfLn || isSpaceRepeating {
					continue
				} else {
					if !inQuote {
						currlnlist = append(currlnlist, currlnbuf)
						currlnbuf, isSpaceRepeating = "", true
					} else {
						currlnbuf += string(element)
						startOfLn, isSpaceRepeating = false, false
					}
				}

			case '"', '\'':
				if !inQuote {
					inQuote = true
				} else {
					currlnbuf = string(element) + currlnbuf + string(element)
					startOfLn, isSpaceRepeating, inQuote = false, false, false
				}

			case '/':
				if len(currln) <= i+1 {
					currlnbuf += string(element)
					startOfLn, isSpaceRepeating = false, false
				} else if currln[i+1] != '/' {
					currlnbuf += string(element)
					startOfLn, isSpaceRepeating = false, false
				} else {
					break cln
				}

			default:
				currlnbuf += string(element)
				startOfLn, isSpaceRepeating = false, false
			}
		}
		if currlnbuf != "" {
			currlnlist = append(currlnlist, currlnbuf)
		}
		if !(reflect.DeepEqual(currlnlist, []string{""}) || reflect.DeepEqual(currlnlist, []string{})) {
			buf = append(buf, currlnlist)
		}
	}

	return buf
}
