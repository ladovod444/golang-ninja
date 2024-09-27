package arithmetic

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

// func GetRand(number, upTo int) map[int]int {
func GetRand() ([]int, error) {

	var res_map map[int]int

	/* for i := 0; i <= upTo; i++ {
		num := rand.IntN(upTo)
	} */

	fmt.Printf("Enter number of values you want to generate: ")
	buff := bufio.NewReader(os.Stdin)

	user_entered, err := buff.ReadString('\n')

	if err != nil {
		log.Fatal("Wrong value entered")
	}

	input := strings.TrimSpace(user_entered)
	number_res, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Wrong number entered")
	} else {

		fmt.Printf("Enter max value used to generate: ")
		buff := bufio.NewReader(os.Stdin)

		user_entered, err := buff.ReadString('\n')

		if err != nil {
			log.Fatal("Wrong value entered")
		}

		input := strings.TrimSpace(user_entered)
		max_value, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal("Wrong value entered")
		} else {

			if (number_res >= max_value) {
				fmt.Printf("Number of values cannot be equal or greater then max value\n")
				//os.Clearenv()
				//GetRand()	
				os.Exit(0);					
			}

			res_map = make(map[int]int, number_res)

			//fmt.Printf("MAX = %d number = %d\n", max_value, number_res)

			for {
				num := rand.IntN(max_value)
				if num != 0 {
					res_map[num] = num
				}

				if len(res_map) == number_res {
					break
				}
			}
		}

	}
    //buff.Reset(os.Stdin)
	return maps.Keys(res_map), err
}
