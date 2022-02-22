package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SendValue(c chan string) {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	c <- (txtlines[0])

}

func main() {
	values := make(chan string)
	defer close(values)
	go SendValue(values)
	value := <-values
	fmt.Println(value)

}
