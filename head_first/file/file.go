package file

import (
	"fmt"
	//"log"
	"os"
)

func GetFileInfo() {
	fileinfo, err := os.Stat("info.txt")

	if err == nil {
		fmt.Println(fileinfo.Size())
	} else {
		fmt.Println("Error---", err.Error())
	}
}