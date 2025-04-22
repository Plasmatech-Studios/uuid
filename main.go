// main.go
package main

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
)

func main() {
	// Generate a new UUID (v4)
	id := uuid.New().String()

	// Copy it to the system clipboard
	if err := clipboard.WriteAll(id); err != nil {
		log.Fatalf("💥 Failed to copy UUID to clipboard: %v", err)
	}

	// Also print it out so you can see it
	fmt.Println(id)
}
