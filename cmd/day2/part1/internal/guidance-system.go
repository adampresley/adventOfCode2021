package internal

type GuidanceSystem struct {
}

func NewGuidanceSystem() GuidanceSystem {
	return GuidanceSystem{}
}

func (g GuidanceSystem) ProcessNavigationInput(input DirectionChangeCollection) (horizontalPos, depth int) {
	for _, directionChange := range input {
		horizontalPos, depth = g.move(horizontalPos, depth, directionChange)
	}

	return
}

func (g GuidanceSystem) move(currentHorizontalPos, currentDepth int, input DirectionChange) (int, int) {
	newHorizontalPos := currentHorizontalPos
	newDepth := currentDepth

	if input.Instruction == MoveDown {
		newDepth = g.moveDown(currentDepth, input)
	}

	if input.Instruction == MoveUp {
		newDepth = g.moveUp(currentDepth, input)
	}

	if input.Instruction == MoveForward {
		newHorizontalPos = g.moveForward(currentHorizontalPos, input)
	}

	return newHorizontalPos, newDepth
}

func (g GuidanceSystem) moveUp(currentDepth int, input DirectionChange) int {
	return currentDepth - input.Delta
}

func (g GuidanceSystem) moveDown(currentDepth int, input DirectionChange) int {
	return currentDepth + input.Delta
}

func (g GuidanceSystem) moveForward(currentHorizontalPos int, input DirectionChange) int {
	return currentHorizontalPos + input.Delta
}
