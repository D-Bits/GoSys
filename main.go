package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var mainOptions = map[int]string{
	0: "Exit Program",
	1: "File System Tasks",
	2: "Misc Tasks",
}

func main() {

	for option, description := range mainOptions {
		fmt.Printf("%d : %s \n", option, description)
	}

	// Create a new scanner
	scanner := bufio.NewScanner(os.Stdin)
	// Get a choice from the user via stdin
	fmt.Print("Enter an int from the options above: ")
	scanner.Scan()

	userChoice, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if userChoice == 0 {
		// End program
	} else if userChoice == 1 {
		filesMain()
	} else if userChoice == 2 {
		miscMain()
	} else {
		fmt.Println("Invalid option.")
		main()
	}
}
