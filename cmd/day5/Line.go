package main

import (
	"fmt"
	"image"
	"strings"
)

type LineDirection string

const (
	Horizontal LineDirection = "horizontal"
	Vertical   LineDirection = "vertical"
	Diagonal   LineDirection = "diagonal"
)

type Line struct {
	Point1    image.Point
	Point2    image.Point
	direction LineDirection
}

type LineCollection []*Line

func (l *Line) IsStraightLine() bool {
	return l.Point1.X == l.Point2.X || l.Point1.Y == l.Point2.Y
}

func NewLine(x1, y1, x2, y2 int) *Line {
	direction := Diagonal

	if x1 == x2 {
		direction = Vertical
	} else if y1 == y2 {
		direction = Horizontal
	}

	result := &Line{
		Point1:    image.Pt(x1, y1),
		Point2:    image.Pt(x2, y2),
		direction: direction,
	}

	return result
}

func (lc LineCollection) String() string {
	result := strings.Builder{}

	for _, l := range lc {
		result.WriteString(l.String())
		result.WriteString("\n")
	}

	return result.String()
}

func (l *Line) ToInts() (int, int, int, int) {
	return l.Point1.X, l.Point1.Y, l.Point2.X, l.Point2.Y
}

func (l *Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d (%s)", l.Point1.X, l.Point1.Y, l.Point2.X, l.Point2.Y, l.direction)
}
