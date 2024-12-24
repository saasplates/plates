package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter project name: ")
	projectName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	projectName = strings.TrimSpace(projectName)

	err = os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created project directory: %s\n", projectName)
} 