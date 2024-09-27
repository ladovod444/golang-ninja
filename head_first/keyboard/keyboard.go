package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DaysInWeek = 7

func GetFloat() (float64, error) {
	fmt.Print("Enter float value: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
