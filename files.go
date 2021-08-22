package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Create multiple directories inside a dir
func createDirs(dirNum int64) {

	for i := 0; int64(i) < dirNum; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your directory's name: ")
		dirName, _ := reader.ReadString('\n')
		os.Mkdir(dirName, 0755)
	}
}

// The menu for file system tasks
func filesMain() {

	fileOptions := map[int]string{
		0: "Main Menu",
		1: "Create multiple directories",
	}

	// Display options
	for option, description := range fileOptions {
		fmt.Printf("%d : %s \n", option, description)
	}

	// Create a new reader
	scanner := bufio.NewScanner(os.Stdin)
	// Get a choice from the user via stdin
	fmt.Print("Enter an int from the options above: ")

	userChoice, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if userChoice == 0 {
		main()
	} else if userChoice == 1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a location to create dirs: ")
		userLocation, _ := reader.ReadString('\n')
		os.Chdir(userLocation)
		fmt.Println("Enter how many directories you want to create: ")
		userDirNum, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		createDirs(userDirNum)
		main()
	} else {
		fmt.Print("*** Invalid option. Returning to main menu. ***")
		main()
	}

}
