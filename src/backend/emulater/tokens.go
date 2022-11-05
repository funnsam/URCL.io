package main

import "errors"

type Instruction struct {
	ID      uint16
	Arg     []uint32
	ArgType []ArgType
}

type ArgType int

const (
	Register ArgType = iota
	Immediate
	Memory
)

type DefineInstruction struct {
	ArgCount int                            // -1 is for any
	OnCall   func([]Instruction, int) error // OnCall(AllInstr, CurrentIndex) error
}

var InstructionPointers = make(map[string]uint16)
var InstructionList = make(map[int]DefineInstruction)
var NextInstrID = 1

const (
	TokenINVALID uint32 = iota
	TokenIN
	TokenOUT
	TokenADD
	TokenRSH
	TokenLOD
	TokenSTR
	TokenBGE
	TokenNOR
	TokenMOV // IMM will be translated to MOV
	TokenSUB
	TokenJMP // NOP will be replaced by MOV R0 R0
	TokenLSH
	TokenINC
	TokenDEC
	TokenNEG
	TokenAND
	TokenOR // NOT will be replaced by NOR R? R0
	TokenXOR
	TokenXNOR
	TokenNAND
	TokenBRL
	TokenBRG
	TokenBNE
	TokenBOD
	TokenBEV
	TokenBLE
	TokenBRZ
	TokenBRN
	TokenBRP
	TokenPSH
	TokenPOP
	TokenCAL
	TokenRET
	TokenHLT
	TokenCPY
	TokenBRC
	TokenBNC
)

func RegisterURCLInstruction(name string, fnc func([]Instruction, int) error, argv int) int {
	InstructionPointers[name] = uint16(NextInstrID)
	InstructionList[NextInstrID] = DefineInstruction{argv, fnc}
	NextInstrID++
	return NextInstrID
}

func OnInvalidInstruction([]Instruction, int) error {
	return errors.New("urcl compiler: Syntax error or an invalid instruction has been registered")
}
