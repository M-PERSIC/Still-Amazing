package main

import (
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("amazing.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	entry := time.Now().Format(time.DateTime)
	if _, err := f.Write([]byte(entry + "\n")); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
