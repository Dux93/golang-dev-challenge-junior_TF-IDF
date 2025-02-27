package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var word string
	fmt.Println(print_doc("data/document_1.txt"))
	fmt.Println(print_doc("data/document_2.txt"))
	fmt.Println(print_doc("data/document_3.txt"))

	fmt.Println("From the Documents, type the word that you notice is repeated: ")
	fmt.Scanln(&word)
	a := document("data/document_1.txt", word)
	b := document("data/document_2.txt", word)
	c := document("data/document_3.txt", word)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	d := a + b + c
	fmt.Println("total repeated in all documents: " + fmt.Sprint(d))

}

func print_doc(url string) string {
	var full string
	file, mistake := os.Open(url)
	defer file.Close()
	if mistake != nil {
		fmt.Println("At ocurred a Mistake")
	}
	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		i++
		line := scanner.Text()
		full += line + " \n"
	}
	return full
}

func document(url, word string) int {

	file, mistake := os.Open(url)
	defer file.Close()
	if mistake != nil {
		fmt.Println("At ocurred a Mistake")
	}
	scanner := bufio.NewScanner(file)
	var i, c int

	for scanner.Scan() {
		var fields []string
		i++
		line := scanner.Text()

		fields = strings.Fields(line)

		for i := 0; i < len(fields); i++ {

			if strings.Contains(line, fields[i]) {
				if strings.EqualFold(fields[i], word) {
					c++
				}
			}

		}

	}
	return c
}
