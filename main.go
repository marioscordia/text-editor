package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println(errors.New("too much or not enough arguments"))
		return
	}

	if strings.Contains(args[0], ".txt") != true || strings.Contains(args[1], ".txt") != true {
		fmt.Println(errors.New("not correct format of file(s)"))
		return
	}

	sample, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sampleArr := strings.Split(string(sample), "\n")

	result, err := os.Create(args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer result.Close()

	writeToFile(sampleArr, result)
}

func writeToFile(sentences []string, file *os.File) {
	for _, v := range sentences {
		arr := strings.Fields(v)
		arr = BigFunc(arr)
		arr = article(arr)
		arr = puncts1(arr)
		arr = puncts2(arr)
		for i := range arr {
			if i != len(arr)-1 {
				file.WriteString(arr[i] + " ")
			} else {
				file.WriteString(arr[i] + "\n")
			}
		}

	}
}
