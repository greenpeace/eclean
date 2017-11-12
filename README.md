# Eclean

**Eclean identifies records you should delete or fix. This script does not change the original file, it parses an Engaging Networks csv and creates other csvs with the results.**

This script is **fast**. It processes 4 queries over 1,300,000 records in about 11 seconds (in a laptop computer). Optionally it formats the files with one email address per line, ready to delete. 

How long will you take to do it using the classic process? Do the 4 queries, download the results in 4 separate csv files and remove the extra columns in each file?

## How to install

* Download the [.zip file](https://github.com/greenpeace/gpes-eclean/releases) for your operating system and unzip the executable binary file.
* Optionally copy the file to a folder in your command line PATH. You will be able to use the script in any folder.

## How to use

### 1 - Export your users

First export a user data or hybrid csv with all your users and **at least** this fields:

* email
* first_name
* last_name
* Suppressed
* contact_codes (sync with Salesforce)

### 2 - Run the script

Run the script from the command line with the desired checks as flags in the command. 

The following example will check invalid emails and some issues with names in the file `your_EN_file_to_check.csv`.

In Mac OS / Linux:

```bash
./eclean -input=your_EN_file_to_check.csv -emails -fake -empty -suppressed -delete
```

In Windows:

```bash
./eclean.exe -input=your_EN_file_to_check.csv -emails -fake -empty -suppressed -delete
```

### List of checks you can do

`-emails` - Creates a csv file with records that don't respect the email regular expression. This emails can be fixed or deleted. Saved in `eclean_INVALID_EMAILS.csv`

`-fake` - Creates a csv file with records with weird first and last names. You should inspect them and delete them as they are not real users. Saved in `eclean_FAKE_NAMES.csv`

`-empty` - Creates a csv file with empty first names and last names. Storing email addresses without first and last names is not recommended. This should be inspected and can be deleted if there's not the risk of being added again to EN (by the CRM). Saved in `eclean_EMPTY_NAMES.csv`

`-suppressed` - Creates a csv file with all the suppressed email addresses. This should be inspected and can be deleted if there's not the risk of being added again to EN (by the CRM). Saved in `eclean_SUPPRESSED_EMAILS.csv`. IMPORTANT: If the `contact_codes` field exists in the exported csv, a second and third files `eclean_SUPPRESSED_EMAILS_CONTACTS.csv` and `eclean_SUPPRESSED_EMAILS_LEADS.csv` are also created.

### Options you can add

`-delete` - Creates files with email addresses **only** and without header. This simpler files can be uploaded in Engaging Networks if you want to delete the records.

### Other

`./eclean -help` - Display help.

`./eclean -stash` - Renames the last files created by eclean by adding a timestamp.

`./eclean -trash` - Delete all files created by eclean in the current folder without a timestamp.

### 3 - Results

**VERY IMPORTANT!!!** - **Think** carefully on the implications of deleting information from EN. **Consult** with your colleagues in your office and the community. Always keep at least one **backup** and triple-check everything. Deleting email addresses will affect your number of signups.

This script will generate csv files on the current folder with filenames prefixed by `eclean_`. 

Inspect this files and, if all is OK, add `-delete` to your command to obtain files with the email addresses only (one per line). To delete the records upload its email addresses in *Supporter Data &gt; Delete Supporters*.

## To do

Features to develop in the near future.

Standardise field values in records:

* Removes all non numeric digits from the **phone_number** field
* Removes all non valid digits from the Spanish **id_number** field

## Install from the source code

This script was developed in **[Go](https://golang.org/)**.

To download and install **eclean** and it's dependencies you must have **Go** installed and run in the command line:

```bash
go get github.com/greenpeace/gpes-eclean
go install github.com/greenpeace/gpes-eclean
```
