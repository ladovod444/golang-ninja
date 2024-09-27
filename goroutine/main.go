package main

import (
	"fmt"
	"os"
	"sort"

	//"ghttp/ghttp/dohttp"
	"math/rand/v2"
	//"sync"
	"time"
)

// Показатель степени, в к-рую надо возвести число, называемое основаниемv base, чтобы получить данное число res.
func Log(base, res int) int {

	//fmt.Println(base, res)
	count := 1
	b := base
	for i := 1; i < res; i++ {
		b = base * b
		count++
		//fmt.Println(b)
		if b >= res {
			return count
		}

	}
	return count
}

func bin_search(arr []int, elem int) int {
	// sort.Ints(arr)
	/* for pos, val := range arr {
		if val == elem {
			return pos
		}
	} */

	low := 0
	high := len(arr) - 1 // 7    3

	check_count := 0
	for low <= high {
		mid := (low + high) / 2 // 10, // 5

		guess := arr[mid]  // 12, 6 
		if guess == elem { // 12 == 3, 6 == 3
			check_count++
			return mid
		}
		if guess > elem { // 12 > 3, 6 > 3
			high = mid - 1 // high = 12 - 1 т.е. 11, 5 - 1 т.е. 4 
			check_count++
		} else {
			low = mid + 1
			check_count++
		}

		fmt.Println("check_count=", check_count)
	}

	return 0
}

func main() {

	/* arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Arr2=", arr2) */


	arr3 := make([] int, 128);
	for i := 1 ; i <= 128; i++ {
		arr3[i-1] = i;
	}
    fmt.Println("Arr3=", arr3) 
	s := bin_search(arr3, 3)
	fmt.Println("Pos=", s)

	log := Log(2, 16)
	fmt.Println("log res=", log)

	log2 := Log(3, 27)
	fmt.Println("log res=", log2)

	arr := []int{1, 6, 2, 5, 32, 11, 4}

	sort.Ints(arr)
	fmt.Println(arr)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(arr)

	os.Exit(0)

	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			message1 <- "Прошло полсекунды"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			message2 <- "Прошло 2 секунды"
		}
	}()

	for {

		// Не блокирующее чтение
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		}

		/* fmt.Println(<-message1)
		fmt.Println(<-message2) */
	}

	/*

		urls := []string{
			"https://example.com/",
			"https://www.youtube.com/",
			"https://github.com/",
			"https://www.instagram.com/",
			"https://www.rambler.ru",
		}

		var wg sync.WaitGroup

		for _, url := range urls {

			wg.Add(1)

			go func(url string) {
				dohttp.Dohttp(url)
				wg.Done()
			}(url)
		}

		wg.Wait()
	*/

	//dohttp.Test();

	/* message := make(chan string, 2) // буфферизированный канал и НЕ блокирует

	message <- "some mess to buff channel"
	message <- "some mess to buff channel2"

	fmt.Println(<-message)
	message <- "some mess to buff channel3"

	//msg:= <-message
	fmt.Println(<-message)
	fmt.Println(<-message) */

	/* message := make(chan string) // В данном случае это НЕ буфферизированный канал

	go func() {
		count := 10
		for i := 0; i <= count; i++ {
			// Записываем в канал числа с интервалом 0.5 с
			time.Sleep(time.Millisecond * 500)

			message <- fmt.Sprintf("%d", i)

		}

		// Закрытие канала
		close(message)
		//time.Sleep(time.Second * 2)

		// Запись в канал
		//message <- "hello to channel"
	}() */

	/* for  {
		// Читаем сообщение из канала
		msg, open := <-message
		if !open {
			break
		}
		fmt.Println(msg)
	} */

	/* for msg:= range message {
		fmt.Println(msg)
	} */

	/* message:= make(chan string)
	message <- "Some message from main" */
	// !!! DEADLOCK

	/* go func ()  {
		time.Sleep(time.Second * 2)

		// Запись в канал
		message <- "hello to channel"
	}() */

	// Читаем сообщение из канала
	/* msg := <- message
	fmt.Println(msg) */

	//fmt.Println(<-message)

	// и выводим в консоль

	/* go fmt.Println("Hello from goroutine")
	go fmt.Println("Hello from main()") */

	//time.Sleep(time.Millisecond)

	// Ф-ция main это главная горутина,
	// когда запускаем time.Sleep, то это блокирует выполнение этой главной горутины
	// и планировщик переключается на другую, ранее запланир. горутину

	/* t:= time.Now();

	go parseUrl("https://example.com/")
	parseUrl("https://www.youtube.com/")


	fmt.Printf("Parsing complete. The elapsed time: %f seconds", time.Since(t).Seconds()); */
}

func parseUrl(url string) {

	for i := 0; i < 5; i++ {
		latency := time.Duration(rand.IntN(500) + 500)
		time.Sleep(time.Duration(latency) * time.Millisecond)

		fmt.Printf("Parsing <%s> - Step %d, Latency %d\n", url, i+1, latency)
	}

}
