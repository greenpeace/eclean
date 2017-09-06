package main

import (
	"fmt"

	"github.com/osvik/simplecsv"
)

// Read the csv file
func readCsvFile(fileName *string) simplecsv.SimpleCsv {
	fmt.Println("Csv file name:", *fileName)
	var x simplecsv.SimpleCsv
	var fileRead bool
	x, fileRead = simplecsv.ReadCsvFile(*fileName)
	if fileRead == true {
		fmt.Println("Sucessfuly read the csv file")
	} else {
		fmt.Println("Error reading the csv file?")
	}

	return x
}

// Prints the headers, one per column
func printHeaders(x simplecsv.SimpleCsv, fileName *string) {
	headers := x.GetHeaders()
	fmt.Println("Csv headers in", *fileName, ":")
	for i, v := range headers {
		fmt.Println(i, ":", v)
	}
}

// Checks if the recommended fields exist in the file and print message
func checkRecommendedFields(x simplecsv.SimpleCsv) {
	fmt.Println("Fields pre-check:")
	if x.GetHeaderPosition("Supporter ID") == -1 {
		fmt.Println("ERR - There's not a Supporter ID field in the CSV")
	} else {
		fmt.Println("OK - Supporter ID field found")
	}
	if x.GetHeaderPosition("email") == -1 {
		fmt.Println("ERR - There's not an email field in the CSV")
	} else {
		fmt.Println("OK - Email field found")
	}
	if x.GetHeaderPosition("first_name") == -1 || x.GetHeaderPosition("last_name") == -1 {
		fmt.Println("ERR - There's not a first_name or last_name field in the CSV")
	} else {
		fmt.Println("OK - first_name and last_name fields found")
	}
	if x.GetHeaderPosition("Suppressed") == -1 {
		fmt.Println("ERR - There's not a Suppressed field in the CSV")
	} else {
		fmt.Println("OK - Suppressed field found")
	}
}

// Creates a Csv with invalid email addresses
func invalidEmailAddresses(x simplecsv.SimpleCsv, deleteFormat *bool) {
	var fieldsList []string
	if *deleteFormat == false {
		fieldsList = []string{"Supporter ID", "email"}
	} else {
		fieldsList = []string{"email"}
	}

	if x.GetHeaderPosition("email") != -1 {
		fmt.Println("Email address field found")
		validEmailsIndex, emailIndexOk := x.MatchInField("email", emailRegex)
		fmt.Println("Number of valid email addresses:", len(validEmailsIndex), "in", x.GetNumberRows()-1)
		if emailIndexOk == true {
			lastRecord := x.GetNumberRows() - 1
			invalidEmailsIndex := simplecsv.NotIndex(validEmailsIndex, 1, lastRecord)
			fmt.Println("Number of invalid email addresses:", len(invalidEmailsIndex))
			invalidEmailsCsv, _ := x.OnlyThisFields(fieldsList)
			invalidEmailsCsv, _ = invalidEmailsCsv.OnlyThisRows(invalidEmailsIndex, true)
			if *deleteFormat == true {
				invalidEmailsCsv, _ = invalidEmailsCsv.DeleteRow(0)
			}
			wasWritten := invalidEmailsCsv.WriteCsvFile("eclean_INVALID_EMAILS.csv")
			if wasWritten == true {
				fmt.Println("Invalid emails saved in the file: eclean_INVALID_EMAILS.csv")
			} else {
				fmt.Println("Could not create the file eclean_INVALID_EMAILS.csv")
			}
		} else {
			fmt.Println("Problems with emails index.")
		}
	} else {
		fmt.Println("There's not an email field in the csv")
	}

}

// Creates a Csv with fake names
func fakeNames(x simplecsv.SimpleCsv, deleteFormat *bool) {
	var fieldsList []string
	if *deleteFormat == false {
		fieldsList = []string{"Supporter ID", "email", "first_name", "last_name"}
	} else {
		fieldsList = []string{"email"}
	}

	if x.GetHeaderPosition("first_name") != -1 && x.GetHeaderPosition("last_name") != -1 {
		fmt.Println("first_name and last_name fields found")
		invalidFirstNameIndex, invalidFirstNameIndexOK := x.MatchInField("first_name", `^\d{2}`)
		invalidLastNameIndex, invalidLastNameIndexOK := x.MatchInField("last_name", `^\d{2}`)
		if invalidFirstNameIndexOK == true || invalidLastNameIndexOK == true {
			invalidNamesIndex := simplecsv.OrIndex(invalidFirstNameIndex, invalidLastNameIndex)
			fmt.Println("Number of records with fake names:", len(invalidNamesIndex))
			invalidNamesCsv, _ := x.OnlyThisRows(invalidNamesIndex, true)
			invalidNamesCsv, _ = invalidNamesCsv.OnlyThisFields(fieldsList)
			if *deleteFormat == true {
				invalidNamesCsv, _ = invalidNamesCsv.DeleteRow(0)
			}
			wasWritten := invalidNamesCsv.WriteCsvFile("eclean_FAKE_NAMES.csv")
			if wasWritten == true {
				fmt.Println("Fake names saved in the file: eclean_FAKE_NAMES.csv")
			} else {
				fmt.Println("Could not create eclean_FAKE_NAMES.csv")
			}
		} else {
			fmt.Println("Problems with fake names index")
		}
	} else {
		fmt.Println("There's not a first_name and last_name fields in the csv")
	}

}