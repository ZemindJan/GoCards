package ui

import (
	"bufio"
	"fmt"
	"os"
)

func PressEnterToContinue() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to continue...")
	reader.ReadString('\n')
}
