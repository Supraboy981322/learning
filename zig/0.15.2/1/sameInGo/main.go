package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	fi, err := os.Open("foo.txt")
	if err != nil { eror(err) }
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		li := scanner.Text()
		fmt.Println(li)
	}
	
	err = scanner.Err()
	if err != nil { eror(err) }
}

func eror(err error) {
	os.Stderr.Write([]byte(err.Error()))
	os.Exit(1)
}

