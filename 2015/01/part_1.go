package d115

import (
	"os"
)

func Part1(input string) int {
	fi, _ := os.ReadFile(input)

	var ans int

	for i := 0; i < len(fi); i++ {
		switch char := string(fi[i : i+1]); char {
		case "(":
			ans += 1
		case ")":
			ans -= 1
		default:
			ans += 0
		}
	}
	return ans
}
