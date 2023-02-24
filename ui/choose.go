package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChooseOption(options []string) uint {
	for index, option := range options {
		fmt.Printf("%d: %s\n", index+1, option)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r", "", -1) // For windows
		text = strings.Replace(text, "\n", "", -1)

		chosenIndex, err := strconv.Atoi(text)

		if err == nil && chosenIndex >= 1 && chosenIndex <= len(options) {
			return uint(chosenIndex) - 1
		}

		fmt.Println("INVALID INPUT!")
	}
}
