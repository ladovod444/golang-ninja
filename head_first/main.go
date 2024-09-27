package main

import (
	//"bufio"
	"fmt"
	"math/rand"

	//"strings"

	//"greeting"
	//"head/essential/file"

	"head/essential/arithmetic"
	"head/essential/dir"
	"head/essential/sum"

	//"head/essential/robot"

	"log"
	"os"
	/* "strconv"
	"strings" *///"golang.org/x/tools/go/analysis/passes/stringintconv"
	/* "head/essential/exc1"
	"strings" */)

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("Opening refrigerator")
}
func (r Refrigerator) Close() {
	fmt.Println("Closing refrigerator")
}
func (r Refrigerator) FindFood(food string) error {
	r.Open()
	defer r.Close()
	if find(food, r) {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}

	return nil
}

func Soc() {
	fmt.Println("Soc")
}

func adder() func(int) int {
	sum := 0
	return func(arg int) int {
		sum += arg
		return sum
	}
}

func Fact(num int) int {
	if num == 0 {
		return 1
	}

	return num * Fact(num-1)
}

func awardPrize() {
	doorNumber := rand.Intn(3) + 10
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("invalid door number")
	}
}

func calmDown() {
	recover()
}
func freakOut() {
	defer calmDown()
	panic("oh no")
}

func DivideNums(a, b float32) float32 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovery:", r)
		}
	}()

	if b == 0.0 {
		panic("Division by zero")
	}

	return a / b
}

func main() {

	//GetRand()
	resr, errr := arithmetic.GetRand()

	if errr !=nil {
		log.Fatal(errr.Error())
	} else {
		fmt.Println(resr)
	}

	os.Exit(0)

	/* freakOut()
	fmt.Println("Exiting normally") */

	fmt.Println(DivideNums(10, 0))
	fmt.Println(DivideNums(10, 5))

	//awardPrize();

	mdir := "D:\\Epam\\EPAM_PAST_PROJECTS\\JAJ_CAN\\test"
	mdir2 := "D:\\Epam\\EPAM_PAST_PROJECTS\\JAJ_CAN\\test1"

	//defer dir.ReportPanic()
	//panic("some other issue")
	dir.GetDirPan(mdir2, "*")

	// НО дальше программа не выполнятеся
	fmt.Println("Continue programm")
	/*
		 Но если использовать
		 defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovery!!!:", r)
			}
		}()
			в ф-ции GetDirPan
			то выполняется, см. также https://proglib.io/p/samouchitel-po-go-dlya-nachinayushchih-chast-11-obrabotka-oshibok-panika-vosstanovlenie-logirovanie-2024-04-16
	*/

	dir.GetDirPan(mdir, ">")

	os.Exit(0)
	//os.Exit(0);

	//dir.GetDir(mdir, " ")
	err := dir.GetDir(mdir, "*")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(Fact(5))

	os.Exit(0)

	adderr, adder2 := adder(), adder()

	for i := 0; i < 11; i++ {
		fmt.Println(adderr(i), adder2(i*-2))
	}

	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer Soc()
	fmt.Println("test def")
	//os.Exit(0)

	//errf, flts := sum.GetFloats("sum/sums.txt");
	//fmt.Println(os.Args)
	errf, flts := sum.GetFloats(os.Args[1])
	sumf := 0
	if errf != nil {
		log.Fatal(errf.Error())
	} else {
		for _, value := range flts {
			sumf += int(value)
		}
		fmt.Println("sum =", sumf)
	}

}
