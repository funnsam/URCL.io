package main

type Instruction struct {
	ID      uint16
	Arg     []uint32
	ArgType []ArgType
}

type ArgType uint8

const (
	Register  = 0
	Immediate = 1
)
