package functions

import (
	"fmt"
	"strings"
)

func AsciiRepresentation(str string, asciiMap map[rune][]string) {
	words := strings.Split(str, "\\n")
	if words[0] == "" {
		words = words[1:]
	}
	slice := [][]string{}
	for _, word := range words {
		for _, char := range word {
			for key, value := range asciiMap {
				if char == key {
					slice = append(slice, value)
				}
			}
		}
		Print(slice)
		slice = nil
	}
}

func Print(slice [][]string) {
	for i := 0; i <= 7; i++ {
		lineSlice := []string{}
		for _, char := range slice {
			lineSlice = append(lineSlice, char[i])
		}
		strPrint := strings.Join(lineSlice, " ")
		if strPrint == "" {
			fmt.Println()
			break
		}
		fmt.Println(strPrint)
	}
}
