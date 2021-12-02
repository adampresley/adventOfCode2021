package internal

type GuidanceSystem struct {
	aim int
}

func NewGuidanceSystem() *GuidanceSystem {
	return &GuidanceSystem{}
}

func (g *GuidanceSystem) ProcessNavigationInput(input DirectionChangeCollection) (horizontalPos, depth int) {
	for _, directionChange := range input {
		horizontalPos, depth = g.move(horizontalPos, depth, directionChange)
	}

	return
}

func (g *GuidanceSystem) move(currentHorizontalPos, currentDepth int, input DirectionChange) (newHorizontalPos, newDepth int) {
	newHorizontalPos = currentHorizontalPos
	newDepth = currentDepth

	g.tilt(input)

	if input.Instruction == MoveForward {
		newHorizontalPos, newDepth = g.moveForward(currentHorizontalPos, currentDepth, input)
	}

	return
}

func (g *GuidanceSystem) tilt(input DirectionChange) {
	if input.Instruction == TiltUp {
		g.aim -= input.Delta
	}

	if input.Instruction == TiltDown {
		g.aim += input.Delta
	}
}

func (g *GuidanceSystem) moveForward(currentHorizontalPos, currentDepth int, input DirectionChange) (newHorizontalPos, newDepth int) {
	newHorizontalPos = currentHorizontalPos + input.Delta
	newDepth = currentDepth + (input.Delta * g.aim)
	return
}
