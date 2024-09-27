package main

import (
	"fmt"
	//"io/ioutil"
	"io"
	"log"
	"net/http"
	//"net/url"
	"time"
)

/* func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}
func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}
func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end main()")
} */

func repeat(s string) {
	for i := 0; i < 25; i++ {
		fmt.Print(s)
	}
}

func greet(myChan chan string, mess string) {
	myChan <- mess
}

type Page struct {
	url string
	size int
}

//func getContentByUlr(url string, rchan chan int) {
func getContentByUlr(url string, rchan chan Page) {
	resp, err := http.Get(url)
	//fmt.Println(url, resp.Status)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	//rchan <- len(result)
	rchan <- Page{ url, len(result)}
	//fmt.Println(len(result))
	//fmt.Println(string(result))

}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}
func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}



func main() {

	/* myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel) */
  
	urls := [] string {
		"https://golang.org/",
		"https://example.com/",
		"https://rambler.ru",
	}

	//mchan := make(chan int)
	//mchan := make(chan Page)
	pages := make(chan Page)

	for _, url := range urls {
		go getContentByUlr(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page:= <-pages
		//fmt.Println(<-pages)
		fmt.Printf("%s: %d\n", page.url, page.size)
	}
	
/* 	go getContentByUlr("https://golang.org/", mchan)
	go getContentByUlr("https://example.com/", mchan)
	go getContentByUlr("https://rambler.ru", mchan) */
	//fromChan := <-mchan
	/* fmt.Println(<-mchan)
	fmt.Println(<-mchan)
	fmt.Println(<-mchan) */

	/* getContentByUlr("https://golang.org/");
	getContentByUlr("https://golang.org/doc"); */

	/* myChannel := make(chan string)

	go greet(myChannel, "mess to channel")

	fmt.Println(<- myChannel) */

	/* var myChannel chan float64
	myChannel = make(chan float64) */
	/*
	   	var size int
	   size = go responseSize("https://example.com/")
	   fmt.Println(size)
	   size = go responseSize("https://golang.org/")
	   fmt.Println(size)
	   size = go responseSize("https://golang.org/doc")
	   fmt.Println(size)
	   time.Sleep(5 * time.Second)
	*/
	/* go repeat("y")
	go repeat("x")
	time.Sleep(time.Second) */
	/* go responseSize("https://example.com/")
	go responseSize("https://golang.org/")
	go responseSize("https://golang.org/doc")
	time.Sleep(time.Second * 5)
	fmt.Println("end main()") */
}

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}
