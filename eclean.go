package main

import (
	"flag"
	"fmt"
)

const emailRegex string = `([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)`

const optOutFieldName = "real_email_ok"

// Use this to configure the field that separates contacts from leads. Normally contact_codes, but you may want to work with another field
const contactFieldName = "contact_codes"
const contactRegularExp = `(\w+|\s)`

// const contactFieldName = "tipo"
// const contactRegularExp = `(2|3|4|5)`

func main() {

	help := flag.Bool("help", false, "Display help")
	trash := flag.Bool("trash", false, "Delete files")
	stash := flag.Bool("stash", false, "Rename files created by this script")
	fileName := flag.String("input", "your_EN_file_to_check.csv", "File to do the operations")
	emails := flag.Bool("emails", false, "Check for invalid email addresses")
	fake := flag.Bool("fake", false, "Check for records that might be fake")
	empty := flag.Bool("empty", false, "Check for records that have empty first and last names")
	suppressed := flag.Bool("suppressed", false, "Check for records that have suppressed emails")
	optOut := flag.Bool("optout", false, "Check for records that have opt outs emails")
	deleteIt := flag.Bool("delete", false, "Return the emails only, without header ")
	flag.Parse()

	if *help == true {
		displayHelp()
	} else if *stash == true {
		stashFiles()
	} else if *trash == true {
		trashFiles()
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

		if *suppressed == true {
			suppresedEmails(x, deleteIt)
			fmt.Printf("\n")
		}
		if *optOut == true {
			optOutEmails(x, deleteIt)
			fmt.Printf("\n")
		}

	}
}
