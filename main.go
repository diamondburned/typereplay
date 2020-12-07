package main

import (
	"flag"
	"log"
	"os"
	"time"
)

var (
	input = "input"
	wait  = 5 * time.Second
)

func init() {
	flag.StringVar(&input, "i", input, "file to read commands from")
	flag.DurationVar(&wait, "w", wait, "time to wait before executing")
}

func main() {
	flag.Parse()

	f, err := os.Open(input)
	if err != nil {
		log.Fatalln("failed to read input:", err)
	}
	defer f.Close()

	t, err := ParseInput(f)
	if err != nil {
		log.Fatalln("failed to parse input:", err)
	}

	f.Close()

	log.Println("Waiting for", wait)
	time.Sleep(wait)
	log.Println("Executing tape")

	t.Do()
}
