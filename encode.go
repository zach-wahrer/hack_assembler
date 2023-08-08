package hackassembler

import (
	"fmt"
	"strconv"
)

func Translate(ins Instruction) (string, error) {
	switch ins.Type {
	case AInstruction:
		num, err := strconv.ParseInt(ins.Symbol, 10, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%016s", strconv.FormatInt(num, 2)), nil
	case CInstruction:
		return fmt.Sprintf("111%s%s%s", TranslateComp(ins.Comp), TranslateDest(ins.Dest), TranslateJump(ins.Jump)), nil
	case LInstruction:
		// TODO: Implement this, later.
		return "", nil
	default:
		return "", fmt.Errorf("unknown instruction type: %s", ins.Type)
	}
}

var allDests = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"DM":  "011",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"MA":  "101",
	"AD":  "110",
	"DA":  "110",
	"ADM": "111",
}

func TranslateDest(dest string) string {
	return allDests[dest]
}

var allComps = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",
	"M":   "1110000",
	"!M":  "1110001",
	"-M":  "1110011",
	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

func TranslateComp(comp string) string {
	return allComps[comp]
}

var allJumps = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

func TranslateJump(jump string) string {
	return allJumps[jump]
}
