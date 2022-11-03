package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
)

func Tokenize(raw []byte) []Instruction {
	println(fmt.Sprintf("%#v", Split(string(raw))))

	return []Instruction{}
}

func Split(raw string) [][]string {
	scanner := bufio.NewScanner(strings.NewReader(raw))
	buf := make([][]string, 0, 100)
	for scanner.Scan() {
		currln := []byte(scanner.Text())

		startOfLn := true
		isSpaceRepeating := false
		currlnlist := make([]string, 0, 10)
		currlnbuf := ""

		for _, element := range currln {
			switch element {
			case ' ', '\t':
				if startOfLn || isSpaceRepeating {
					continue
				} else {
					currlnlist = append(currlnlist, currlnbuf)
					currlnbuf, isSpaceRepeating = "", true
				}
			default:
				currlnbuf += string(element)
				startOfLn, isSpaceRepeating = false, false
			}
		}
		currlnlist = append(currlnlist, currlnbuf)
		if !reflect.DeepEqual(currlnlist, []string{""}) {
			buf = append(buf, currlnlist)
		}
	}

	return buf
}
