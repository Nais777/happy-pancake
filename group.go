package main

func flipSection(section []bool) {
	for i := 0; i < len(section); i++ {
		section[i] = !section[i]
	}
}

func squash(stack []bool) []bool {
	ret := make([]bool, 0, len(stack))
	prior := stack[0]

	for _, up := range stack {
		if up == prior {
			continue
		}

		ret = append(ret, prior)
		prior = up
	}

	return append(ret, prior)
}
