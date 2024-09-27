package dohttp

import (
	"fmt"
	"net/http"
	"time"
)

func Test() {
	fmt.Println("ddd")
}

func Dohttp(url string) {
	t:= time.Now()
	resp, err := http.Get(url)

	if (err != nil) {
		fmt.Printf("Failed to get: <%s> %s\n ", url, err.Error())
	}

	defer resp.Body.Close()

	fmt.Printf("<%s> - Status Code [%d] - Latency %d ms\n", url, resp.StatusCode, time.Since(t).Milliseconds());
}
