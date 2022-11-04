package main

import "errors"

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

type DefineInstruction struct {
	ArgCount int                            // 0 is for any
	OnCall   func([]Instruction, int) error // OnCall(AllInstr, CurrentIndex) error
}

var InstructionPointers = make(map[string]uint16)
var InstructionList = make(map[uint16]DefineInstruction)

var (
	TokenIN   = 1
	TokenOUT  = 2
	TokenADD  = 3
	TokenRSH  = 4
	TokenLOD  = 5
	TokenSTR  = 6
	TokenBGE  = 7
	TokenNOR  = 8
	TokenMOV  = 9 // IMM will be translated to MOV
	TokenSUB  = 10
	TokenJMP  = 11 // NOP will be replaced by MOV R0 R0
	TokenLSH  = 12
	TokenINC  = 13
	TokenDEC  = 14
	TokenNEG  = 15
	TokenAND  = 16
	TokenOR   = 17 // NOT will be replaced by NOR R? R0
	TokenXOR  = 18
	TokenXNOR = 19
	TokenNAND = 20
	TokenBRL  = 21
	TokenBRG  = 22
	TokenBNE  = 23
	TokenBOD  = 24
	TokenBEV  = 25
	TokenBLE  = 26
	TokenBRZ  = 27
	TokenBRN  = 28
	TokenBRP  = 29
	TokenPSH  = 30
	TokenPOP  = 31
	TokenCAL  = 32
	TokenRET  = 33
	TokenHLT  = 34
	TokenCPY  = 35
	TokenBRC  = 36
	TokenBNC  = 37
)

func init() {
	InstructionList[0] = DefineInstruction{0, func([]Instruction, int) error {
		return errors.New("Syntax error or an invalid instruction has been registered")
	}}
}
