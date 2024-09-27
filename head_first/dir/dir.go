package dir

import (
	"fmt"
	//"io/ioutil"
	//"log"
	"os"
)

func ReportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println("Recovering...", err)
	} else {
		panic(p)
	}

}

func GetDir(dir string, prefix string) error {
	//files, err := ioutil.ReadDir(dir); //deprecated
	next_pref := prefix
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
		//log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {

			fmt.Println(next_pref, "Directory:", dir)
			dir = dir + "\\" + file.Name()
			next_pref += prefix
			err = GetDir(dir, next_pref)
		} else {
			fmt.Println(next_pref, "File:", dir, "\\", file.Name())
		}
	}

	return nil
}

func GetDirPan(dir string, prefix string) {
	//files, err := ioutil.ReadDir(dir); //deprecated
	files, err := os.ReadDir(dir)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovery!!!:", r)
		}
	}()

	if err != nil {
		panic(err)
		//fmt.Println(err.Error())
		//log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {

			fmt.Println(prefix, "Directory:", dir)
			dir = dir + "\\" + file.Name()
			prefix += prefix
			GetDirPan(dir, prefix)
		} else {
			fmt.Println(prefix, "File:", dir, "\\", file.Name())
		}
	}

	//return nil
}
