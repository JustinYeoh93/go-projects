package main

import "strings"

func processRow(text string) (firstName, fullName, month string) {
	row := strings.Split(text, "|")

	// extract full name
	fullName = strings.Replace(strings.TrimSpace(row[7]), " ", "", -1)

	// extract first name
	name := strings.TrimSpace(row[7])
	if name != "" {
		startOfName := strings.Index(name, ", ") + 2
		if endOfName := strings.Index(name[startOfName:], " "); endOfName < 0 {
			firstName = name[startOfName:]
		} else {
			firstName = name[startOfName : startOfName+endOfName]
		}
		if strings.HasSuffix(firstName, ",") {
			firstName = strings.Replace(firstName, ",", "", -1)
		}
	}

	// extract month
	date := strings.TrimSpace(row[13])
	if len(date) == 8 {
		month = date[:2]
	} else {
		month = "--"
	}

	return firstName, fullName, month
}
