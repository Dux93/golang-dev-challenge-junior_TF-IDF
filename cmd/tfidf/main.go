package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func print_doc(url string) {
	file, mistake := os.Open(url)
	defer file.Close()
	if mistake != nil {
		fmt.Println("At ocurred a Mistake")
	}
	scanner := bufio.NewScanner(file)
	var i int

	for scanner.Scan() {
		var fields []string
		i++
		line := scanner.Text()

		fields = strings.Fields(line)
		fmt.Println(fields)

	}
}

func wordAllDocument(url string) string{
  var word, words string
   file, mistake := os.Open(url)
    defer file.Close()
	if mistake != nil {
		fmt.Println("At ocurred a Mistake")
	}
	scanner := bufio.NewScanner(file)
	var i, c,k int

	for scanner.Scan() {
		var fields []string
		i++
		line := scanner.Text()

		fields = strings.Fields(line)
		fmt.Println(fields)
		for i := 0; i < len(fields); i++ {
		    if strings.Contains(line, fields[i]) {
				if strings.EqualFold(fields[i], fields[i]) {
				    k++
				    if(k>=1){
				        word = fields[i]
				        c++
				    }
				}
			}
		}
		
words += word + string(c) + "/n"
	}
	return words
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
		fmt.Println(fields)
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


/*
func main() {
	var word string
	print_doc("data/document_1.txt")
	print_doc("data/document_2.txt")
	print_doc("data/document_3.txt")
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
*/


func main() {
	wordAllDocument("data/document_1.txt")
	wordAllDocument("data/document_2.txt")
	wordAllDocument("data/document_3.txt")
}
