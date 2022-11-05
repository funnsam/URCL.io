package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Tokenize(raw []byte) []Instruction {
	SetOutput("Output of lexer: \n" + fmt.Sprintf("%v", Split(string(raw))))
	tm, _ := Compile(Split(string(raw)))
	ShowDialog("simple", "Compiler: ", fmt.Sprintf("%v", tm))

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

func Compile(raw [][]string) ([]Instruction, error) {
	tmpInstrList := []Instruction{}

	for _, e := range raw {
		instrName := ""
		instrArgs := []uint32{}
		instrType := []ArgType{}

		for j, e2 := range e {
			if j != 0 {
				switch []byte(e2)[0] {
				case 'R', 'r', '$':
					tm, er := strconv.Atoi(e2[1:])
					if er != nil {
						return nil, er
					}
					instrArgs = append(instrArgs, uint32(tm))
					instrType = append(instrType, Register)
				case 'M', 'm', '#':
					tm, er := strconv.Atoi(e2[1:])
					if er != nil {
						return nil, er
					}
					instrArgs = append(instrArgs, uint32(tm))
					instrType = append(instrType, Memory)
				case '%':
					instrArgs = append(instrArgs, uint32(PortsPointer[e2[1:]]))
					instrType = append(instrType, Immediate)
				default:
				}
			} else {
				instrName = e2
			}
		}

		tmpInstrList = append(tmpInstrList, Instruction{InstructionPointers[instrName], instrArgs, instrType})
	}

	return tmpInstrList, nil
}
