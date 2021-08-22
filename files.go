package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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

// Calculate a SHA256 checksum for a file
func calcHash(filePath string) {

	hasher := sha256.New()
	fileName, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileName.Close()

	if _, err := io.Copy(hasher, fileName); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Your hash is: ")
	fmt.Printf("%x", hasher.Sum(nil))
	// New line for readability
	fmt.Println()
}

// The menu for file system tasks
func filesMain() {

	fileOptions := map[int]string{
		0: "Main Menu",
		1: "Create multiple directories",
		2: "Calculate file hash",
	}

	// Display options
	for option, description := range fileOptions {
		fmt.Printf("%d : %s \n", option, description)
	}

	// Create a new reader and scanner
	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	// Get a choice from the user via stdin
	fmt.Print("Enter an int from the options above: ")
	scanner.Scan()

	userChoice, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if userChoice == 0 {
		main()
	} else if userChoice == 1 {
		fmt.Print("Enter a location to create dirs: ")
		userLocation, _ := reader.ReadString('\n')
		// Remove the newline before passing into createDirs() function
		formattedUserLocation := strings.TrimSuffix(userLocation, "\n")
		os.Chdir(formattedUserLocation)
		fmt.Println("Enter how many directories you want to create: ")
		userDirNum, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		createDirs(userDirNum)

		main()
	} else if userChoice == 2 {
		fmt.Print("Enter the absolute path of the file: ")
		userFile, _ := reader.ReadString('\n')
		// Remove the newline before passing into getData() function
		formattedFilePath := strings.TrimSuffix(userFile, "\n")
		calcHash(formattedFilePath)

		main()
	} else {
		fmt.Print("*** Invalid option. Returning to main menu. ***")
	}

}
