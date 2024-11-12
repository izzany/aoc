package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var volumetotal int
	var ribbontotal int
	var count int

	for fiarr.Scan() {
		line := fiarr.Text()

		x := func(c rune) bool {
			return c == 'x'
		}

		fields := strings.FieldsFunc(line, x)
		if len(fields) > 0 {
			_length, _ := strconv.Atoi(fields[0])
			_width, _ := strconv.Atoi(fields[1])
			_height, _ := strconv.Atoi(fields[2])
			dimension := Dimension{_length, _width, _height}
			volumetotal += getVolume(dimension)
			ribbontotal += getRibbon(dimension)
			count += 1
		}
	}
	fmt.Printf("volume : %v\n", volumetotal)
	fmt.Printf("ribbon : %v\n", ribbontotal)

}

func getVolume(dimension Dimension) int {

	sidelw := dimension.length * dimension.width
	sidewh := dimension.width * dimension.height
	sidehl := dimension.height * dimension.length

	slack := func(sidelw, sidewh, sidehl int) int {

		if (sidelw <= sidewh) && (sidelw <= sidehl) {
			return sidelw
		} else if (sidewh <= sidelw) && (sidewh <= sidehl) {
			return sidewh
		} else if (sidehl <= sidelw) && (sidehl <= sidewh) {
			return sidehl
		} else {
			return 0
		}
	}

	slackvalue := slack(sidelw, sidewh, sidehl)
	volume := (2 * sidelw) + (2 * sidewh) + (2 * sidehl) + slackvalue

	return volume
}

func getRibbon(dimension Dimension) int {
	sides := []int{dimension.length, dimension.width, dimension.height}
	sort.Slice(sides, func(i, j int) bool {
		return sides[i] < sides[j]
	})
	wrap := (2 * sides[0]) + (2 * sides[1])
	bow := dimension.length * dimension.width * dimension.height
	ribbon := wrap + bow

	return ribbon
}
