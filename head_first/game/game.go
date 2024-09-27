// guess is a game in which the player must guess a random number

package game

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	//"time"
)

func Game() {
	fmt.Println("The guess game")

	var guess_range int = 20

	// Generate a number to guess
	/* seconds := time.Now().Unix()
	rand.Seed(seconds) */
	//var conceived int64 = rand.Int63n(10) + 1
	var conceived int = rand.Intn(guess_range) + 1
	//fmt.Println(conceived)

	//var attempts int = 0
	var possible_attempts int = guess_range
	success := false
	for guesses := 0; guesses < possible_attempts; guesses++ {
		//attempts++ // Increments a number of attempts
		buff := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a number between 1 and ", guess_range, ": ")
		input, err := buff.ReadString('\n')

		if err != nil {
			log.Fatal("Some error")
		}

		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input) //strconv.ParseInt(input, 10, 32)

		if err != nil {
			log.Fatal("Wrong value was enter, not a number")
		}

		if number == conceived {
			fmt.Println("You've guessed !!!")
			success = true
			break
		} else if number < conceived {
			fmt.Println("Oops. Your guess was LOW")
		} else if number > conceived {
			fmt.Println("Oops. Your guess was HIGH")
		}

		/* if guesses > possible_attempts {
			fmt.Println("Sorry But You've reached maximum of the attempts")
			break
		} */
	}
	if !success {
		fmt.Println("Sorry But You've reached maximum of the attempts")
	}

	//fmt.Println(input, err)

	// Получить число от пользователя

}
