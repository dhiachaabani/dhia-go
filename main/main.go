package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	val := os.Getenv("es")
	fmt.Println(val)

	fmt.Println(Hello("chris"))
}
