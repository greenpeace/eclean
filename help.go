package main

import (
	"fmt"
)

func displayHelp() {
	helpText := `
Eclean identifies records you should delete or fix. This script does not change the original file, it parses an Engaging Networks csv and creates other csvs with the results.

EXAMPLE:

./eclean -input=your_EN_file_to_check.csv -emails -fake -empty -suppressed -optout -delete

LIST OF CHECKS YOU CAN DO:

-emails - Creates a file with records that don't respect the email regular expression. This emails can be fixed or deleted. Saved in eclean_INVALID_EMAILS.csv

-fake - Creates a file with records with weird first and last names. You should inspect them and delete them as they are not real users. Saved in eclean_FAKE_NAMES.csv

-empty - Creates a csv with empty first names and last names. Storing email addresses without first and last names is not recommended. This should be inspected and can be deleted if there's not the risk of being added again by the CRM. Saved in eclean_EMPTY_NAMES.csv

-suppressed - Creates a csv file with all the suppressed email addresses. This should be inspected and can be deleted if there's not the risk of being added again to EN (by the CRM). Saved in eclean_SUPPRESSED_EMAILS.csv. IMPORTANT: If the contact_codes field exists in the exported csv, a second and third files eclean_SUPPRESSED_EMAILS_CONTACTS.csv and eclean_SUPPRESSED_EMAILS_LEADS.csv are also created.

-optout - Creates a CSV file with all the opt-out email addresses

OPTIONS YOU CAN ADD:

-delete - Creates files with email addresses only and without header. This simpler files can be uploaded in Engaging Networks if you want to delete the records.  

OTHER:

-help - Display help.

-trash - Delete all files created by eclean in the current folder.

-stash - Renames all files created by eclean by adding a timestamp.

`

	fmt.Println(helpText)
}
