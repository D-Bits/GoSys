package main

import (
	"bufio"
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

func miscMain() {

	fileOptions := map[int]string{
		0: "Main Menu",
		1: "Download JSON Data from a REST API",
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
	} else {
		fmt.Println("*** Invalid option. ***")
	}

}
