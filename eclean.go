package main

import (
	"flag"
	"fmt"
)

const emailRegex string = `([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)`

func main() {

	help := flag.Bool("help", false, "Display help")
	fileName := flag.String("input", "your_EN_file_to_check.csv", "File to do the operations")
	emails := flag.Bool("emails", false, "Check for invalid email addresses")
	fake := flag.Bool("fake", false, "Check for records that might be fake")
	empty := flag.Bool("empty", false, "Check for records that have empty first and last names")
	deleteIt := flag.Bool("delete", false, "Return the emails only, without header ")
	flag.Parse()

	if *help == true {
		displayHelp()
	} else {

		fmt.Printf("\nREPORT\n\n")

		x := readCsvFile(fileName)

		fmt.Printf("\n")

		printHeaders(x, fileName)

		fmt.Printf("\n")

		fmt.Println("Total number of records:", x.GetNumberRows()-1)

		fmt.Printf("\n")

		checkRecommendedFields(x)

		fmt.Printf("\n")

		if *emails == true {
			invalidEmailAddresses(x, deleteIt)
			fmt.Printf("\n")
		}

		if *fake == true {
			fakeNames(x, deleteIt)
			fmt.Printf("\n")
		}

		if *empty == true {
			emptyNames(x, deleteIt)
			fmt.Printf("\n")
		}

	}
}
