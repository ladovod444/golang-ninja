package readfile2

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetArgs() {
	//fmt.Println(os.Args)
	if len(os.Args) > 1 {
		for value := range os.Args {
			fmt.Println(value)
		}

	}	
}

func GetArgsAverage() float64 {
	//average := 0.0
	//fmt.Println(os.Args)

	clean_arg := os.Args[1:]

	//fmt.Println(clean_arg)

	sum := 0.0
	if len(clean_arg) > 1 {
		for _, value := range clean_arg {
			//fmt.Println(value)
			number, err := strconv.ParseFloat(value, 64)
			if (err != nil) {
				log.Fatal("Not float data was entered");
			}
			sum += number
		}

	}

	return sum / float64(len(os.Args))
}