package main

import "strings"

var (
	PortsPointer = make(map[string]uint32)
	PortsUsages  = make(map[uint32]Port)
)

type Port struct {
	OnInput  func([]uint32, uint32) error // OnInput(Regs, TargetReg)
	OnOutput func([]uint32, uint32) error // OnOutput(Regs, InputData)
}

func GetPortNumber(name string) uint32 {
	return PortsPointer[strings.ToUpper(name)]
}

func RegisterPortWithID(name string, id uint32, port Port) {
	PortsPointer[strings.ToUpper(name)] = id
	PortsUsages[id] = port
}

func init() {
	RegisterPortWithID("TEXT", 1, Port{})
}
