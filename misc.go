package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Get data from REST API endpoint, and save it to a file
func getData(endpoint string, filename string) {

	// make a GET request to fetch the JSON data
	res, err := http.Get(endpoint)

	// Log any errors
	if err != nil {
		log.Fatal(err)
	}

	// store the response body in a variable
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Create a file to write to
	file, err := os.Create("./data/" + filename + ".json")
	// Write the JSON to an external file
	file.WriteString(string(body))

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Download complete. Check the 'data' directory.")
	}

}

// Check if a password is in a data breach via HaveIBeenPwned
func checkPassword(password string) {

	// Convert the password string to bytes, so they can be hashed
	passwordBytes := []byte(password)
	// Hash the password
	passwordHash := sha1.Sum(passwordBytes)
	// Convert the hash bytes to a string
	hashString := hex.EncodeToString(passwordHash[:])

	//fmt.Printf("Your hash: %x", passwordHash)
	// Make a request to the HaveIBeenPwned REST API
	res, err := http.Get("https://api.pwnedpasswords.com/range/" + hashString[:5])

	if err != nil {
		log.Fatal(err)
	}

	// Parse the response body
	body, err := ioutil.ReadAll(res.Body)
	bodyString := string(body)
	hashesArray := strings.Split(bodyString, ":")

	if err != nil {
		log.Fatal(err)
	}

	tail := hashString[5:]

	// FIXME: #1 Figure out why it always returns false
	search := func(userPassword string, list []string) bool {
		fmt.Println(tail)
		for _, password := range list {

			fmt.Println(password)
			/*
				truncHash := password[5:]
				if truncHash == tail {
					return true
				} else {
					return false
				}
			*/
		}
		return false
	}

	fmt.Println(search(tail, hashesArray))
	/*
		fmt.Println(hashesArray)
		fmt.Println(hashString)
		fmt.Println(hashString[5:])
	*/
}

func miscMain() {

	fileOptions := map[int]string{
		0: "Main menu",
		1: "Download JSON data from a REST API",
		2: "Check if password has been compromised in a data breach",
	}

	// Display options
	for option, description := range fileOptions {
		fmt.Printf("%d : %s \n", option, description)
	}

	// Create a scanner for menu selections
	scanner := bufio.NewScanner(os.Stdin)
	// Create a new reader for other values from stdin
	reader := bufio.NewReader(os.Stdin)

	// Get a choice from the user via stdin
	fmt.Print("Enter an int from the options above: ")
	scanner.Scan()

	userChoice, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if userChoice == 0 {
		main()
	} else if userChoice == 1 {
		fmt.Print("Enter a URL: ")
		url, _ := reader.ReadString('\n')
		// Remove the newline before passing into getData() function
		formattedUrl := strings.TrimSuffix(url, "\n")
		fmt.Print("Enter a file name (without a file extension): ")
		filename, _ := reader.ReadString('\n')
		formattedFilename := strings.TrimSuffix(filename, "\n")

		getData(formattedUrl, formattedFilename)

		fmt.Println()
		main()
	} else if userChoice == 2 {
		fmt.Print("Enter a password: ")
		userPass, _ := reader.ReadString('\n')
		formattedPass := strings.TrimSuffix(userPass, "\n")

		checkPassword(formattedPass)

		fmt.Println()
		main()
	} else {
		fmt.Println("*** Invalid option. ***")
		main()
	}
}
