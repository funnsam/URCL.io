package main

// URCL project codename "Red Panda", released under BSD 2-Clause license
/*
BSD 2-Clause License

Copyright (c) 2022, funnsam
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

import "errors"

var NotHLTed = true
var TokenList []Instruction
var CurrentTokenIndex = 0

func InitEmulater() {
	NotHLTed = true
}

func Step() {
	//for NotHLTed {
	//}
}

func OnINHit([]Instruction, int) error {
	return errors.New("urcl compiler: IN is not supported")
}

func OnOUTHit([]Instruction, int) error {
	return errors.New("urcl compiler: OUT is not supported")
}

func init() {
	InstructionList[0] = DefineInstruction{-1, OnInvalidInstruction}
	RegisterURCLInstruction("IN", OnINHit, 2)
	RegisterURCLInstruction("OUT", OnOUTHit, 2)
}
