package sum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	file.Close()

	fmt.Println("Closing file")
}

func GetFloats(fileName string) (error, []float64) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return err, nil
	}
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return err, nil
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return scanner.Err(), nil
	}
	return nil, numbers
}
