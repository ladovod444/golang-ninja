package elects

import (
	"bufio"
	//"fmt"

	//"log"
	"os"
	//"strings"
)

func ReadElect(filname string) (map[string]int, error) {
	//
	/* var electsMap map[string]int
	electsMap = make(map[string]int) */
	electsMap := map[string]int{}
	file, err := os.Open(filname)

	if err != nil {
		//log.Fatal(err.Error())
		return electsMap, err
	} else {

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			//fmt.Println(scanner.Text())
			line := scanner.Text()

			electsMap[line]++

			/* _, ok := electsMap[line]

			if ok {
				electsMap[line]++
			} else {
				electsMap[line] = 1
			} */

		}
		err = file.Close()

	}

	return electsMap, err
}
