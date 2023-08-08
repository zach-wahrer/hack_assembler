package hackassembler

import (
	"bufio"
	"io"
	"strings"
)

type (
	Parser struct {
		Contents io.Reader
	}

	InstructionType string

	Instruction struct {
		Type   InstructionType
		Symbol string
		Dest   string
		Comp   string
		Jump   string
	}
)

const (
	AInstruction InstructionType = "A"
	CInstruction InstructionType = "C"
	LInstruction InstructionType = "L"
)

func NewParser(contents io.Reader) *Parser {
	return &Parser{Contents: contents}
}

func (p *Parser) Parse() <-chan Instruction {
	instructions := make(chan Instruction)
	scanner := bufio.NewScanner(p.Contents)
	scanner.Split(bufio.ScanLines)

	go func() {
		defer close(instructions)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if p.SkipLine(line) {
				continue
			}

			switch p.GetInstructionType(line) {
			case AInstruction:
				instructions <- Instruction{
					Type:   AInstruction,
					Symbol: p.GetSymbol(line),
				}
			case CInstruction:
				instructions <- Instruction{
					Type: CInstruction,
					Dest: p.GetDest(line),
					Comp: p.GetComp(line),
					Jump: p.GetJump(line),
				}
			case LInstruction:
				// TODO: Handle this instruction, later.
			}

		}
	}()

	return instructions
}

func (p *Parser) SkipLine(line string) bool {
	return strings.HasPrefix(line, "//") || line == ""
}

func (p *Parser) GetInstructionType(line string) InstructionType {
	// TODO: Error handling, later
	if strings.HasPrefix(line, "@") {
		return AInstruction
	} else if strings.HasPrefix(line, "(") {
		return LInstruction
	} else {
		return CInstruction
	}
}

func (p *Parser) GetSymbol(line string) string {
	return strings.TrimLeft(line, "@")
}

func (p *Parser) GetDest(line string) string {
	splitDest := strings.Split(line, "=")
	if len(splitDest) == 1 {
		return ""
	}
	return splitDest[0]
}

func (p *Parser) GetComp(line string) string {
	splitDest := strings.Split(line, "=")

	var noDest string
	if len(splitDest) > 1 {
		noDest = splitDest[1]
	} else {
		noDest = splitDest[0]
	}

	return strings.Split(noDest, ";")[0]
}

func (p *Parser) GetJump(line string) string {
	splitJump := strings.Split(line, ";")
	if len(splitJump) == 1 {
		return ""
	}
	return splitJump[1]
}
