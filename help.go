package main

import (
	"fmt"
)

func displayHelp() {
	helpText := `
Eclean identifies records you should delete or fix. This script does not change the original file, it parses an Engaging Networks csv and creates other csvs with the results.

Example:

./eclean -input=your_EN_file_to_check.csv -emails -fake -delete

List of checks you can do:

-emails - Creates a file with records that don't respect the email regular expression. This emails can be fixed or deleted. Saved in eclean_INVALID_EMAILS.csv

-fake - Creates a file with records with weird first and last names. You should inspect them and delete them as they are not real users. Saved in eclean_FAKE_NAMES.csv

-empty - Creates a csv with empty first names and last names. This whould be inspected and can be deleted if there's not the risk of being added again by the CRM. Saved in eclean_EMPTY_NAMES.csv

Options you can add:

-delete - Creates files with email addresses only and without header. This simpler files can be uploaded in Engaging Networks if you want to delete the records.  


`

	fmt.Println(helpText)
}
