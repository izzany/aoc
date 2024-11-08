package d115

import "os"

func Part2(input string) int {
	fi, _ := os.ReadFile(input)

	var ans int
    var index int

	for i := 0; i < len(fi); i++ {
        if ans == -1 {
            index = i
            break
        }

		switch char := string(fi[i : i+1]); char {
		case "(":
			ans += 1
		case ")":
			ans -= 1
		default:
			ans += 0
		}
	}
	return index
}
