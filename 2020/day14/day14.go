package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/a-buck/adventofcode/2020/utils"
)

var (
	inputFilePath = flag.String("input", "day14.txt", "input file path")
	lineExpr      = regexp.MustCompile(`^mem\[(?P<addr>\d+)\]\s=\s(?P<val>\d+)$`)
	partB         = flag.Bool("partB", false, "enable part b")
)

type instruction struct {
	addr, val uint64
}

type program struct {
	mask   string
	instrs []instruction
}

func main() {

	flag.Parse()

	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(content), "mask = ")[1:]

	programs := make([]program, 0)

	for _, part := range parts {
		lines := strings.Split(part, "\n")
		mask := lines[0]

		prog := program{mask: mask}

		for _, line := range lines[1:] {
			if line == "" {
				continue
			}
			result := utils.Match(line, lineExpr)

			addr := utils.ToUint(result["addr"])
			val := utils.ToUint(result["val"])

			prog.instrs = append(prog.instrs, instruction{addr: addr, val: val})
		}
		programs = append(programs, prog)
	}

	mem := make(map[uint64]uint64)

	if *partB {
		for _, prog := range programs {
			options := maskOptions(prog.mask)

			maskAStr := strings.ReplaceAll(prog.mask, "X", "0")
			maskA, _ := strconv.ParseUint(maskAStr, 2, 0)

			maskBStr := make([]rune, len(prog.mask))
			for i, v := range prog.mask {
				if v == 'X' {
					maskBStr[i] = '0'
				} else {
					maskBStr[i] = '1'
				}
			}
			maskB, _ := strconv.ParseUint(string(maskBStr), 2, 0)

			for _, instr := range prog.instrs {
				m := instr.addr & maskB // set X positions to zero
				m = m | maskA           // overwrite with 1's from mask

				for _, optStr := range options {
					opt, _ := strconv.ParseUint(optStr, 2, 0)
					addr := m | opt
					mem[addr] = instr.val
				}
			}
		}

	} else {
		// part A
		for _, prog := range programs {
			mask1, _ := strconv.ParseUint(strings.ReplaceAll(prog.mask, "X", "0"), 2, 0)
			mask2, _ := strconv.ParseUint(strings.ReplaceAll(prog.mask, "X", "1"), 2, 0)
			for _, instr := range prog.instrs {
				a := mask1 | instr.val
				b := mask2 & a
				mem[instr.addr] = b
			}
		}
	}

	var sum uint64
	for _, v := range mem {
		sum += v
	}
	fmt.Println(sum)
}

func maskOptions(mask string) []string {
	options := make([]string, 0)

	xIndexes := make([]int, 0)
	for i, v := range mask {
		if v == 'X' {
			xIndexes = append(xIndexes, i)
		}
	}

	n := (1 << len(xIndexes))
	for i := 0; i < n; i++ {
		m := make([]rune, len(mask))
		for j := range m {
			m[j] = '0'
		}

		for j, xIdx := range xIndexes {
			interval := 1 << (len(xIndexes) - j - 1)
			val := (i / interval) % 2
			r := strconv.Itoa(val)[0]
			m[xIdx] = rune(r)
		}

		options = append(options, string(m))
	}

	return options
}
