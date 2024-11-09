package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dimension struct {
	length int
	width  int
	height int
}

func main() {
	fi, _ := os.Open("./input.txt")
	fiarr := bufio.NewScanner(fi)
	fmt.Print(fiarr)

	var volumetotal int

	for fiarr.Scan() {
		line := fiarr.Text()

		x := func(c rune) bool {
			return c == 'x'
		}

		fields := strings.FieldsFunc(line, x)
		dimension := Dimension{strconv.Atoi(fields[0]), strconv.Atoi(fields[1]), strconv.Atoi(fields[2])}
		d := &dimension

		volumetotal += getVolume(d)

	}

}

func getDimension() struct{}

func getVolume(class Dimension) int {

	return 0
}
