package main

import (
	"fmt"
	"os"
	"time"

	"github.com/osvik/simplecsv"
)

// Read the csv file
func readCsvFile(fileName *string) simplecsv.SimpleCsv {
	fmt.Println("Csv file name:", *fileName)
	if _, err := os.Stat(*fileName); os.IsNotExist(err) {
		fmt.Println("ERROR: The file/path", *fileName, "does not exist here")
		os.Exit(-1)
	}
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
	if x.GetHeaderPosition(optOutFieldName) == -1 {
		fmt.Println("WARNING - There's not an opt-out field in the CSV. The CSV should have a column named", optOutFieldName)
	} else {
		fmt.Println("OK - Opt-out field found, it's the", optOutFieldName, "field")
	}
	if x.GetHeaderPosition(contactFieldName) == -1 {
		fmt.Println("ERR - There's not a", contactFieldName, "field in the CSV")
	} else {
		fmt.Println("OK -", contactFieldName, "field found")
	}
}

// Creates a Csv with invalid email addresses
func invalidEmailAddresses(x simplecsv.SimpleCsv, deleteFormat *bool) {
	var fieldsList []string
	if *deleteFormat == false {
		fieldsList = x.GetHeaders()
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
		fieldsList = x.GetHeaders()
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

// Creates a Csv with empty names
func emptyNames(x simplecsv.SimpleCsv, deleteFormat *bool) {
	var fieldsList []string
	if *deleteFormat == false {
		fieldsList = x.GetHeaders()
	} else {
		fieldsList = []string{"email"}
	}

	if x.GetHeaderPosition("first_name") != -1 && x.GetHeaderPosition("last_name") != -1 {
		fmt.Println("first_name and last_name fields found")
		invalidFirstNameIndex, invalidFirstNameIndexOK := x.FindInField("first_name", "")
		invalidLastNameIndex, invalidLastNameIndexOK := x.FindInField("last_name", "")
		if invalidFirstNameIndexOK == true || invalidLastNameIndexOK == true {
			invalidNamesIndex := simplecsv.AndIndex(invalidFirstNameIndex, invalidLastNameIndex)
			fmt.Println("Number of records with empty names:", len(invalidNamesIndex))
			invalidNamesCsv, _ := x.OnlyThisRows(invalidNamesIndex, true)
			invalidNamesCsv, _ = invalidNamesCsv.OnlyThisFields(fieldsList)
			if *deleteFormat == true {
				invalidNamesCsv, _ = invalidNamesCsv.DeleteRow(0)
			}
			wasWritten := invalidNamesCsv.WriteCsvFile("eclean_EMPTY_NAMES.csv")
			if wasWritten == true {
				fmt.Println("Fake names saved in the file: eclean_EMPTY_NAMES.csv")
			} else {
				fmt.Println("Could not create eclean_EMPTY_NAMES.csv")
			}
		} else {
			fmt.Println("Problems with fake names index")
		}

	} else {
		fmt.Println("There's not a first_name and last_name fields in the csv")
	}

}

// Removes all the result files
func trashFiles() {
	os.Remove("eclean_INVALID_EMAILS.csv")
	os.Remove("eclean_FAKE_NAMES.csv")
	os.Remove("eclean_EMPTY_NAMES.csv")
	os.Remove("eclean_SUPPRESSED_EMAILS.csv")
	os.Remove("eclean_SUPPRESSED_EMAILS_CONTACTS.csv")
	os.Remove("eclean_SUPPRESSED_EMAILS_LEADS.csv")
	os.Remove("eclean_OPT-OUT_EMAILS.csv")
	os.Remove("eclean_OPT-OUT_EMAILS_CONTACTS.csv")
	os.Remove("eclean_OPT-OUT_EMAILS_LEADS.csv")
}

// nowDateTimeString Returns the date as a string in a specific format
func nowDateTimeString() string {
	t := time.Now()
	return t.Format("2006-01-02-15h04m05")
}

// stashFiles Renames files created by eclean by adding the date+time
func stashFiles() {
	now := nowDateTimeString()
	os.Rename("eclean_INVALID_EMAILS.csv", "eclean_INVALID_EMAILS-"+now+".csv")
	os.Rename("eclean_FAKE_NAMES.csv", "eclean_FAKE_NAMES-"+now+".csv")
	os.Rename("eclean_EMPTY_NAMES.csv", "eclean_EMPTY_NAMES-"+now+".csv")
	os.Rename("eclean_SUPPRESSED_EMAILS.csv", "eclean_SUPPRESSED_EMAILS-"+now+".csv")
	os.Rename("eclean_SUPPRESSED_EMAILS_CONTACTS.csv", "eclean_SUPPRESSED_EMAILS_CONTACTS-"+now+".csv")
	os.Rename("eclean_SUPPRESSED_EMAILS_LEADS.csv", "eclean_SUPPRESSED_EMAILS_LEADS-"+now+".csv")
	os.Rename("eclean_OPT-OUT_EMAILS.csv", "eclean_OPT-OUT_EMAILS-"+now+".csv")
	os.Rename("eclean_OPT-OUT_EMAILS_CONTACTS.csv", "eclean_OPT-OUT_EMAILS_CONTACTS-"+now+".csv")
	os.Rename("eclean_OPT-OUT_EMAILS_LEADS.csv", "eclean_OPT-OUT_EMAILS_LEADS-"+now+".csv")

	os.Exit(0)
}

// Creates a csv with the supressed emails
func suppresedEmails(x simplecsv.SimpleCsv, deleteFormat *bool) {
	var suppressedEmailIndex []int
	var suppressedEmailIndexOK bool
	var fieldsList []string
	if *deleteFormat == false {
		fieldsList = x.GetHeaders()
	} else {
		fieldsList = []string{"email"}
	}
	if x.GetHeaderPosition("Suppressed") != -1 {
		fmt.Println("Suppressed field found")
		suppressedEmailIndex, suppressedEmailIndexOK = x.FindInField("Suppressed", "Y")
		if suppressedEmailIndexOK == true {
			fmt.Println("Number of records with suppressed emails:", len(suppressedEmailIndex))
			suppressedEmailsCsv, _ := x.OnlyThisFields(fieldsList)
			suppressedEmailsCsv, _ = suppressedEmailsCsv.OnlyThisRows(suppressedEmailIndex, true)
			if *deleteFormat == true {
				suppressedEmailsCsv, _ = suppressedEmailsCsv.DeleteRow(0)
			}
			wasWritten := suppressedEmailsCsv.WriteCsvFile("eclean_SUPPRESSED_EMAILS.csv")
			if wasWritten == true {
				fmt.Println("Suppressed emails saved in the file: eclean_SUPPRESSED_EMAILS.csv")
			} else {
				fmt.Println("Could not create eclean_SUPPRESSED_EMAILS.csv")
			}
		} else {
			fmt.Println("Problems with supressed index")
		}
	} else {
		fmt.Println("There's not a Suppressed field in the csv")
	}

	if x.GetHeaderPosition("Suppressed") != -1 && x.GetHeaderPosition(contactFieldName) != -1 {
		fmt.Println(contactFieldName, "field found")

		lastRecord := x.GetNumberRows() - 1
		contactsIndex, contactsIndexOK := x.MatchInField(contactFieldName, contactRegularExp)
		leadsIndex := simplecsv.NotIndex(contactsIndex, 1, lastRecord)

		if contactsIndexOK == true && suppressedEmailIndexOK == true {
			suppressedEmailsContactsIndex := simplecsv.AndIndex(contactsIndex, suppressedEmailIndex)
			suppressedEmailsLeadsIndex := simplecsv.AndIndex(leadsIndex, suppressedEmailIndex)

			fmt.Println("Number of SF contacts with suppressed emails:", len(suppressedEmailsContactsIndex))
			fmt.Println("Number of SF leads with suppressed emails:", len(suppressedEmailsLeadsIndex))

			if *deleteFormat == false {
				fieldsList = x.GetHeaders()
			} else {
				fieldsList = []string{"email"}
			}
			suppressedContactsEmailsCsv, _ := x.OnlyThisFields(fieldsList)
			suppressedLeadsEmailsCsv, _ := x.OnlyThisFields(fieldsList)

			suppressedContactsEmailsCsv, _ = suppressedContactsEmailsCsv.OnlyThisRows(suppressedEmailsContactsIndex, true)
			suppressedLeadsEmailsCsv, _ = suppressedLeadsEmailsCsv.OnlyThisRows(suppressedEmailsLeadsIndex, true)

			if *deleteFormat == true {
				suppressedContactsEmailsCsv, _ = suppressedContactsEmailsCsv.DeleteRow(0)
				suppressedLeadsEmailsCsv, _ = suppressedLeadsEmailsCsv.DeleteRow(0)
			}
			wasWrittenContacts := suppressedContactsEmailsCsv.WriteCsvFile("eclean_SUPPRESSED_EMAILS_CONTACTS.csv")
			if wasWrittenContacts == true {
				fmt.Println("Suppressed emails from contacts saved in the file: eclean_SUPPRESSED_EMAILS_CONTACTS.csv")
			} else {
				fmt.Println("Could not create eclean_SUPPRESSED_EMAILS_CONTACTS.csv")
			}

			wasWrittenLeads := suppressedLeadsEmailsCsv.WriteCsvFile("eclean_SUPPRESSED_EMAILS_LEADS.csv")
			if wasWrittenLeads == true {
				fmt.Println("Suppressed emails from leads saved in the file: eclean_SUPPRESSED_EMAILS_LEADS.csv")
			} else {
				fmt.Println("Could not create eclean_SUPPRESSED_EMAILS_LEADS.csv")
			}

		}
	}

}

// Creates a csv with the opt-out emails
func optOutEmails(x simplecsv.SimpleCsv, deleteFormat *bool) {
	fmt.Println("Analising opt-out emails")
	var optOutsEmailIndex []int
	var optOutsEmailIndexOK bool
	// Defining the fields in the col
	var fieldsList []string
	if *deleteFormat == false {
		fieldsList = x.GetHeaders()
	} else {
		fieldsList = []string{"email"}
	}

	// It the opt out field exists
	if x.GetHeaderPosition(optOutFieldName) != -1 {
		fmt.Println("Opt-out field found, it's", optOutFieldName)
		optOutsEmailIndex, optOutsEmailIndexOK = x.FindInField(optOutFieldName, "N")

		if optOutsEmailIndexOK == true {
			fmt.Println("Number of records with opt-outs emails:", len(optOutsEmailIndex))
			optOutsEmailsCsv, _ := x.OnlyThisFields(fieldsList)
			optOutsEmailsCsv, _ = optOutsEmailsCsv.OnlyThisRows(optOutsEmailIndex, true)
			if *deleteFormat == true {
				optOutsEmailsCsv, _ = optOutsEmailsCsv.DeleteRow(0)
			}
			wasWritten := optOutsEmailsCsv.WriteCsvFile("eclean_OPT-OUT_EMAILS.csv")
			if wasWritten == true {
				fmt.Println("Opt-out emails saved in the file: eclean_OPT-OUT_EMAILS.csv")
			} else {
				fmt.Println("Could not create eclean_OPT-OUT_EMAILS.csv")
			}
		} else {
			fmt.Println("Problems with opt-outs index")
		}
	} else { // If the opt-out field does not exist
		fmt.Println("There's not an opt-out field in the csv. It should be", optOutFieldName)
	}

	// It the opt out and contact_codes fields exist
	if x.GetHeaderPosition(optOutFieldName) != -1 && x.GetHeaderPosition(contactFieldName) != -1 {
		fmt.Println(contactFieldName, "field found")

		lastRecord := x.GetNumberRows() - 1
		contactsIndex, contactsIndexOK := x.MatchInField(contactFieldName, contactRegularExp)
		leadsIndex := simplecsv.NotIndex(contactsIndex, 1, lastRecord)

		if contactsIndexOK == true && optOutsEmailIndexOK == true {
			optOutsEmailsContactsIndex := simplecsv.AndIndex(contactsIndex, optOutsEmailIndex)
			optOutsEmailsLeadsIndex := simplecsv.AndIndex(leadsIndex, optOutsEmailIndex)

			fmt.Println("Number of SF contacts with opt-out emails:", len(optOutsEmailsContactsIndex))
			fmt.Println("Number of SF leads with opt-out emails:", len(optOutsEmailsLeadsIndex))

			if *deleteFormat == false {
				fieldsList = x.GetHeaders()
			} else {
				fieldsList = []string{"email"}
			}

			optOutsContactsEmailsCsv, _ := x.OnlyThisFields(fieldsList)
			optOutsLeadsEmailsCsv, _ := x.OnlyThisFields(fieldsList)

			optOutsContactsEmailsCsv, _ = optOutsContactsEmailsCsv.OnlyThisRows(optOutsEmailsContactsIndex, true)
			optOutsLeadsEmailsCsv, _ = optOutsLeadsEmailsCsv.OnlyThisRows(optOutsEmailsLeadsIndex, true)

			if *deleteFormat == true {
				optOutsContactsEmailsCsv, _ = optOutsContactsEmailsCsv.DeleteRow(0)
				optOutsLeadsEmailsCsv, _ = optOutsLeadsEmailsCsv.DeleteRow(0)
			}

			wasWrittenContacts2 := optOutsContactsEmailsCsv.WriteCsvFile("eclean_OPT-OUT_EMAILS_CONTACTS.csv")
			if wasWrittenContacts2 == true {
				fmt.Println("Opt-out emails from contacts saved in the file: eclean_OPT-OUT_EMAILS_CONTACTS.csv")
			} else {
				fmt.Println("Could not create eclean_OPT-OUT_EMAILS_CONTACTS.csv")
			}
			wasWrittenLeads2 := optOutsLeadsEmailsCsv.WriteCsvFile("eclean_OPT-OUT_EMAILS_LEADS.csv")
			if wasWrittenLeads2 == true {
				fmt.Println("Opt-out emails from leads saved in the file: eclean_OPT-OUT_EMAILS_LEADS.csv")
			} else {
				fmt.Println("Could not create eclean_OPT-OUT_EMAILS_LEADS.csv")
			}
		}

	}

}
