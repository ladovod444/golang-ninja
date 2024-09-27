package datafile

import (
	"bufio"
	//"fmt"
	"os"
	"log"
)

type Elector struct {
	name string
	vote int
}

/* func in_array(val string, array []string) (ok bool, i int) {
    for i = range array {
        if ok = array[i] == val; ok {
            return
        }
    }
    return
} */

// GetStrings читает строку из каждой строки данных файла.
func GetStrings(fileName string) ([]string, []Elector, error) {
	var lines []string
	var linesElector []Elector
	//var linesElectorFinal []Elector

	//var counts [] int

	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(file)
	votes := 0

	//count := 0
	for scanner.Scan() {
		line := scanner.Text()

		//matched := false

		elect := Elector{
			name: line,
			vote: votes,
		}

		linesElector = append(linesElector, elect)
		/* for i, value := range linesElector {
			if value.name == line {
				elect.vote++
				matched = true
				count++

				counts[i]++

				//linesElector[count] = elect
				//linesElectorFinal = append(linesElectorFinal, elect)
			}
		}

		if matched {
			//elect.vote++
			fmt.Println("matched")

			//linesElector[count] = elect
			//if in_array()
			counts = append(counts, 1)
			linesElectorFinal = append(linesElectorFinal, elect)

		}
 */
		lines = append(lines, line)

	}
	err = file.Close()
	if err != nil {
		return nil, nil, err
	}
	if scanner.Err() != nil {
		return nil, nil, scanner.Err()
	}
	return lines, linesElector, nil
	//return lines, linesElectorFinal, nil
}

func ReadFileMap(filename string) (map[string]int, error) {
	votes := map[string]int{}

	file1, err1 := os.Open(filename)
	if err1 != nil {
		log.Fatal("error file not found")
	} else {
		rf1 := bufio.NewScanner(file1)

		votes_count := 1
		for rf1.Scan() {

			//fmt.Println(rf1.Text());
			//votes_count++

			elect := rf1.Text()   // Получаем строку
			_, ok := votes[elect] // Проверяем наличие в массиве
			if ok {               // Если же есть такой кандидат,
				//fmt.Println("user: ", elect)
				//fmt.Println("user elect_count: ", votes[elect])
				votes[elect]++ // то увеличиваем значение голосов на 1
			} else {
				// Иначе задаем 1
				votes[elect] = votes_count
			}
		}
		err1 = file1.Close()
	}

	return votes, err1

	/* if err1 == nil {
		fmt.Println("Результаты: ", votes)
	} */
}
