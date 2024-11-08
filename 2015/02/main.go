package main

import (
	"fmt"
	"os"
	"strings"
) 

func main() {
    fi, _ := os.ReadFile("./input.txt")
    fiarr := strings.Split(string(fi), "\\s+")
    fmt.Print(fiarr)    

}

func getVolume(length, width, height int) int {
    
    return 0;
}
