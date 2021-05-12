package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func lookupCSV(filePath string) {
	records := csvReader(filePath)
	fmt.Println(records)
}

func csvReader(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("An error encountered ::", err)
		log.Fatal(err)
	}
	//init the csv reader
	reader := csv.NewReader(file)
	//reader.Comma = ' '

	//read all records into a [][]string
	records, _ := reader.ReadAll()
	return records
}

func csvReaderByLine(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file :", err)
		log.Fatal(err)
	}
	//init the csv reader
	reader := csv.NewReader(file)

	// line, err := reader.Read()
	// if err != nil {
	// 	fmt.Println("Error reading the line :", err)
	// 	log.Fatal(err)
	// }

	//do something with the line
	for i := 0; ; i++ {
		line, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Println("An error encountered ::", err)
		}

		fmt.Printf("Row %d : %v \n", i, line)
	}
}

func main() {
	lookupCSV("amino.csv")
	csvReaderByLine("amino.csv")
}
