package readfile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ReadFileText struct {
	name string
	size int
}

func (r ReadFileText) GetSize() int {
	return r.size
}

func (r ReadFileText) GetName() string {
	return r.name
}



func ReadFile(filename string) ([]float64, error) {
	rf := ReadFileText{"data.txt", 30}
	fmt.Println(rf)

	var numbers []float64

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 32)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, num)
		//fmt.Println(scanner.Text())

	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, err
}
