package main

import (
	"fmt"
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

	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	file, err := os.OpenFile(inputFilename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	checklistItems := []*ChecklistItem{}

	if err := gocsv.UnmarshalFile(file, &checklistItems); err != nil { // Load clients from file
		panic(err)
	}
	for _, item := range checklistItems {
		fmt.Println(item)
	}
	fmt.Println(checklistItems[0])
	fmt.Println("-------------------------------------")

	// New list to print a new csv file
	var newCSVList []ChecklistItem

	//build new list
	for i := 0; i < len(checklistItems); i++ {
		formArray := strings.Split(checklistItems[i].Form, "\n")
		formSize := len(formArray)

		for f := 0; f < formSize; f++ {
			newCSVList = append(newCSVList, ChecklistItem{
				Category:  checklistItems[i].Category,
				FormGroup: checklistItems[i].FormGroup,
				Form:      formArray[f],
			})
		}

	}

	outfile, err := os.Create(outputFilename)
	gocsv.MarshalFile(&newCSVList, outfile)
}
