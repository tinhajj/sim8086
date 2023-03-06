package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("missing file")
	}

	bytes, err := os.ReadFile(args[0])

	if err != nil {
		log.Fatalf("unable to read file: %s", err)
	}

	fmt.Printf("bits 16\n\n")
	decodeFromBytes(bytes)
}

func decodeFromBytes(b []byte) {
	for i := 0; i < len(b); i += 2 {
		b1 := b[i]
		b2 := b[i+1]

		instruction := b1 >> 2
		dw := b1 &^ 0b11111100
		d := b1 &^ 0b11111101 >> 1
		w := b1 &^ 0b11111110

		mod := b2 >> 6
		reg := b2 &^ 0b11000111 >> 3
		rm := b2 &^ 0b11111000

		// fmt.Printf("full:\t%08b %08b\n", b1, b2)
		// fmt.Printf("inst:\t%08b\n", instruction)
		// fmt.Printf("dw:\t%08b\n", dw)
		// fmt.Printf("d:\t%08b\n", d)
		// fmt.Printf("w:\t%08b\n", w)
		// fmt.Printf("mod:\t%08b\n", mod)
		// fmt.Printf("reg:\t%08b\n", reg)
		// fmt.Printf("rm:\t%08b\n", rm)

		decode(instruction, dw, d, w, mod, reg, rm)
	}
}

func decode(instruction, dw, d, w, mod, reg, rm byte) {
	result := ""
	if instruction == 0b100010 {
		result += "mov "
	}

	// TODO: remove duplication
	if d == 0b1 {
		if w == 0b1 {
			result += lutW[reg] + ", "
			result += lutW[rm]
		} else if w == 0b0 {
			result += lutB[reg] + ", "
			result += lutB[rm]
		}
	} else if d == 0b0 {
		if w == 0b1 {
			result += lutW[rm] + ", "
			result += lutW[reg]
		} else if w == 0b0 {
			result += lutB[rm] + ", "
			result += lutB[reg]
		}
	}

	fmt.Println(result)
}

// 100010 11
