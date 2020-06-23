package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(instruction string, parInput *string) {
	fmt.Print(instruction)
	reader := bufio.NewReader(os.Stdin)
	*parInput, _ = reader.ReadString('\n')
	*parInput = strings.Replace(*parInput, "\r", "", -1)
	*parInput = strings.Replace(*parInput, "\n", "", -1)
}
