package csvreader

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

/*CSVtoStringSlice : Reads a csv file and sends back a 2D Array of strings
 * params:  file object, seperator
 * returns: A 2D array (2D slices of Strings), or returns err if any error
 */
func CSVtoStringSlice(file *os.File, separator rune) ([][]string, error) {
	data := [][]string{}
	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = separator

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		data = append(data, row)
	}
	return data, nil
}
