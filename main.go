package main

import (
	"fmt"

	"github.com/t3chn0tr0n/file_fetcher/csvreader"
)

func main() {
	file, separator, err := csvreader.ConnectToSource()
	if err != nil {
		fmt.Println("ERROR: FILE CANNOT BE OPEN!")
		return
	}
	data, err := csvreader.CSVtoStringSlice(file, separator)
	if err != nil {
		fmt.Println("ERROR: CANNOT PARSE CSV!")
		return
	}
	fmt.Println(data)
}
