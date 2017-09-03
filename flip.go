package main

func getFlipCount(stack []bool) int {
	if len(stack) == 1 {
		if stack[0] {
			return 0
		}

		return 1
	}

	stack[0] = !stack[0]
	return 1 + getFlipCount(squash(stack))
}
