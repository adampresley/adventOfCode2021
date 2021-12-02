package internal

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Instruction string

const (
	MoveUp      Instruction = "up"
	MoveDown    Instruction = "down"
	MoveForward Instruction = "forward"
)

type DirectionChange struct {
	Instruction Instruction
	Delta       int
}

type DirectionChangeCollection []DirectionChange

func NewDirectionChangeCollection() DirectionChangeCollection {
	return make(DirectionChangeCollection, 0, 200)
}

func (i Instruction) String() string {
	return string(i)
}

func Parse(reader io.Reader) DirectionChangeCollection {
	result := NewDirectionChangeCollection()
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		result = append(result, parseLine(scanner.Text()))
	}

	return result
}

func parseLine(line string) DirectionChange {
	split := strings.Split(line, " ")

	delta, _ := strconv.Atoi(split[1])

	return DirectionChange{
		Instruction: Instruction(split[0]),
		Delta:       delta,
	}
}
