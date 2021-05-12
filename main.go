package main

import (
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type ChecklistItem struct {
	Category  string `csv:"Category"`
	FormGroup string `csv:"FormGroup"`
	Form      string `csv:"Form"`
}

func main() {

	// grab user arguments
	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	// open file
	file, err := os.OpenFile(inputFilename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read into collection
	checklistItems := []*ChecklistItem{}
	if err := gocsv.UnmarshalFile(file, &checklistItems); err != nil { // Load clients from file
		panic(err)
	}

	//build the new list
	var newCSVList []ChecklistItem
	newCSVList = splitData(checklistItems)

	// print to new file
	outfile, err := os.Create(outputFilename)
	gocsv.MarshalFile(&newCSVList, outfile)
}

func splitData(items []*ChecklistItem) []ChecklistItem {
	// New list to print a new csv file
	var newCSVList []ChecklistItem

	//build new list
	for i := 0; i < len(items); i++ {
		formArray := strings.Split(items[i].Form, "\n")
		formSize := len(formArray)

		for f := 0; f < formSize; f++ {
			newCSVList = append(newCSVList, ChecklistItem{
				Category:  items[i].Category,
				FormGroup: items[i].FormGroup,
				Form:      formArray[f],
			})
		}

	}

	return newCSVList
}
