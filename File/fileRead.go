package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Simple File Reading

	// data, err := ioutil.ReadFile("file1.txt")

	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// }
	// str := string(data)
	// fmt.Print(str)

	//File Reading with Flag

	fileptr := flag.String("filePath", "file1.txt", "file path to read from")
	flag.Parse()

	data1, err := ioutil.ReadFile(*fileptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("File Read with Flag : ", string(data1))

	//Reading file Line by Line
	f, err := os.Open(*fileptr)

	s := bufio.NewScanner(f)
	for i := 1; s.Scan(); i++ {
		fmt.Println(i, ": ", s.Text())
	}

	//File reading by sentence

	//error
	// fmt.Println("Full Stop break")
	// for s.Scan() {
	// 	fmt.Println(s.Text())
	// }

	//printing characters
	fmt.Println("Printing characters")
	s1 := bufio.NewScanner(f)
	for i := 1; s1.Scan(); i++ {
		str1 := string(s1.Text())
		fmt.Println(i, " line ", str1)
	}
}
