**Eclean identifies records you should delete or fix. This script does not change the original file, it parses an Engaging Networks csv and creates other csvs with the results.**

## How to install

* Download the .zip file for your operating system and unzip the executable binary file.
* Optionally copy the file to a folder in your command line PATH. You will be able to use the script in any folder.


## How to use

### 1 - Export your users

First export a user data or hybrid csv with all your users and **at least** this fields:
* email
* first_name
* last_name 

### 2 - Run the script

Run the script from the command line with the desired checks as flags in the command. 

The following example will check invalid emails and some issues with names in the file `your_EN_file_to_check.csv`.

In Mac OS / Linux:

```
./eclean -input=your_EN_file_to_check.csv -emails -fake -empty -suppressed -delete
```

In Windows: 

```
./eclean.exe -input=your_EN_file_to_check.csv -emails -fake -empty -suppressed -delete
```

### List of checks you can do

`-emails` - Creates a csv file with records that don't respect the email regular expression. This emails can be fixed or deleted. Saved in `eclean_INVALID_EMAILS.csv`

`-fake` - Creates a csv file with records with weird first and last names. You should inspect them and delete them as they are not real users. Saved in `eclean_FAKE_NAMES.csv`

`-empty` - Creates a csv file with empty first names and last names. Storing email addresses without first and last names is not recommended. This should be inspected and can be deleted if there's not the risk of being added again to EN (by the CRM). Saved in `eclean_EMPTY_NAMES.csv`

`-suppressed` - Creates a csv file with suppressed email addresses. This should be inspected and can be deleted if there's not the risk of being added again to EN (by the CRM). Saved in `eclean_SUPPRESSED_EMAILS.csv`

### Options you can add

`-delete` - Creates files with email addresses **only** and without header. This simpler files can be uploaded in Engaging Networks if you want to delete the records.   

### Other 

`./eclean -help` - Display help.

`./eclean -trash` - Delete all files created by eclean in the current folder.

### 3 - Results

The script will generate csv files on the current folder with filenames prefixed by `eclean_`. 

Inspect this files and, if all is OK, add `-delete` to your command to obtain files with the email addresses only (one per line). To delete the records upload its email addresses in *Supporter Data &gt; Delete Supporters*.

## To do

Features to develop in the near future.

Records that can be removed:

* Differentiate **suppressed** donors from other suppressed emails.

Standardize field values in records:

* Removes all non numeric digits from the **phone_number** field
* Removes all non valid digits from the Spanish **id_number** field

## Install from the source code

This script was developed in **[Go](https://golang.org/)**.

To download and install **eclean** and it's dependencies you must have **Go** installed and run in the command line:

```
go get github.com/greenpeace/eclean
go install github.com/greenpeace/eclean
```
