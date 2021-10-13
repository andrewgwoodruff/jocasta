package main

import (
    "fmt"
    "net/http"
)

func main() {
	fmt.Println("Hello! My name is Jocasta, and I'll be helping you with your data backups today.")

	resp, err := http.Get("http://sample-aws-glacier-archive-url.com")
	if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}
