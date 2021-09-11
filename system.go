package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func osInfo() {

	fmt.Println("OS: " + runtime.GOOS)
	fmt.Println("Architecture: " + runtime.GOARCH)
	fmt.Println("Go Version: " + runtime.Version())

}

func systemMain() {

	sysOptions := map[int]string{
		0: "Main Menu",
		1: "Get OS Info",
	}

	// Display options
	for option, description := range sysOptions {
		fmt.Printf("%d : %s \n", option, description)
	}

	// Create a new reader and scanner
	scanner := bufio.NewScanner(os.Stdin)
	// Get a choice from the user via stdin
	fmt.Print("Enter an int from the options above: ")
	scanner.Scan()

	userChoice, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if userChoice == 0 {
		main()
	} else if userChoice == 1 {
		fmt.Println()
		osInfo()
		fmt.Println()
		main()
	} else {
		fmt.Print("*** Invalid option. Returning to main menu. ***")
	}

}
