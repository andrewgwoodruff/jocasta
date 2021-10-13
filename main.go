package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello! My name is Jocasta, and I'll be helping you with your data backups today.")

	response, error := http.Get("http://example.com")
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)

	scanner := bufio.NewScanner(response.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if error := scanner.Err(); error != nil {
		panic(error)
	}
}
