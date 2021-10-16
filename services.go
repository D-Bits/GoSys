/*
* Functions for managing Linux services via systemctl
 */
package main

import (
	"log"
	"os/exec"
)

// Function for managing Postgres
func mngPostgres(option int) {

	if option == 0 {
		main()
	} else if option == 1 {
		cmdArgs := exec.Command(
			"sudo",
			"systemctl",
			"start",
			"postgressql",
		)
		err := cmdArgs.Run()
		if err != nil {
			log.Fatal(err)
		}

	} else if option == 2 {
		cmdArgs := exec.Command(
			"sudo",
			"systemctl",
			"start",
			"postgressql",
		)
		err := cmdArgs.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func servicesMain(option int) {

}
